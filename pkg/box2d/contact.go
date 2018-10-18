package box2d

import (
	"math"
)

/// Friction mixing law. The idea is to allow either fixture to drive the friction to zero.
/// For example, anything slides on ice.
func MixFriction(friction1, friction2 float64) float64 {
	return math.Sqrt(friction1 * friction2)
}

/// Restitution mixing law. The idea is allow for anything to bounce off an inelastic surface.
/// For example, a superball bounces on anything.
func MixRestitution(restitution1, restitution2 float64) float64 {
	if restitution1 > restitution2 {
		return restitution1
	}

	return restitution2
}

type ContactInterface interface {
	Data() *Contact
	Evaluate(*Manifold, Transform, Transform)
}

type ContactCreateFunc func(fixtureA *Fixture, indexA int, fixtureB *Fixture, indexB int) ContactInterface // returned contact should be a pointer

type ContactRegister struct {
	CreateFunc ContactCreateFunc
	Primary    bool
}

/// A contact edge is used to connect bodies and contacts together
/// in a contact graph where each body is a node and each contact
/// is an edge. A contact edge belongs to a doubly linked list
/// maintained in each attached body. Each contact has two contact
/// nodes, one for each attached body.
type ContactEdge struct {
	Other   *Body            ///< provides quick access to the other body attached.
	Contact ContactInterface ///< the contact
	Prev    *ContactEdge     ///< the previous contact edge in the body's contact list
	Next    *ContactEdge     ///< the next contact edge in the body's contact list
}

func NewContactEdge() *ContactEdge {
	return &ContactEdge{}
}

type ContactFlag uint32

const (
	// ContactFlagIsland is used when crawling contact graph when forming islands.
	ContactFlagIsland ContactFlag = 0x0001
	// ContactFlagTouching is set when the shapes are touching.
	ContactFlagTouching ContactFlag = 0x0002
	// ContactFlagEnabled means this contact can be disabled (by user)
	ContactFlagEnabled ContactFlag = 0x0004
	// ContactFlagFilter means this contact needs filtering because a fixture filter was changed.
	ContactFlagFilter ContactFlag = 0x0008
	// ContactFlagBulletHit means this bullet contact had a TOI event
	ContactFlagBulletHit ContactFlag = 0x0010
	// ContactFlagToi means this contact has a valid TOI in m_toi
	ContactFlagToi ContactFlag = 0x0020
)

// /// The class manages contact between two shapes. A contact exists for each overlapping
// /// AABB in the broad-phase (except if filtered). Therefore a contact object may exist
// /// that has no contact points.
var s_registers [][]ContactRegister
var s_initialized = false

type Contact struct {
	Flags ContactFlag

	// World pool and list pointers.
	Prev ContactInterface //should be backed by a pointer
	Next ContactInterface //should be backed by a pointer

	// Nodes for connecting bodies.
	NodeA *ContactEdge
	NodeB *ContactEdge

	FixtureA *Fixture
	FixtureB *Fixture

	IndexA int
	IndexB int

	Manifold *Manifold

	ToiCount     int
	Toi          float64
	Friction     float64
	Restitution  float64
	TangentSpeed float64
}

func (contact Contact) GetFlags() ContactFlag {
	return contact.Flags
}

func (contact *Contact) SetFlags(flags ContactFlag) {
	contact.Flags = flags
}

func (contact Contact) GetPrev() ContactInterface {
	return contact.Prev
}

func (contact *Contact) SetPrev(prev ContactInterface) {
	contact.Prev = prev
}

func (contact Contact) GetNext() ContactInterface {
	return contact.Next
}

func (contact *Contact) SetNext(next ContactInterface) {
	contact.Next = next
}

func (contact Contact) GetNodeA() *ContactEdge {
	return contact.NodeA
}

func (contact *Contact) SetNodeA(node *ContactEdge) {
	contact.NodeA = node
}

func (contact Contact) GetNodeB() *ContactEdge {
	return contact.NodeB
}

func (contact *Contact) SetNodeB(node *ContactEdge) {
	contact.NodeB = node
}

func (contact Contact) GetFixtureA() *Fixture {
	return contact.FixtureA
}

func (contact *Contact) SetFixtureA(fixture *Fixture) {
	contact.FixtureA = fixture
}

func (contact Contact) GetFixtureB() *Fixture {
	return contact.FixtureB
}

