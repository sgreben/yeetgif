package box2d

import "math"

/// The world class manages all physics entities, dynamic simulation,
/// and asynchronous queries. The world also contains efficient memory
/// management facilities.

type WorldFlag uint32

const (
	WorldFlagNewFixture  WorldFlag = 0x0001
	WorldFlagLocked      WorldFlag = 0x0002
	WorldFlagClearForces WorldFlag = 0x0004
)

// /// The world class manages all physics entities, dynamic simulation,
// /// and asynchronous queries. The world also contains efficient memory
// /// management facilities.
type World struct {
	Flags WorldFlag

	ContactManager ContactManager

	BodyList *Body // linked list

	BodyCount int

	Gravity    Point
	AllowSleep bool

	DestructionListener DestructionListenerInterface

	// This is used to compute the time step ratio to
	// support a variable time step.
	InverseDt0 float64

	// These are for debugging the solver.
	WarmStarting      bool
	ContinuousPhysics bool
	SubStepping       bool

	StepComplete bool
}

func (world World) GetBodyList() *Body {
	return world.BodyList
}

func (world World) GetContactList() ContactInterface { // returns a pointer
	return world.ContactManager.ContactList
}

func (world World) GetBodyCount() int {
	return world.BodyCount
}

func (world World) GetContactCount() int {
	return world.ContactManager.ContactCount
}

func (world *World) SetGravity(gravity Point) {
	world.Gravity = gravity
}

func (world World) GetGravity() Point {
	return world.Gravity
}

func (world World) IsLocked() bool {
	return (world.Flags & WorldFlagLocked) == WorldFlagLocked
}

func (world *World) SetAutoClearForces(flag bool) {
	if flag {
		world.Flags |= WorldFlagClearForces
	} else {
		world.Flags &= ^WorldFlagClearForces
	}
}

/// Get the flag that controls automatic clearing of forces after each time step.
func (world World) GetAutoClearForces() bool {
	return (world.Flags & WorldFlagClearForces) == WorldFlagClearForces
}

func (world World) GetContactManager() ContactManager {
	return world.ContactManager
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// World.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func MakeWorld(gravity Point) World {

	world := World{}

	world.DestructionListener = nil

	world.BodyList = nil

	world.BodyCount = 0

	world.WarmStarting = true
	world.ContinuousPhysics = true
	world.SubStepping = false

	world.StepComplete = true

	world.AllowSleep = true
	world.Gravity = gravity

	world.Flags = WorldFlagClearForces

	world.InverseDt0 = 0.0

	world.ContactManager = MakeContactManager()

	return world
}

func (world *World) Destroy() {

	// Some shapes allocate using b2Alloc.
	b := world.BodyList
	for b != nil {
		bNext := b.Next

		f := b.FixtureList
		for f != nil {
			fNext := f.Next
			f.ProxyCount = 0
			f.Destroy()
			f = fNext
		}

		b = bNext
	}
}

func (world *World) SetDestructionListener(listener DestructionListenerInterface) {
	world.DestructionListener = listener
}

func (world *World) SetContactFilter(filter ContactFilterInterface) {
	world.ContactManager.ContactFilter = filter
}

func (world *World) SetContactListener(listener ContactListenerInterface) {
	world.ContactManager.ContactListener = listener
}

// void (world *World) SetDebugDraw(b2Draw* debugDraw)
// {
// 	g_debugDraw = debugDraw;
// }

func (world *World) CreateBody(def *BodyDef) *Body {

	if world.IsLocked() {
		return nil
	}

	b := NewBody(def, world)

	// Add to world doubly linked list.
	b.Prev = nil
	b.Next = world.BodyList
	if world.BodyList != nil {
		world.BodyList.Prev = b
	}
	world.BodyList = b
	world.BodyCount++

	return b
}

func (world *World) DestroyBody(b *Body) {

	if world.IsLocked() {
		return
	}

	// Delete the attached contacts.
	ce := b.ContactList
	for ce != nil {
		ce0 := ce
		ce = ce.Next
		world.ContactManager.Destroy(ce0.Contact)
	}
	b.ContactList = nil

	// Delete the attached fixtures. This destroys broad-phase proxies.
	f := b.FixtureList
	for f != nil {
		f0 := f
		f = f.Next

		if world.DestructionListener != nil {
			world.DestructionListener.SayGoodbyeToFixture(f0)
		}

		f0.DestroyProxies(&world.ContactManager.BroadPhase)
		f0.Destroy()

		b.FixtureList = f
		b.FixtureCount -= 1
	}

	b.FixtureList = nil
	b.FixtureCount = 0

	// Remove world body list.
	if b.Prev != nil {
		b.Prev.Next = b.Next
	}

	if b.Next != nil {
		b.Next.Prev = b.Prev
	}

	if b == world.BodyList {
		world.BodyList = b.Next
	}

	world.BodyCount--
}

