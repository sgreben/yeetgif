package box2d

import (
	"math"
)

/// This is an internal class.
type Island struct {
	Listener ContactListenerInterface

	Bodies   []*Body
	Contacts []ContactInterface // has to be backed by pointers

	Positions  []Position
	Velocities []Velocity

	BodyCount    int
	ContactCount int

	BodyCapacity    int
	ContactCapacity int
}

func (island *Island) Clear() {
	island.BodyCount = 0
	island.ContactCount = 0
}

func (island *Island) AddBody(body *Body) {
	body.IslandIndex = island.BodyCount
	island.Bodies[island.BodyCount] = body
	island.BodyCount++
}

func (island *Island) AddContact(contact ContactInterface) { // contact has to be a pointer
	island.Contacts[island.ContactCount] = contact
	island.ContactCount++
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Island.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

/*
Position Correction Notes
=========================
I tried the several algorithms for position correction of the 2D revolute joint.
I looked at these systems:
- simple pendulum (1m diameter sphere on massless 5m stick) with initial angular velocity of 100 rad/s.
- suspension bridge with 30 1m long planks of length 1m.
- multi-link chain with 30 1m long links.

Here are the algorithms:

Baumgarte - A fraction of the position error is added to the velocity error. There is no
separate position solver.

Pseudo Velocities - After the velocity solver and position integration,
the position error, Jacobian, and effective mass are recomputed. Then
the velocity constraints are solved with pseudo velocities and a fraction
of the position error is added to the pseudo velocity error. The pseudo
velocities are initialized to zero and there is no warm-starting. After
the position solver, the pseudo velocities are added to the positions.
This is also called the First Order World method or the Position LCP method.

Modified Nonlinear Gauss-Seidel (NGS) - Like Pseudo Velocities except the
position error is re-computed for each constraint and the positions are updated
after the constraint is solved. The radius vectors (aka Jacobians) are
re-computed too (otherwise the algorithm has horrible instability). The pseudo
velocity states are not needed because they are effectively zero at the beginning
of each iteration. Since we have the current position error, we allow the
iterations to terminate early if the error becomes smaller than b2_linearSlop.

Full NGS or just NGS - Like Modified NGS except the effective mass are re-computed
each time a constraint is solved.

Here are the results:
Baumgarte - this is the cheapest algorithm but it has some stability problems,
especially with the bridge. The chain links separate easily close to the root
and they jitter as they struggle to pull together. This is one of the most common
methods in the field. The big drawback is that the position correction artificially
affects the momentum, thus leading to instabilities and false bounce. I used a
bias factor of 0.2. A larger bias factor makes the bridge less stable, a smaller
factor makes joints and contacts more spongy.

Pseudo Velocities - the is more stable than the Baumgarte method. The bridge is
stable. However, joints still separate with large angular velocities. Drag the
simple pendulum in a circle quickly and the joint will separate. The chain separates
easily and does not recover. I used a bias factor of 0.2. A larger value lead to
the bridge collapsing when a heavy cube drops on it.

Modified NGS - this algorithm is better in some ways than Baumgarte and Pseudo
Velocities, but in other ways it is worse. The bridge and chain are much more
stable, but the simple pendulum goes unstable at high angular velocities.

Full NGS - stable in all tests. The joints display good stiffness. The bridge
still sags, but this is better than infinite forces.

Recommendations
Pseudo Velocities are not really worthwhile because the bridge and chain cannot
recover from joint separation. In other cases the benefit over Baumgarte is small.

Modified NGS is not a robust method for the revolute joint due to the violent
instability seen in the simple pendulum. Perhaps it is viable with other constraint
types, especially scalar constraints where the effective mass is a scalar.

This leaves Baumgarte and Full NGS. Baumgarte has small, but manageable instabilities
and is very fast. I don't think we can escape Baumgarte, especially in highly
demanding cases where high constraint fidelity is not needed.

Full NGS is robust and easy on the eyes. I recommend this as an option for
higher fidelity simulation and certainly for suspension bridges and long chains.
Full NGS might be a good choice for ragdolls, especially motorized ragdolls where
joint separation can be problematic. The number of NGS iterations can be reduced
for better performance without harming robustness much.

Each joint in a can be handled differently in the position solver. So I recommend
a system where the user can select the algorithm on a per joint basis. I would
probably default to the slower Full NGS and let the user select the faster
Baumgarte method in performance critical scenarios.
*/

/*
Cache Performance

The Box2D solvers are dominated by cache misses. Data structures are designed
to increase the number of cache hits. Much of misses are due to random access
to body data. The constraint structures are iterated over linearly, which leads
to few cache misses.

The bodies are not accessed during iteration. Instead read only data, such as
the mass values are stored with the constraints. The mutable data are the constraint
impulses and the bodies velocities/positions. The impulses are held inside the
constraint structures. The body velocities/positions are held in compact, temporary
arrays to increase the number of cache hits. Linear and angular velocity are
stored in a single array since multiple arrays lead to multiple misses.
*/

/*
2D Rotation

R = [cos(theta) -sin(theta)]
    [sin(theta) cos(theta) ]

thetaDot = omega

Let q1 = cos(theta), q2 = sin(theta).
R = [q1 -q2]
    [q2  q1]

q1Dot = -thetaDot * q2
q2Dot = thetaDot * q1

q1_new = q1_old - dt * w * q2
q2_new = q2_old + dt * w * q1
then normalize.

This might be faster than computing sin+cos.
However, we can compute sin+cos of the same angle fast.
*/

func MakeIsland(bodyCapacity int, contactCapacity int, listener ContactListenerInterface) Island {

	island := Island{}

	island.BodyCapacity = bodyCapacity
	island.ContactCapacity = contactCapacity
	island.BodyCount = 0
	island.ContactCount = 0

	island.Listener = listener

	island.Bodies = make([]*Body, bodyCapacity)
	island.Contacts = make([]ContactInterface, contactCapacity)

	island.Velocities = make([]Velocity, bodyCapacity)
	island.Positions = make([]Position, bodyCapacity)

	return island
}

func (island *Island) Destroy() {

}

func (island *Island) Solve(step TimeStep, gravity Point, allowSleep bool) {
	h := step.Dt
	// Integrate velocities and apply damping. Initialize the body state.
	for i := 0; i < island.BodyCount; i++ {
		b := island.Bodies[i]

		c := b.Sweep.C
		a := b.Sweep.A
		v := b.LinearVelocity
		w := b.AngularVelocity

		// Store positions for continuous collision.
		b.Sweep.C0 = b.Sweep.C
		b.Sweep.A0 = b.Sweep.A

		if b.Type == BodyTypeDynamicBody {

			// Integrate velocities.
			v.OperatorPlusInplace(
				PointMulScalar(
					h,
					PointAdd(
						PointMulScalar(b.GravityScale, gravity),
						PointMulScalar(b.InvMass, b.Force),
					),
				),
			)
			w += h * b.InvI * b.Torque

			// Apply damping.
			// ODE: dv/dt + c * v = 0
			// Solution: v(t) = v0 * exp(-c * t)
			// Time step: v(t + dt) = v0 * exp(-c * (t + dt)) = v0 * exp(-c * t) * exp(-c * dt) = v * exp(-c * dt)
			// v2 = exp(-c * dt) * v1
			// Pade approximation:
			// v2 = v1 * 1 / (1 + c * dt)
			v.OperatorScalarMulInplace(1.0 / (1.0 + h*b.LinearDamping))
			w *= 1.0 / (1.0 + h*b.AngularDamping)
		}

		island.Positions[i].C = c
		island.Positions[i].A = a
		island.Velocities[i].V = v
		island.Velocities[i].W = w
	}

	// Solver data
	solverData := SolverData{}
	solverData.Step = step
	solverData.Positions = island.Positions
	solverData.Velocities = island.Velocities

	// Initialize velocity constraints.
	contactSolverDef := MakeContactSolverDef()
	contactSolverDef.Step = step
	contactSolverDef.Contacts = island.Contacts
	contactSolverDef.Count = island.ContactCount
	contactSolverDef.Positions = island.Positions
	contactSolverDef.Velocities = island.Velocities

	contactSolver := MakeContactSolver(&contactSolverDef)
	contactSolver.InitializeVelocityConstraints()

	if step.WarmStarting {
		contactSolver.WarmStart()
	}

	// Solve velocity constraints
	for i := 0; i < step.VelocityIterations; i++ {
		contactSolver.SolveVelocityConstraints()
	}

	// Store impulses for warm starting
	contactSolver.StoreImpulses()

	// Integrate positions
	for i := 0; i < island.BodyCount; i++ {
		c := island.Positions[i].C
		a := island.Positions[i].A
		v := island.Velocities[i].V
		w := island.Velocities[i].W

		// Check for large velocities
		translation := PointMulScalar(h, v)
		if PointDot(translation, translation) > _maxTranslationSquared {
			ratio := _maxTranslation / translation.Length()
			v.OperatorScalarMulInplace(ratio)
		}

		rotation := h * w
		if rotation*rotation > _maxRotationSquared {
			ratio := _maxRotation / math.Abs(rotation)
			w *= ratio
		}

		// Integrate
		c.OperatorPlusInplace(PointMulScalar(h, v))
		a += h * w

		island.Positions[i].C = c
		island.Positions[i].A = a
		island.Velocities[i].V = v
		island.Velocities[i].W = w
	}

	// Solve position constraints
	positionSolved := false
	for i := 0; i < step.PositionIterations; i++ {
		contactsOkay := contactSolver.SolvePositionConstraints()

		jointsOkay := true
		if contactsOkay && jointsOkay {
			// Exit early if the position errors are small.
			positionSolved = true
			break
		}
	}

	// Copy state buffers back to the bodies
	for i := 0; i < island.BodyCount; i++ {
		body := island.Bodies[i]
		body.Sweep.C = island.Positions[i].C
		body.Sweep.A = island.Positions[i].A
		body.LinearVelocity = island.Velocities[i].V
		body.AngularVelocity = island.Velocities[i].W
		body.SynchronizeTransform()
	}

	island.Report(contactSolver.VelocityConstraints)

	if allowSleep {
		minSleepTime := math.MaxFloat64

		linTolSqr := _linearSleepTolerance * _linearSleepTolerance
		angTolSqr := _angularSleepTolerance * _angularSleepTolerance

		for i := 0; i < island.BodyCount; i++ {
			b := island.Bodies[i]
			if b.GetType() == BodyTypeStaticBody {
				continue
			}

			if (b.Flags&BodyFlagAutoSleep) == 0 || b.AngularVelocity*b.AngularVelocity > angTolSqr || PointDot(b.LinearVelocity, b.LinearVelocity) > linTolSqr {
				b.SleepTime = 0.0
				minSleepTime = 0.0
			} else {
				b.SleepTime += h
				minSleepTime = math.Min(minSleepTime, b.SleepTime)
			}
		}

		if minSleepTime >= _timeToSleep && positionSolved {
			for i := 0; i < island.BodyCount; i++ {
				b := island.Bodies[i]
				b.SetAwake(false)
			}
		}
	}
}

func (island *Island) SolveTOI(subStep TimeStep, toiIndexA int, toiIndexB int) {

	// Initialize the body state.
	for i := 0; i < island.BodyCount; i++ {
		b := island.Bodies[i]
		island.Positions[i].C = b.Sweep.C
		island.Positions[i].A = b.Sweep.A
		island.Velocities[i].V = b.LinearVelocity
		island.Velocities[i].W = b.AngularVelocity
	}

	contactSolverDef := MakeContactSolverDef()

	contactSolverDef.Contacts = island.Contacts
	contactSolverDef.Count = island.ContactCount
	contactSolverDef.Step = subStep
	contactSolverDef.Positions = island.Positions
	contactSolverDef.Velocities = island.Velocities
	contactSolver := MakeContactSolver(&contactSolverDef)

	// Solve position constraints.
	for i := 0; i < subStep.PositionIterations; i++ {
		contactsOkay := contactSolver.SolveTOIPositionConstraints(toiIndexA, toiIndexB)
		if contactsOkay {
			break
		}
	}

	// Leap of faith to new safe state.
	island.Bodies[toiIndexA].Sweep.C0 = island.Positions[toiIndexA].C
	island.Bodies[toiIndexA].Sweep.A0 = island.Positions[toiIndexA].A
	island.Bodies[toiIndexB].Sweep.C0 = island.Positions[toiIndexB].C
	island.Bodies[toiIndexB].Sweep.A0 = island.Positions[toiIndexB].A

	// No warm starting is needed for TOI events because warm
	// starting impulses were applied in the discrete solver.
	contactSolver.InitializeVelocityConstraints()

	// Solve velocity constraints.
	for i := 0; i < subStep.VelocityIterations; i++ {
		contactSolver.SolveVelocityConstraints()
	}

	// Don't store the TOI contact forces for warm starting
	// because they can be quite large.

	h := subStep.Dt

	// Integrate positions
	for i := 0; i < island.BodyCount; i++ {
		c := island.Positions[i].C
		a := island.Positions[i].A
		v := island.Velocities[i].V
		w := island.Velocities[i].W

		// Check for large velocities
		translation := PointMulScalar(h, v)
		if PointDot(translation, translation) > _maxTranslationSquared {
			ratio := _maxTranslation / translation.Length()
			v.OperatorScalarMulInplace(ratio)
		}

		rotation := h * w
		if rotation*rotation > _maxRotationSquared {
			ratio := _maxRotation / math.Abs(rotation)
			w *= ratio
		}

		// Integrate
		c.OperatorPlusInplace(PointMulScalar(h, v))
		a += h * w

		island.Positions[i].C = c
		island.Positions[i].A = a
		island.Velocities[i].V = v
		island.Velocities[i].W = w

		// Sync bodies
		body := island.Bodies[i]
		body.Sweep.C = c
		body.Sweep.A = a
		body.LinearVelocity = v
		body.AngularVelocity = w
		body.SynchronizeTransform()
	}

	island.Report(contactSolver.VelocityConstraints)
}

func (island *Island) Report(constraints []ContactVelocityConstraint) {
	if island.Listener == nil {
		return
	}

	for i := 0; i < island.ContactCount; i++ {
		c := island.Contacts[i]

		vc := constraints[i]

		impulse := MakeContactImpulse()
		impulse.Count = vc.PointCount

		for j := 0; j < vc.PointCount; j++ {
			impulse.NormalImpulses[j] = vc.Points[j].NormalImpulse
			impulse.TangentImpulses[j] = vc.Points[j].TangentImpulse
		}

		island.Listener.PostSolve(c, &impulse)
	}
}