func (contact *Contact) SetFixtureB(fixture *Fixture) {
	contact.FixtureB = fixture
}

func (contact Contact) GetChildIndexA() int {
	return contact.IndexA
}

func (contact *Contact) SetChildIndexA(index int) {
	contact.IndexA = index
}

func (contact Contact) GetChildIndexB() int {
	return contact.IndexB
}

func (contact *Contact) SetChildIndexB(index int) {
	contact.IndexB = index
}

func (contact Contact) GetManifold() *Manifold {
	return contact.Manifold
}

func (contact *Contact) SetManifold(manifold *Manifold) {
	contact.Manifold = manifold
}

func (contact Contact) GetTOICount() int {
	return contact.ToiCount
}

func (contact *Contact) SetTOICount(toiCount int) {
	contact.ToiCount = toiCount
}

func (contact Contact) GetTOI() float64 {
	return contact.Toi
}

func (contact *Contact) SetTOI(toi float64) {
	contact.Toi = toi
}

func (contact Contact) GetFriction() float64 {
	return contact.Friction
}

func (contact *Contact) SetFriction(friction float64) {
	contact.Friction = friction
}

func (contact *Contact) ResetFriction() {
	contact.Friction = MixFriction(contact.FixtureA.Friction, contact.FixtureB.Friction)
}

func (contact Contact) GetRestitution() float64 {
	return contact.Restitution
}

func (contact *Contact) SetRestitution(restitution float64) {
	contact.Restitution = restitution
}

func (contact *Contact) ResetRestitution() {
	contact.Restitution = MixRestitution(contact.FixtureA.Restitution, contact.FixtureB.Restitution)
}

func (contact Contact) GetTangentSpeed() float64 {
	return contact.TangentSpeed
}

func (contact *Contact) SetTangentSpeed(speed float64) {
	contact.TangentSpeed = speed
}

func (contact Contact) GetWorldManifold(worldManifold *WorldManifold) {
	bodyA := contact.FixtureA.GetBody()
	bodyB := contact.FixtureB.GetBody()
	shapeA := contact.FixtureA.GetShape()
	shapeB := contact.FixtureB.GetShape()

	worldManifold.Initialize(contact.Manifold, bodyA.GetTransform(), shapeA.GetRadius(), bodyB.GetTransform(), shapeB.GetRadius())
}

func (contact *Contact) SetEnabled(flag bool) {
	if flag {
		contact.Flags |= ContactFlagEnabled
	} else {
		contact.Flags &= ^ContactFlagEnabled
	}
}

func (contact Contact) IsEnabled() bool {
	return (contact.Flags & ContactFlagEnabled) == ContactFlagEnabled
}

func (contact Contact) IsTouching() bool {
	return (contact.Flags & ContactFlagTouching) == ContactFlagTouching
}