func (world *World) SetAllowSleeping(flag bool) {
	if flag == world.AllowSleep {
		return
	}

	world.AllowSleep = flag
	if world.AllowSleep == false {
		for b := world.BodyList; b != nil; b = b.Next {
			b.SetAwake(true)
		}
	}
}

// Find islands, integrate and solve constraints, solve position constraints
func (world *World) Solve(step TimeStep) {

	// Size the island for the worst case.
	island := MakeIsland(
		world.BodyCount,
		world.ContactManager.ContactCount,
		world.ContactManager.ContactListener,
	)

	// Clear all the island flags.
	for b := world.BodyList; b != nil; b = b.Next {
		b.Flags &= ^BodyFlagIsland
	}
	for c := world.ContactManager.ContactList; c != nil; c = c.Data().GetNext() {
		c.Data().SetFlags(c.Data().GetFlags() & ^ContactFlagIsland)
	}

	// Build and simulate all awake islands.
	stackSize := world.BodyCount
	stack := make([]*Body, stackSize)

	for seed := world.BodyList; seed != nil; seed = seed.Next {
		if (seed.Flags & BodyFlagIsland) != 0x0000 {
			continue
		}

		if seed.IsAwake() == false || seed.IsActive() == false {
			continue
		}

		// The seed can be dynamic or kinematic.
		if seed.GetType() == BodyTypeStaticBody {
			continue
		}

		// Reset island and stack.
		island.Clear()
		stackCount := 0
		stack[stackCount] = seed
		stackCount++
		seed.Flags |= BodyFlagIsland

		// Perform a depth first search (DFS) on the constraint graph.
		for stackCount > 0 {
			// Grab the next body off the stack and add it to the island.
			stackCount--
			b := stack[stackCount]
			island.AddBody(b)

			// Make sure the body is awake (without resetting sleep timer).
			b.Flags |= BodyFlagAwake

			// To keep islands as small as possible, we don't
			// propagate islands across static bodies.
			if b.GetType() == BodyTypeStaticBody {
				continue
			}

			// Search all contacts connected to this body.
			for ce := b.ContactList; ce != nil; ce = ce.Next {
				contact := ce.Contact
				contactData := contact.Data()

				// Has this contact already been added to an island?
				if (contactData.GetFlags() & ContactFlagIsland) != 0x0000 {
					continue
				}

				// Is this contact solid and touching?
				if contactData.IsEnabled() == false || contactData.IsTouching() == false {
					continue
				}

				// Skip sensors.
				sensorA := contactData.GetFixtureA().IsSensor
				sensorB := contactData.GetFixtureB().IsSensor

				if sensorA || sensorB {
					continue
				}

				island.AddContact(contact)
				contactData.SetFlags(contactData.GetFlags() | ContactFlagIsland)

				other := ce.Other

				// Was the other body already added to this island?
				if (other.Flags & BodyFlagIsland) != 0x0000 {
					continue
				}

				stack[stackCount] = other
				stackCount++
				other.Flags |= BodyFlagIsland
			}
		}

		island.Solve(step, world.Gravity, world.AllowSleep)

		// Post solve cleanup.
		for i := 0; i < island.BodyCount; i++ {
			// Allow static bodies to participate in other islands.
			b := island.Bodies[i]
			if b.GetType() == BodyTypeStaticBody {
				b.Flags &= ^BodyFlagIsland
			}
		}
	}

	stack = nil

	{
		// Synchronize fixtures, check for out of range bodies.
		for b := world.BodyList; b != nil; b = b.GetNext() {
			// If a body was not in an island then it did not move.
			if (b.Flags & BodyFlagIsland) == 0 {
				continue
			}

			if b.GetType() == BodyTypeStaticBody {
				continue
			}

			// Update fixtures (for broad-phase).
			b.SynchronizeFixtures()
		}

		// Look for new contacts.
		world.ContactManager.FindNewContacts()
	}
}

