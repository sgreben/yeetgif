package box2d

/// This holds contact filtering data.
type Filter struct {
	/// The collision category bits. Normally you would just set one bit.
	CategoryBits uint16

	/// The collision mask bits. This states the categories that this
	/// shape would accept for collision.
	MaskBits uint16

	/// Collision groups allow a certain group of objects to never collide (negative)
	/// or always collide (positive). Zero means no collision group. Non-zero group
	/// filtering always wins against the mask bits.
	GroupIndex int16
}

func MakeFilter() Filter {
	return Filter{
		CategoryBits: 0x0001,
		MaskBits:     0xFFFF,
		GroupIndex:   0,
	}
}

/// A fixture definition is used to create a fixture. This class defines an
/// abstract fixture definition. You can reuse fixture definitions safely.
type FixtureDef struct {

	/// The shape, this must be set. The shape will be cloned, so you
	/// can create the shape on the stack.
	Shape ShapeInterface

	/// Use this to store application specific fixture data.
	UserData interface{}

	/// The friction coefficient, usually in the range [0,1].
	Friction float64

	/// The restitution (elasticity) usually in the range [0,1].
	Restitution float64

	/// The density, usually in kg/m^2.
	Density float64

	/// A sensor shape collects contact information but never generates a collision
	/// response.
	IsSensor bool

	/// Contact filtering data.
	Filter Filter
}

/// The constructor sets the default fixture definition values.
func MakeFixtureDef() FixtureDef {
	return FixtureDef{
		Shape:       nil,
		UserData:    nil,
		Friction:    0.2,
		Restitution: 0.0,
		Density:     0.0,
		IsSensor:    false,
	}
}

/// This proxy is used internally to connect fixtures to the broad-phase.
type FixtureProxy struct {
	Aabb       AABB
	Fixture    *Fixture
	ChildIndex int
	ProxyId    int
}

// /// A fixture is used to attach a shape to a body for collision detection. A fixture
// /// inherits its transform from its parent. Fixtures hold additional non-geometric data
// /// such as friction, collision filters, etc.
// /// Fixtures are created via b2Body::CreateFixture.
// /// @warning you cannot reuse fixtures.
type Fixture struct {
	Density float64

	Next *Fixture
	Body *Body

	Shape ShapeInterface

	Friction    float64
	Restitution float64

	Proxies    []FixtureProxy
	ProxyCount int

	Filter Filter

	IsSensor bool

	UserData interface{}
}

func NewFixture() *Fixture {
	return &Fixture{
		Next:   nil,
		Body:   nil,
		Filter: MakeFilter(),
	}
}

func (fix Fixture) GetType() ShapeType {
	return fix.Shape.GetType()
}

func (fix Fixture) GetShape() ShapeInterface {
	return fix.Shape
}

func (fix Fixture) GetFilterData() Filter {
	return fix.Filter
}

func (fix Fixture) GetUserData() interface{} {
	return fix.UserData
}

func (fix *Fixture) SetUserData(data interface{}) {
	fix.UserData = data
}

func (fix Fixture) GetBody() *Body {
	return fix.Body
}

func (fix Fixture) GetNext() *Fixture {
	return fix.Next
}

func (fix *Fixture) SetDensity(density float64) {
	fix.Density = density
}

func (fix Fixture) GetDensity() float64 {
	return fix.Density
}

func (fix Fixture) GetFriction() float64 {
	return fix.Friction
}

func (fix *Fixture) SetFriction(friction float64) {
	fix.Friction = friction
}

func (fix Fixture) GetRestitution() float64 {
	return fix.Restitution
}

func (fix *Fixture) SetRestitution(restitution float64) {
	fix.Restitution = restitution
}

func (fix Fixture) TestPoint(p Point) bool {
	return fix.Shape.TestPoint(fix.Body.GetTransform(), p)
}

func (fix Fixture) RayCast(output *RayCastOutput, input RayCastInput, childIndex int) bool {
	return fix.Shape.RayCast(output, input, fix.Body.GetTransform(), childIndex)
}

func (fix Fixture) GetMassData(massData *MassData) {
	fix.Shape.ComputeMass(massData, fix.Density)
}

