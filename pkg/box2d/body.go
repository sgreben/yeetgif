package box2d

type BodyType uint8

// Body types
// static: zero mass, zero velocity, may be manually moved
// kinematic: zero mass, non-zero velocity set by user, moved by solver
// dynamic: positive mass, non-zero velocity determined by forces, moved by solver
const (
	BodyTypeStaticBody    BodyType = 0
	BodyTypeKinematicBody BodyType = 1
	BodyTypeDynamicBody   BodyType = 2
)

// BodyDef is a body definition. It holds all the data needed to construct a rigid body.
// You can safely re-use body definitions. Shapes are added to a body after construction.
type BodyDef struct {

	/// The body type: static, kinematic, or dynamic.
	/// Note: if a dynamic body would have zero mass, the mass is set to one.
	Type BodyType

	/// The world position of the body. Avoid creating bodies at the origin
	/// since this can lead to many overlapping shapes.
	Position Point

	/// The world angle of the body in radians.
	Angle float64

	/// The linear velocity of the body's origin in world co-ordinates.
	LinearVelocity Point

	/// The angular velocity of the body.
	AngularVelocity float64

	/// Linear damping is use to reduce the linear velocity. The damping parameter
	/// can be larger than 1.0 but the damping effect becomes sensitive to the
	/// time step when the damping parameter is large.
	/// Units are 1/time
	LinearDamping float64

	/// Angular damping is use to reduce the angular velocity. The damping parameter
	/// can be larger than 1.0 but the damping effect becomes sensitive to the
	/// time step when the damping parameter is large.
	/// Units are 1/time
	AngularDamping float64

	/// Set this flag to false if this body should never fall asleep. Note that
	/// this increases CPU usage.
	AllowSleep bool

	/// Is this body initially awake or sleeping?
	Awake bool

	/// Should this body be prevented from rotating? Useful for characters.
	FixedRotation bool

	/// Is this a fast moving body that should be prevented from tunneling through
	/// other moving bodies? Note that all bodies are prevented from tunneling through
	/// kinematic and static bodies. This setting is only considered on dynamic bodies.
	/// @warning You should use this flag sparingly since it increases processing time.
	Bullet bool

	/// Does this body start out active?
	Active bool

	/// Scale the gravity applied to this body.
	GravityScale float64
}

// MakeBodyDef sets the body definition default values.
func MakeBodyDef() BodyDef {
	return BodyDef{
		Position:        Point{},
		Angle:           0.0,
		LinearVelocity:  Point{},
		AngularVelocity: 0.0,
		LinearDamping:   0.0,
		AngularDamping:  0.0,
		AllowSleep:      true,
		Awake:           true,
		FixedRotation:   false,
		Bullet:          false,
		Type:            BodyTypeStaticBody,
		Active:          true,
		GravityScale:    1.0,
	}
}

func NewBodyDef() *BodyDef {
	res := MakeBodyDef()
	return &res
}

type BodyFlag uint32

const (
	BodyFlagIsland        BodyFlag = 0x0001
	BodyFlagAwake         BodyFlag = 0x0002
	BodyFlagAutoSleep     BodyFlag = 0x0004
	BodyFlagBullet        BodyFlag = 0x0008
	BodyFlagFixedRotation BodyFlag = 0x0010
	BodyFlagActive        BodyFlag = 0x0020
	BodyFlagToi           BodyFlag = 0x0040
)

type Body struct {
	Type BodyType

	Flags BodyFlag

	IslandIndex int

	Xf    Transform // the body origin transform
	Sweep Sweep     // the swept motion for CCD

	LinearVelocity  Point
	AngularVelocity float64

	Force  Point
	Torque float64

	World *World
	Prev  *Body
	Next  *Body

	FixtureList  *Fixture // linked list
	FixtureCount int

	ContactList *ContactEdge // linked list

	Mass, InvMass float64

	// Rotational inertia about the center of mass.
	I, InvI float64

	LinearDamping  float64
	AngularDamping float64
	GravityScale   float64

	SleepTime float64

	UserData interface{}
}

func (body Body) GetType() BodyType {
	return body.Type
}

func (body Body) GetTransform() Transform {
	return body.Xf
}

func (body Body) GetPosition() Point {
	return body.Xf.P
}

func (body Body) GetAngle() float64 {
	return body.Sweep.A
}

func (body Body) GetWorldCenter() Point {
	return body.Sweep.C
}

func (body Body) GetLocalCenter() Point {
	return body.Sweep.LocalCenter
}