// Find TOI contacts and solve them.
func (world *World) SolveTOI(step TimeStep) {

	island := MakeIsland(2*_maxTOIContacts, _maxTOIContacts, world.ContactManager.ContactListener)

	if world.StepComplete {
		for b := world.BodyList; b != nil; b = b.Next {
			b.Flags &= ^BodyFlagIsland
			b.Sweep.Alpha0 = 0.0
		}

		for c := world.ContactManager.ContactList; c != nil; c = c.Data().GetNext() {
			contactData := c.Data()
			// Invalidate TOI
			contactData.SetFlags(contactData.GetFlags() & ^(ContactFlagToi | ContactFlagIsland))
			contactData.SetTOICount(0)
			contactData.SetTOI(1.0)
		}
	}

	// Find TOI events and solve them.
	for {
		// Find the first TOI.
		var minContact ContactInterface = nil // has to be a pointer
		minAlpha := 1.0

		for c := world.ContactManager.ContactList; c != nil; c = c.Data().GetNext() {
			contactData := c.Data()

			// Is this contact disabled?
			if contactData.IsEnabled() == false {
				continue
			}

			// Prevent excessive sub-stepping.
			if contactData.GetTOICount() > _maxSubSteps {
				continue
			}

			alpha := 1.0
			if (contactData.GetFlags() & ContactFlagToi) != 0x0000 {
				// This contact has a valid cached TOI.
				alpha = contactData.GetTOI()
			} else {
				fA := contactData.GetFixtureA()
				fB := contactData.GetFixtureB()

				// Is there a sensor?
				if fA.IsSensor || fB.IsSensor {
					continue
				}

				bA := fA.GetBody()
				bB := fB.GetBody()

				typeA := bA.Type
				typeB := bB.Type

				activeA := bA.IsAwake() && typeA != BodyTypeStaticBody
				activeB := bB.IsAwake() && typeB != BodyTypeStaticBody

				// Is at least one body active (awake and dynamic or kinematic)?
				if activeA == false && activeB == false {
					continue
				}

				collideA := bA.IsBullet() || typeA != BodyTypeDynamicBody
				collideB := bB.IsBullet() || typeB != BodyTypeDynamicBody

				// Are these two non-bullet dynamic bodies?
				if collideA == false && collideB == false {
					continue
				}

				// Compute the TOI for this contact.
				// Put the sweeps onto the same time interval.
				alpha0 := bA.Sweep.Alpha0

				if bA.Sweep.Alpha0 < bB.Sweep.Alpha0 {
					alpha0 = bB.Sweep.Alpha0
					bA.Sweep.Advance(alpha0)
				} else if bB.Sweep.Alpha0 < bA.Sweep.Alpha0 {
					alpha0 = bA.Sweep.Alpha0
					bB.Sweep.Advance(alpha0)
				}

				indexA := contactData.GetChildIndexA()
				indexB := contactData.GetChildIndexB()

				// Compute the time of impact in interval [0, minTOI]
				input := TOIInput{}
				input.ProxyA.Set(fA.GetShape(), indexA)
				input.ProxyB.Set(fB.GetShape(), indexB)
				input.SweepA = bA.Sweep
				input.SweepB = bB.Sweep
				input.TMax = 1.0

				output := TOIOutput{}
				TimeOfImpact(&output, &input)

				// Beta is the fraction of the remaining portion of the .
				beta := output.T
				if output.State == TOIOutputStateTouching {
					alpha = math.Min(alpha0+(1.0-alpha0)*beta, 1.0)
				} else {
					alpha = 1.0
				}

				contactData.SetTOI(alpha)
				contactData.SetFlags(contactData.GetFlags() | ContactFlagToi)
			}

			if alpha < minAlpha {
				// This is the minimum TOI found so far.
				minContact = c
				minAlpha = alpha
			}
		}

		if minContact == nil || 1.0-10.0*_epsilon < minAlpha {
			// No more TOI events. Done!
			world.StepComplete = true
			break
		}

		// Advance the bodies to the TOI.
		fA := minContact.Data().GetFixtureA()
		fB := minContact.Data().GetFixtureB()
		bA := fA.GetBody()
		bB := fB.GetBody()

		backup1 := bA.Sweep
		backup2 := bB.Sweep

		bA.Advance(minAlpha)
		bB.Advance(minAlpha)

		// The TOI contact likely has some new contact points.
		ContactUpdate(minContact, world.ContactManager.ContactListener)
		minContactData := minContact.Data()
		minContactData.SetFlags(minContactData.GetFlags() & ^ContactFlagToi)
		minContactData.SetTOICount(minContactData.GetTOICount() + 1)

		// Is the contact solid?
		if minContactData.IsEnabled() == false || minContactData.IsTouching() == false {
			// Restore the sweeps.
			minContactData.SetEnabled(false)
			bA.Sweep = backup1
			bB.Sweep = backup2
			bA.SynchronizeTransform()
			bB.SynchronizeTransform()
			continue
		}

		bA.SetAwake(true)
		bB.SetAwake(true)

		// Build the island
		island.Clear()
		island.AddBody(bA)
		island.AddBody(bB)
		island.AddContact(minContact)

		bA.Flags |= BodyFlagIsland
		bB.Flags |= BodyFlagIsland
		minContactData.SetFlags(minContactData.GetFlags() | ContactFlagIsland)

		// Get contacts on bodyA and bodyB.
		bodies := [2]*Body{bA, bB}

		for i := 0; i < 2; i++ {
			body := bodies[i]
			if body.Type == BodyTypeDynamicBody {
				for ce := body.ContactList; ce != nil; ce = ce.Next {
					if island.BodyCount == island.BodyCapacity {
						break
					}

					if island.ContactCount == island.ContactCapacity {
						break
					}

					contact := ce.Contact
					contactData := contact.Data()

					// Has this contact already been added to the island?
					if (contactData.GetFlags() & ContactFlagIsland) != 0x0000 {
						continue
					}

					// Only add static, kinematic, or bullet bodies.
					other := ce.Other
					if other.Type == BodyTypeDynamicBody && body.IsBullet() == false && other.IsBullet() == false {
						continue
					}

					// Skip sensors.
					sensorA := contactData.GetFixtureA().IsSensor
					sensorB := contactData.GetFixtureB().IsSensor
					if sensorA || sensorB {
						continue
					}

					// Tentatively advance the body to the TOI.
					backup := other.Sweep
					if (other.Flags & BodyFlagIsland) == 0 {
						other.Advance(minAlpha)
					}

					// Update the contact points
					ContactUpdate(contact, world.ContactManager.ContactListener)

					// Was the contact disabled by the user?
					if contactData.IsEnabled() == false {
						other.Sweep = backup
						other.SynchronizeTransform()
						continue
					}

					// Are there contact points?
					if contactData.IsTouching() == false {
						other.Sweep = backup
						other.SynchronizeTransform()
						continue
					}

					// Add the contact to the island
					contactData.SetFlags(contactData.GetFlags() | ContactFlagIsland)
					island.AddContact(contact)

					// Has the other body already been added to the island?
					if (other.Flags & BodyFlagIsland) != 0x0000 {
						continue
					}

					// Add the other body to the island.
					other.Flags |= BodyFlagIsland

					if other.Type != BodyTypeStaticBody {
						other.SetAwake(true)
					}

					island.AddBody(other)
				}
			}
		}

		subStep := TimeStep{}
		subStep.Dt = (1.0 - minAlpha) * step.Dt
		subStep.InverseDt = 1.0 / subStep.Dt
		subStep.DtRatio = 1.0
		subStep.PositionIterations = 20
		subStep.VelocityIterations = step.VelocityIterations
		subStep.WarmStarting = false
		island.SolveTOI(subStep, bA.IslandIndex, bB.IslandIndex)

		// Reset island flags and synchronize broad-phase proxies.
		for i := 0; i < island.BodyCount; i++ {
			body := island.Bodies[i]
			body.Flags &= ^BodyFlagIsland

			if body.Type != BodyTypeDynamicBody {
				continue
			}

			body.SynchronizeFixtures()

			// Invalidate all contact TOIs on this displaced body.
			for ce := body.ContactList; ce != nil; ce = ce.Next {
				contactData := ce.Contact.Data()
				contactData.SetFlags(contactData.GetFlags() & ^(ContactFlagToi | ContactFlagIsland))
			}
		}

		// Commit fixture proxy movements to the broad-phase so that new contacts are created.
		// Also, some contacts can be destroyed.
		world.ContactManager.FindNewContacts()

		if world.SubStepping {
			world.StepComplete = false
			break
		}
	}
}