func (fix Fixture) GetAABB(childIndex int) AABB {
	return fix.Proxies[childIndex].Aabb
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Fixture.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func MakeFixture() Fixture {
	return Fixture{
		UserData:   nil,
		Body:       nil,
		Next:       nil,
		Proxies:    nil,
		ProxyCount: 0,
		Shape:      nil,
		Density:    0.0,
	}
}

func (fix *Fixture) Create(body *Body, def *FixtureDef) {
	fix.UserData = def.UserData
	fix.Friction = def.Friction
	fix.Restitution = def.Restitution

	fix.Body = body
	fix.Next = nil

	fix.Filter = def.Filter

	fix.IsSensor = def.IsSensor

	fix.Shape = def.Shape.Clone()

	// Reserve proxy space
	childCount := fix.Shape.GetChildCount()
	fix.Proxies = make([]FixtureProxy, childCount)

	for i := 0; i < childCount; i++ {
		fix.Proxies[i].Fixture = nil
		fix.Proxies[i].ProxyId = nullProxyID
	}
	fix.ProxyCount = 0

	fix.Density = def.Density
}

func (fix *Fixture) Destroy() {
	// The proxies must be destroyed before calling this.
	fix.Proxies = nil
	fix.Shape = nil
}

func (fix *Fixture) CreateProxies(broadPhase *BroadPhase, xf Transform) {

	// Create proxies in the broad-phase.
	fix.ProxyCount = fix.Shape.GetChildCount()

	for i := 0; i < fix.ProxyCount; i++ {
		proxy := &fix.Proxies[i]
		fix.Shape.ComputeAABB(&proxy.Aabb, xf, i)
		proxy.ProxyId = broadPhase.CreateProxy(proxy.Aabb, proxy)
		proxy.Fixture = fix
		proxy.ChildIndex = i
	}
}

func (fix *Fixture) DestroyProxies(broadPhase *BroadPhase) {
	// Destroy proxies in the broad-phase.
	for i := 0; i < fix.ProxyCount; i++ {
		proxy := &fix.Proxies[i]
		broadPhase.DestroyProxy(proxy.ProxyId)
		proxy.ProxyId = nullProxyID
	}

	fix.ProxyCount = 0
}

func (fix *Fixture) Synchronize(broadPhase *BroadPhase, transform1 Transform, transform2 Transform) {

	if fix.ProxyCount == 0 {
		return
	}

	for i := 0; i < fix.ProxyCount; i++ {

		proxy := &fix.Proxies[i]

		// Compute an AABB that covers the swept shape (may miss some rotation effect).
		aabb1 := AABB{}
		aabb2 := AABB{}
		fix.Shape.ComputeAABB(&aabb1, transform1, proxy.ChildIndex)
		fix.Shape.ComputeAABB(&aabb2, transform2, proxy.ChildIndex)

		proxy.Aabb.CombineTwoInPlace(aabb1, aabb2)

		displacement := PointSub(transform2.P, transform1.P)

		broadPhase.MoveProxy(proxy.ProxyId, proxy.Aabb, displacement)
	}
}

func (fix *Fixture) SetFilterData(filter Filter) {
	fix.Filter = filter
	fix.Refilter()
}

func (fix *Fixture) Refilter() {

	if fix.Body == nil {
		return
	}

	// Flag associated contacts for filtering.
	edge := fix.Body.GetContactList()
	for edge != nil {
		contact := edge.Contact.Data()
		fixtureA := contact.GetFixtureA()
		fixtureB := contact.GetFixtureB()
		if fixtureA == fix || fixtureB == fix {
			contact.FlagForFiltering()
		}

		edge = edge.Next
	}

	world := fix.Body.GetWorld()

	if world == nil {
		return
	}

	// Touch each proxy so that new pairs may be created
	broadPhase := &world.ContactManager.BroadPhase
	for i := 0; i < fix.ProxyCount; i++ {
		broadPhase.TouchProxy(fix.Proxies[i].ProxyId)
	}
}

func (fix *Fixture) SetSensor(sensor bool) {
	if sensor != fix.IsSensor {
		fix.Body.SetAwake(true)
		fix.IsSensor = sensor
	}
}