func (body *Body) SetLinearVelocity(v Point) {
	if body.Type == BodyTypeStaticBody {
		return
	}

	if PointDot(v, v) > 0.0 {
		body.SetAwake(true)
	}

	body.LinearVelocity = v
}

func (body Body) GetLinearVelocity() Point {
	return body.LinearVelocity
}

func (body *Body) SetAngularVelocity(w float64) {
	if body.Type == BodyTypeStaticBody {
		return
	}

	if w*w > 0.0 {
		body.SetAwake(true)
	}

	body.AngularVelocity = w
}

func (body Body) GetAngularVelocity() float64 {
	return body.AngularVelocity
}

func (body Body) GetMass() float64 {
	return body.Mass
}

func (body Body) GetInertia() float64 {
	return body.I + body.Mass*PointDot(body.Sweep.LocalCenter, body.Sweep.LocalCenter)
}

func (body Body) GetMassData(data *MassData) {
	data.Mass = body.Mass
	data.I = body.I + body.Mass*PointDot(body.Sweep.LocalCenter, body.Sweep.LocalCenter)
	data.Center = body.Sweep.LocalCenter
}

func (body Body) GetWorldPoint(localPoint Point) Point {
	return TransformPointMul(body.Xf, localPoint)
}

func (body Body) GetWorldVector(localVector Point) Point {
	return RotPointMul(body.Xf.Q, localVector)
}

func (body Body) GetLocalPoint(worldPoint Point) Point {
	return TransformPointMulT(body.Xf, worldPoint)
}

func (body Body) GetLocalVector(worldVector Point) Point {
	return RotPointMulT(body.Xf.Q, worldVector)
}

func (body Body) GetLinearVelocityFromWorldPoint(worldPoint Point) Point {
	return PointAdd(body.LinearVelocity, PointCrossScalarVector(body.AngularVelocity, PointSub(worldPoint, body.Sweep.C)))
}

func (body Body) GetLinearVelocityFromLocalPoint(localPoint Point) Point {
	return body.GetLinearVelocityFromWorldPoint(body.GetWorldPoint(localPoint))
}

func (body Body) GetLinearDamping() float64 {
	return body.LinearDamping
}

func (body *Body) SetLinearDamping(linearDamping float64) {
	body.LinearDamping = linearDamping
}

func (body Body) GetAngularDamping() float64 {
	return body.AngularDamping
}

func (body *Body) SetAngularDamping(angularDamping float64) {
	body.AngularDamping = angularDamping
}

func (body Body) GetGravityScale() float64 {
	return body.GravityScale
}

func (body *Body) SetGravityScale(scale float64) {
	body.GravityScale = scale
}

func (body *Body) SetBullet(flag bool) {
	if flag {
		body.Flags |= BodyFlagBullet
	} else {
		body.Flags &= ^BodyFlagBullet
	}
}

func (body Body) IsBullet() bool {
	return (body.Flags & BodyFlagBullet) == BodyFlagBullet
}

func (body *Body) SetAwake(flag bool) {
	if flag {
		body.Flags |= BodyFlagAwake
		body.SleepTime = 0.0
	} else {
		body.Flags &= ^BodyFlagAwake
		body.SleepTime = 0.0
		body.LinearVelocity.SetZero()
		body.AngularVelocity = 0.0
		body.Force.SetZero()
		body.Torque = 0.0
	}
}

func (body Body) IsAwake() bool {
	return (body.Flags & BodyFlagAwake) == BodyFlagAwake
}

func (body Body) IsActive() bool {
	return (body.Flags & BodyFlagActive) == BodyFlagActive
}

func (body Body) IsFixedRotation() bool {
	return (body.Flags & BodyFlagFixedRotation) == BodyFlagFixedRotation
}

func (body *Body) SetSleepingAllowed(flag bool) {
	if flag {
		body.Flags |= BodyFlagAutoSleep
	} else {
		body.Flags &= ^BodyFlagAutoSleep
		body.SetAwake(true)
	}
}

func (body Body) IsSleepingAllowed() bool {
	return (body.Flags & BodyFlagAutoSleep) == BodyFlagAutoSleep
}

func (body Body) GetFixtureList() *Fixture {
	return body.FixtureList
}

func (body Body) GetContactList() *ContactEdge {
	return body.ContactList
}

func (body Body) GetNext() *Body {
	return body.Next
}

func (body *Body) SetUserData(data interface{}) {
	body.UserData = data
}

func (body Body) GetUserData() interface{} {
	return body.UserData
}