func (world *World) Step(dt float64, velocityIterations int, positionIterations int) {

	// If new fixtures were added, we need to find the new contacts.
	if (world.Flags & WorldFlagNewFixture) != 0x0000 {
		world.ContactManager.FindNewContacts()
		world.Flags &= ^WorldFlagNewFixture
	}

	world.Flags |= WorldFlagLocked

	step := TimeStep{}
	step.Dt = dt
	step.VelocityIterations = velocityIterations
	step.PositionIterations = positionIterations
	if dt > 0.0 {
		step.InverseDt = 1.0 / dt
	} else {
		step.InverseDt = 0.0
	}

	step.DtRatio = world.InverseDt0 * dt

	step.WarmStarting = world.WarmStarting

	// Update contacts. This is where some contacts are destroyed.
	{
		world.ContactManager.Collide()
	}

	// Integrate velocities, solve velocity constraints, and integrate positions.
	if world.StepComplete && step.Dt > 0.0 {
		world.Solve(step)
	}

	// Handle TOI events.
	if world.ContinuousPhysics && step.Dt > 0.0 {
		world.SolveTOI(step)
	}

	if step.Dt > 0.0 {
		world.InverseDt0 = step.InverseDt
	}

	if (world.Flags & WorldFlagClearForces) != 0x0000 {
		world.ClearForces()
	}

	world.Flags &= ^WorldFlagLocked

}