func (contact *Contact) FlagForFiltering() {
	contact.Flags |= ContactFlagFilter
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Contact.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func ContactInitializeRegisters() {
	s_registers = make([][]ContactRegister, ShapeTypeTypeCount)
	for i := 0; i < int(ShapeTypeTypeCount); i++ {
		s_registers[i] = make([]ContactRegister, ShapeTypeTypeCount)
	}

	AddType(PolygonContact_Create, ShapeTypePolygon, ShapeTypePolygon)
	AddType(EdgeAndPolygonContact_Create, ShapeTypeEdge, ShapeTypePolygon)
}

func AddType(createFunc ContactCreateFunc, type1 ShapeType, type2 ShapeType) {
	s_registers[type1][type2].CreateFunc = createFunc
	s_registers[type1][type2].Primary = true

	if type1 != type2 {
		s_registers[type2][type1].CreateFunc = createFunc
		s_registers[type2][type1].Primary = false
	}
}

func ContactFactory(fixtureA *Fixture, indexA int, fixtureB *Fixture, indexB int) ContactInterface { // returned contact should be a pointer

	if s_initialized == false {
		ContactInitializeRegisters()
		s_initialized = true
	}

	type1 := fixtureA.GetType()
	type2 := fixtureB.GetType()

	createFunc := s_registers[type1][type2].CreateFunc
	if createFunc != nil {
		if s_registers[type1][type2].Primary {
			return createFunc(fixtureA, indexA, fixtureB, indexB)
		} else {
			return createFunc(fixtureB, indexB, fixtureA, indexA)
		}
	}

	return nil
}

func ContactDestroy(contact ContactInterface) {
	contactData := contact.Data()
	fixtureA := contactData.FixtureA
	fixtureB := contactData.FixtureB

	if contactData.GetManifold().PointCount > 0 && fixtureA.IsSensor == false && fixtureB.IsSensor == false {
		fixtureA.GetBody().SetAwake(true)
		fixtureB.GetBody().SetAwake(true)
	}
}

func MakeContact(fA *Fixture, indexA int, fB *Fixture, indexB int) Contact {

	contact := Contact{}
	contact.Flags = ContactFlagEnabled

	contact.FixtureA = fA
	contact.FixtureB = fB

	contact.IndexA = indexA
	contact.IndexB = indexB

	contact.Manifold = &Manifold{}
	contact.Manifold.PointCount = 0

	contact.Prev = nil
	contact.Next = nil

	contact.NodeA = NewContactEdge()

	contact.NodeA.Contact = nil
	contact.NodeA.Prev = nil
	contact.NodeA.Next = nil
	contact.NodeA.Other = nil

	contact.NodeB = NewContactEdge()

	contact.NodeB.Contact = nil
	contact.NodeB.Prev = nil
	contact.NodeB.Next = nil
	contact.NodeB.Other = nil

	contact.ToiCount = 0

	contact.Friction = MixFriction(contact.FixtureA.Friction, contact.FixtureB.Friction)
	contact.Restitution = MixRestitution(contact.FixtureA.Restitution, contact.FixtureB.Restitution)

	contact.TangentSpeed = 0.0

	return contact
}

// Update the contact manifold and touching status.
// Note: do not assume the fixture AABBs are overlapping or are valid.
func ContactUpdate(contact ContactInterface, listener ContactListenerInterface) {
	contactData := contact.Data()
	oldManifold := *contactData.GetManifold()

	// Re-enable this contact.
	contactData.SetFlags(contactData.GetFlags() | ContactFlagEnabled)

	touching := false
	wasTouching := (contactData.GetFlags() & ContactFlagTouching) == ContactFlagTouching

	sensorA := contactData.GetFixtureA().IsSensor
	sensorB := contactData.GetFixtureB().IsSensor
	sensor := sensorA || sensorB

	bodyA := contactData.GetFixtureA().GetBody()
	bodyB := contactData.GetFixtureB().GetBody()
	xfA := bodyA.GetTransform()
	xfB := bodyB.GetTransform()

	// Is this contact a sensor?
	if sensor {
		shapeA := contactData.GetFixtureA().GetShape()
		shapeB := contactData.GetFixtureB().GetShape()
		touching = TestOverlapShapes(shapeA, contactData.GetChildIndexA(), shapeB, contactData.GetChildIndexB(), xfA, xfB)

		// Sensors don't generate manifolds.
		contactData.GetManifold().PointCount = 0
	} else {
		// ContactInterface is extended by specialized contact structs and mentionned by ContactInterface but not implemented on specialized structs
		contact.Evaluate(contactData.GetManifold(), xfA, xfB) // should be evaluated on specialisations of contact (like CircleContact)
		touching = contactData.GetManifold().PointCount > 0

		// Match old contact ids to new contact ids and copy the
		// stored impulses to warm start the solver.
		for i := 0; i < contactData.GetManifold().PointCount; i++ {
			mp2 := &contactData.GetManifold().Points[i]
			mp2.NormalImpulse = 0.0
			mp2.TangentImpulse = 0.0
			id2 := mp2.Id

			for j := 0; j < oldManifold.PointCount; j++ {
				mp1 := &oldManifold.Points[j]

				if mp1.Id.Key() == id2.Key() {
					mp2.NormalImpulse = mp1.NormalImpulse
					mp2.TangentImpulse = mp1.TangentImpulse
					break
				}
			}
		}

		if touching != wasTouching {
			bodyA.SetAwake(true)
			bodyB.SetAwake(true)
		}
	}

	if touching {
		contactData.SetFlags(contactData.GetFlags() | ContactFlagTouching)
	} else {
		contactData.SetFlags(contactData.GetFlags() & ^ContactFlagTouching)
	}

	if wasTouching == false && touching == true && listener != nil {
		listener.BeginContact(contact)
	}

	if wasTouching == true && touching == false && listener != nil {
		listener.EndContact(contact)
	}

	if sensor == false && touching && listener != nil {
		listener.PreSolve(contact, oldManifold)
	}
}