func (body *Body) ApplyForce(force Point, point Point, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate a force if the body is sleeping.
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.Force.OperatorPlusInplace(force)
		body.Torque += PointCross(
			PointSub(point, body.Sweep.C),
			force,
		)
	}
}

func (body *Body) ApplyForceToCenter(force Point, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate a force if the body is sleeping
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.Force.OperatorPlusInplace(force)
	}
}

func (body *Body) ApplyTorque(torque float64, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate a force if the body is sleeping
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.Torque += torque
	}
}

func (body *Body) ApplyLinearImpulse(impulse Point, point Point, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate velocity if the body is sleeping
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.LinearVelocity.OperatorPlusInplace(PointMulScalar(body.InvMass, impulse))
		body.AngularVelocity += body.InvI * PointCross(
			PointSub(point, body.Sweep.C),
			impulse,
		)
	}
}

func (body *Body) ApplyLinearImpulseToCenter(impulse Point, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate velocity if the body is sleeping
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.LinearVelocity.OperatorPlusInplace(PointMulScalar(body.InvMass, impulse))
	}
}

func (body *Body) ApplyAngularImpulse(impulse float64, wake bool) {
	if body.Type != BodyTypeDynamicBody {
		return
	}

	if wake && (body.Flags&BodyFlagAwake) == 0 {
		body.SetAwake(true)
	}

	// Don't accumulate velocity if the body is sleeping
	if (body.Flags & BodyFlagAwake) != 0x0000 {
		body.AngularVelocity += body.InvI * impulse
	}
}

func (body *Body) SynchronizeTransform() {
	body.Xf.Q.Set(body.Sweep.A)
	body.Xf.P = PointSub(body.Sweep.C, RotPointMul(body.Xf.Q, body.Sweep.LocalCenter))
}

func (body *Body) Advance(alpha float64) {
	// Advance to the new safe time. This doesn't sync the broad-phase.
	body.Sweep.Advance(alpha)
	body.Sweep.C = body.Sweep.C0
	body.Sweep.A = body.Sweep.A0
	body.Xf.Q.Set(body.Sweep.A)
	body.Xf.P = PointSub(body.Sweep.C, RotPointMul(body.Xf.Q, body.Sweep.LocalCenter))
}