func (world *World) ClearForces() {
	for body := world.BodyList; body != nil; body = body.GetNext() {
		body.Force.SetZero()
		body.Torque = 0.0
	}
}

type WorldQueryWrapper struct {
	BroadPhase *BroadPhase
	Callback   BroadPhaseQueryCallback
}

func MakeWorldQueryWrapper() WorldQueryWrapper {
	return WorldQueryWrapper{}
}

func (query *WorldQueryWrapper) QueryCallback(proxyId int) bool {
	proxy := query.BroadPhase.GetUserData(proxyId).(*FixtureProxy)
	return query.Callback(proxy.Fixture)
}

func (world *World) QueryAABB(callback BroadPhaseQueryCallback, aabb AABB) {
	wrapper := MakeWorldQueryWrapper()
	wrapper.BroadPhase = &world.ContactManager.BroadPhase
	wrapper.Callback = callback
	world.ContactManager.BroadPhase.Query(wrapper.QueryCallback, aabb)
}

func (world *World) RayCast(callback RaycastCallback, point1 Point, point2 Point) {

	// TreeRayCastCallback
	wrapper := func(input RayCastInput, nodeId int) float64 {

		userData := world.ContactManager.BroadPhase.GetUserData(nodeId)
		proxy := userData.(*FixtureProxy)
		fixture := proxy.Fixture
		index := proxy.ChildIndex
		output := RayCastOutput{}
		hit := fixture.RayCast(&output, input, index)

		if hit {
			fraction := output.Fraction
			point := PointAdd(PointMulScalar((1.0-fraction), input.P1), PointMulScalar(fraction, input.P2))
			return callback(fixture, point, output.Normal, fraction)
		}

		return input.MaxFraction
	}

	input := RayCastInput{}
	input.MaxFraction = 1.0
	input.P1 = point1
	input.P2 = point2
	world.ContactManager.BroadPhase.RayCast(wrapper, input)
}

func (world World) GetProxyCount() int {
	return world.ContactManager.BroadPhase.GetProxyCount()
}

func (world World) GetTreeHeight() int {
	return world.ContactManager.BroadPhase.GetTreeHeight()
}

func (world World) GetTreeBalance() int {
	return world.ContactManager.BroadPhase.GetTreeBalance()
}

func (world World) GetTreeQuality() float64 {
	return world.ContactManager.BroadPhase.GetTreeQuality()
}

func (world *World) ShiftOrigin(newOrigin Point) {

	if (world.Flags & WorldFlagLocked) == WorldFlagLocked {
		return
	}

	for b := world.BodyList; b != nil; b = b.Next {
		b.Xf.P.OperatorMinusInplace(newOrigin)
		b.Sweep.C0.OperatorMinusInplace(newOrigin)
		b.Sweep.C.OperatorMinusInplace(newOrigin)
	}

	world.ContactManager.BroadPhase.ShiftOrigin(newOrigin)
}