func (body Body) GetWorld() *World {
	return body.World
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Body.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func NewBody(bd *BodyDef, world *World) *Body {
	body := &Body{}

	body.Flags = 0

	if bd.Bullet {
		body.Flags |= BodyFlagBullet
	}

	if bd.FixedRotation {
		body.Flags |= BodyFlagFixedRotation
	}

	if bd.AllowSleep {
		body.Flags |= BodyFlagAutoSleep
	}

	if bd.Awake {
		body.Flags |= BodyFlagAwake
	}

	if bd.Active {
		body.Flags |= BodyFlagActive
	}

	body.World = world

	body.Xf.P = bd.Position
	body.Xf.Q.Set(bd.Angle)

	body.Sweep.LocalCenter.SetZero()
	body.Sweep.C0 = body.Xf.P
	body.Sweep.C = body.Xf.P
	body.Sweep.A0 = bd.Angle
	body.Sweep.A = bd.Angle
	body.Sweep.Alpha0 = 0.0

	body.ContactList = nil
	body.Prev = nil
	body.Next = nil

	body.LinearVelocity = bd.LinearVelocity
	body.AngularVelocity = bd.AngularVelocity

	body.LinearDamping = bd.LinearDamping
	body.AngularDamping = bd.AngularDamping
	body.GravityScale = bd.GravityScale

	body.Force.SetZero()
	body.Torque = 0.0

	body.SleepTime = 0.0

	body.Type = bd.Type

	if body.Type == BodyTypeDynamicBody {
		body.Mass = 1.0
		body.InvMass = 1.0
	} else {
		body.Mass = 0.0
		body.InvMass = 0.0
	}

	body.I = 0.0
	body.InvI = 0.0

	body.FixtureList = nil
	body.FixtureCount = 0

	return body
}

func (body *Body) SetType(bodytype BodyType) {

	if body.World.IsLocked() == true {
		return
	}

	if body.Type == bodytype {
		return
	}

	body.Type = bodytype

	body.ResetMassData()

	if body.Type == BodyTypeStaticBody {
		body.LinearVelocity.SetZero()
		body.AngularVelocity = 0.0
		body.Sweep.A0 = body.Sweep.A
		body.Sweep.C0 = body.Sweep.C
		body.SynchronizeFixtures()
	}

	body.SetAwake(true)

	body.Force.SetZero()
	body.Torque = 0.0

	// Delete the attached contacts.
	ce := body.ContactList
	for ce != nil {
		ce0 := ce
		ce = ce.Next
		body.World.ContactManager.Destroy(ce0.Contact)
	}

	body.ContactList = nil

	// Touch the proxies so that new contacts will be created (when appropriate)
	broadPhase := body.World.ContactManager.BroadPhase
	for f := body.FixtureList; f != nil; f = f.Next {
		proxyCount := f.ProxyCount
		for i := 0; i < proxyCount; i++ {
			broadPhase.TouchProxy(f.Proxies[i].ProxyId)
		}
	}
}

func (body *Body) CreateFixtureFromDef(def *FixtureDef) *Fixture {

	if body.World.IsLocked() == true {
		return nil
	}

	fixture := NewFixture()
	fixture.Create(body, def)

	if (body.Flags & BodyFlagActive) != 0x0000 {
		broadPhase := &body.World.ContactManager.BroadPhase
		fixture.CreateProxies(broadPhase, body.Xf)
	}

	fixture.Next = body.FixtureList
	body.FixtureList = fixture
	body.FixtureCount++

	fixture.Body = body

	// Adjust mass properties if needed.
	if fixture.Density > 0.0 {
		body.ResetMassData()
	}

	// Let the world know we have a new fixture. This will cause new contacts
	// to be created at the beginning of the next time step.
	body.World.Flags |= WorldFlagNewFixture

	return fixture
}

func (body *Body) CreateFixture(shape ShapeInterface, density float64) *Fixture {

	def := MakeFixtureDef()
	def.Shape = shape
	def.Density = density

	return body.CreateFixtureFromDef(&def)
}

func (body *Body) DestroyFixture(fixture *Fixture) {

	if fixture == nil {
		return
	}

	if body.World.IsLocked() == true {
		return
	}

	// Remove the fixture from this body's singly linked list.
	node := &body.FixtureList
	for *node != nil {
		if *node == fixture {
			*node = fixture.Next
			break
		}

		node = &(*node).Next
	}

	// You tried to remove a shape that is not attached to this body.

	// Destroy any contacts associated with the fixture.
	edge := body.ContactList
	for edge != nil {
		c := edge.Contact
		edge = edge.Next

		fixtureA := c.Data().GetFixtureA()
		fixtureB := c.Data().GetFixtureB()

		if fixture == fixtureA || fixture == fixtureB {
			// This destroys the contact and removes it from
			// this body's contact list.
			body.World.ContactManager.Destroy(c)
		}
	}

	if (body.Flags & BodyFlagActive) != 0x0000 {
		broadPhase := &body.World.ContactManager.BroadPhase
		fixture.DestroyProxies(broadPhase)
	}

	fixture.Body = nil
	fixture.Next = nil
	fixture.Destroy()

	body.FixtureCount--

	// Reset the mass data.
	body.ResetMassData()
}

func (body *Body) ResetMassData() {

	// Compute mass data from shapes. Each shape has its own density.
	body.Mass = 0.0
	body.InvMass = 0.0
	body.I = 0.0
	body.InvI = 0.0
	body.Sweep.LocalCenter.SetZero()

	// Static and kinematic bodies have zero mass.
	if body.Type == BodyTypeStaticBody || body.Type == BodyTypeKinematicBody {
		body.Sweep.C0 = body.Xf.P
		body.Sweep.C = body.Xf.P
		body.Sweep.A0 = body.Sweep.A
		return
	}

	// Accumulate mass over all fixtures.
	localCenter := Point{}
	for f := body.FixtureList; f != nil; f = f.Next {
		if f.Density == 0.0 {
			continue
		}

		massData := NewMassData()
		f.GetMassData(massData)
		body.Mass += massData.Mass
		localCenter.OperatorPlusInplace(PointMulScalar(massData.Mass, massData.Center))
		body.I += massData.I
	}

	// Compute center of mass.
	if body.Mass > 0.0 {
		body.InvMass = 1.0 / body.Mass
		localCenter.OperatorScalarMulInplace(body.InvMass)
	} else {
		// Force all dynamic bodies to have a positive mass.
		body.Mass = 1.0
		body.InvMass = 1.0
	}

	if body.I > 0.0 && (body.Flags&BodyFlagFixedRotation) == 0 {
		// Center the inertia about the center of mass.
		body.I -= body.Mass * PointDot(localCenter, localCenter)
		body.InvI = 1.0 / body.I

	} else {
		body.I = 0.0
		body.InvI = 0.0
	}

	// Move center of mass.
	oldCenter := body.Sweep.C
	body.Sweep.LocalCenter = localCenter
	body.Sweep.C0 = TransformPointMul(body.Xf, body.Sweep.LocalCenter)
	body.Sweep.C = TransformPointMul(body.Xf, body.Sweep.LocalCenter)

	// Update center of mass velocity.
	body.LinearVelocity.OperatorPlusInplace(PointCrossScalarVector(
		body.AngularVelocity,
		PointSub(body.Sweep.C, oldCenter),
	))
}

func (body *Body) SetMassData(massData *MassData) {

	if body.World.IsLocked() == true {
		return
	}

	if body.Type != BodyTypeDynamicBody {
		return
	}

	body.InvMass = 0.0
	body.I = 0.0
	body.InvI = 0.0

	body.Mass = massData.Mass
	if body.Mass <= 0.0 {
		body.Mass = 1.0
	}

	body.InvMass = 1.0 / body.Mass

	if massData.I > 0.0 && (body.Flags&BodyFlagFixedRotation) == 0 {
		body.I = massData.I - body.Mass*PointDot(massData.Center, massData.Center)
		body.InvI = 1.0 / body.I
	}

	// Move center of mass.
	oldCenter := body.Sweep.C
	body.Sweep.LocalCenter = massData.Center
	body.Sweep.C0 = TransformPointMul(body.Xf, body.Sweep.LocalCenter)
	body.Sweep.C = TransformPointMul(body.Xf, body.Sweep.LocalCenter)

	// Update center of mass velocity.
	body.LinearVelocity.OperatorPlusInplace(
		PointCrossScalarVector(
			body.AngularVelocity,
			PointSub(body.Sweep.C, oldCenter),
		),
	)
}

func (body Body) ShouldCollide(other *Body) bool {

	// At least one body should be dynamic.
	if body.Type != BodyTypeDynamicBody && other.Type != BodyTypeDynamicBody {
		return false
	}

	return true
}

func (body *Body) SetTransform(position Point, angle float64) {

	if body.World.IsLocked() == true {
		return
	}

	body.Xf.Q.Set(angle)
	body.Xf.P = position

	body.Sweep.C = TransformPointMul(body.Xf, body.Sweep.LocalCenter)
	body.Sweep.A = angle

	body.Sweep.C0 = body.Sweep.C
	body.Sweep.A0 = angle

	broadPhase := &body.World.ContactManager.BroadPhase
	for f := body.FixtureList; f != nil; f = f.Next {
		f.Synchronize(broadPhase, body.Xf, body.Xf)
	}
}

func (body *Body) SynchronizeFixtures() {
	xf1 := Transform{}
	xf1.Q.Set(body.Sweep.A0)
	xf1.P = PointSub(body.Sweep.C0, RotPointMul(xf1.Q, body.Sweep.LocalCenter))

	broadPhase := &body.World.ContactManager.BroadPhase
	for f := body.FixtureList; f != nil; f = f.Next {
		f.Synchronize(broadPhase, xf1, body.Xf)
	}
}

func (body *Body) SetActive(flag bool) {

	if flag == body.IsActive() {
		return
	}

	if flag {
		body.Flags |= BodyFlagActive

		// Create all proxies.
		broadPhase := &body.World.ContactManager.BroadPhase
		for f := body.FixtureList; f != nil; f = f.Next {
			f.CreateProxies(broadPhase, body.Xf)
		}

		// Contacts are created the next time step.
	} else {
		body.Flags &= ^BodyFlagActive

		// Destroy all proxies.
		broadPhase := &body.World.ContactManager.BroadPhase
		for f := body.FixtureList; f != nil; f = f.Next {
			f.DestroyProxies(broadPhase)
		}

		// Destroy the attached contacts.
		ce := body.ContactList
		for ce != nil {
			ce0 := ce
			ce = ce.Next
			body.World.ContactManager.Destroy(ce0.Contact)
		}

		body.ContactList = nil
	}
}

func (body *Body) SetFixedRotation(flag bool) {
	status := (body.Flags & BodyFlagFixedRotation) == BodyFlagFixedRotation

	if status == flag {
		return
	}

	if flag {
		body.Flags |= BodyFlagFixedRotation
	} else {
		body.Flags &= ^BodyFlagFixedRotation
	}

	body.AngularVelocity = 0.0

	body.ResetMassData()
}
