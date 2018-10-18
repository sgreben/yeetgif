package box2d

type ContactManager struct {
	BroadPhase      BroadPhase
	ContactList     ContactInterface
	ContactCount    int
	ContactFilter   ContactFilterInterface
	ContactListener ContactListenerInterface
}

var b2_defaultFilter ContactFilterInterface
var b2_defaultListener ContactListenerInterface

func MakeContactManager() ContactManager {
	return ContactManager{
		BroadPhase:      MakeBroadPhase(),
		ContactList:     nil,
		ContactCount:    0,
		ContactFilter:   b2_defaultFilter,
		ContactListener: b2_defaultListener,
	}
}

func NewContactManager() *ContactManager {
	res := MakeContactManager()
	return &res
}

func (mgr *ContactManager) Destroy(contact ContactInterface) {
	c := contact.Data()
	fixtureA := c.GetFixtureA()
	fixtureB := c.GetFixtureB()
	bodyA := fixtureA.GetBody()
	bodyB := fixtureB.GetBody()

	if mgr.ContactListener != nil && c.IsTouching() {
		mgr.ContactListener.EndContact(contact)
	}

	// Remove from the world.
	if c.GetPrev() != nil {
		c.GetPrev().Data().SetNext(c.GetNext())
	}

	if c.GetNext() != nil {
		c.GetNext().Data().SetPrev(c.GetPrev())
	}

	if contact == mgr.ContactList {
		mgr.ContactList = c.GetNext()
	}

	// Remove from body 1
	if c.GetNodeA().Prev != nil {
		c.GetNodeA().Prev.Next = c.GetNodeA().Next
	}

	if c.GetNodeA().Next != nil {
		c.GetNodeA().Next.Prev = c.GetNodeA().Prev
	}

	if c.GetNodeA() == bodyA.ContactList {
		bodyA.ContactList = c.GetNodeA().Next
	}

	// Remove from body 2
	if c.GetNodeB().Prev != nil {
		c.GetNodeB().Prev.Next = c.GetNodeB().Next
	}

	if c.GetNodeB().Next != nil {
		c.GetNodeB().Next.Prev = c.GetNodeB().Prev
	}

	if c.GetNodeB() == bodyB.ContactList {
		bodyB.ContactList = c.GetNodeB().Next
	}

	// Call the factory.
	ContactDestroy(contact)
	mgr.ContactCount--
}

// This is the top level collision call for the time step. Here
// all the narrow phase collision is processed for the world
// contact list.
func (mgr *ContactManager) Collide() {
	// Update awake contacts.
	c := mgr.ContactList

	for c != nil {
		contactData := c.Data()
		fixtureA := contactData.GetFixtureA()
		fixtureB := contactData.GetFixtureB()
		indexA := contactData.GetChildIndexA()
		indexB := contactData.GetChildIndexB()
		bodyA := fixtureA.GetBody()
		bodyB := fixtureB.GetBody()

		// Is this contact flagged for filtering?
		if (contactData.GetFlags() & ContactFlagFilter) != 0x0000 {
			// Should these bodies collide?
			if bodyB.ShouldCollide(bodyA) == false {
				cNuke := c
				c = cNuke.Data().GetNext()
				mgr.Destroy(cNuke)
				continue
			}

			// Check user filtering.
			if mgr.ContactFilter != nil && mgr.ContactFilter.ShouldCollide(fixtureA, fixtureB) == false {
				cNuke := c
				c = cNuke.Data().GetNext()
				mgr.Destroy(cNuke)
				continue
			}

			// Clear the filtering flag.
			contactData.SetFlags(contactData.GetFlags() & ^ContactFlagFilter)
		}

		activeA := bodyA.IsAwake() && bodyA.Type != BodyTypeStaticBody
		activeB := bodyB.IsAwake() && bodyB.Type != BodyTypeStaticBody

		// At least one body must be awake and it must be dynamic or kinematic.
		if activeA == false && activeB == false {
			c = contactData.GetNext()
			continue
		}

		proxyIdA := fixtureA.Proxies[indexA].ProxyId
		proxyIdB := fixtureB.Proxies[indexB].ProxyId
		overlap := mgr.BroadPhase.TestOverlap(proxyIdA, proxyIdB)

		// Here we destroy contacts that cease to overlap in the broad-phase.
		if overlap == false {
			cNuke := c
			c = cNuke.Data().GetNext()
			mgr.Destroy(cNuke)
			continue
		}

		// The contact persists.
		ContactUpdate(c, mgr.ContactListener)
		c = contactData.GetNext()
	}
}

func (mgr *ContactManager) FindNewContacts() {
	mgr.BroadPhase.UpdatePairs(mgr.AddPair)
}

func (mgr *ContactManager) AddPair(proxyUserDataA interface{}, proxyUserDataB interface{}) {

	proxyA := proxyUserDataA.(*FixtureProxy)
	proxyB := proxyUserDataB.(*FixtureProxy)

	fixtureA := proxyA.Fixture
	fixtureB := proxyB.Fixture

	indexA := proxyA.ChildIndex
	indexB := proxyB.ChildIndex

	bodyA := fixtureA.GetBody()
	bodyB := fixtureB.GetBody()

	// Are the fixtures on the same body?
	if bodyA == bodyB {
		return
	}

	// TODO_ERIN use a hash table to remove a potential bottleneck when both
	// bodies have a lot of contacts.
	// Does a contact already exist?
	edge := bodyB.GetContactList()
	for edge != nil {
		if edge.Other == bodyA {
			fA := edge.Contact.Data().GetFixtureA()
			fB := edge.Contact.Data().GetFixtureB()
			iA := edge.Contact.Data().GetChildIndexA()
			iB := edge.Contact.Data().GetChildIndexB()

			if fA == fixtureA && fB == fixtureB && iA == indexA && iB == indexB {
				// A contact already exists.
				return
			}

			if fA == fixtureB && fB == fixtureA && iA == indexB && iB == indexA {
				// A contact already exists.
				return
			}
		}

		edge = edge.Next
	}

	// Does a joint override collision? Is at least one body dynamic?
	if bodyB.ShouldCollide(bodyA) == false {
		return
	}

	// Check user filtering.
	if mgr.ContactFilter != nil && mgr.ContactFilter.ShouldCollide(fixtureA, fixtureB) == false {
		return
	}

	// Call the factory.
	c := ContactFactory(fixtureA, indexA, fixtureB, indexB)
	if c == nil {
		return
	}

	contactData := c.Data()
	// Contact creation may swap fixtures.
	fixtureA = contactData.GetFixtureA()
	fixtureB = contactData.GetFixtureB()
	indexA = contactData.GetChildIndexA()
	indexB = contactData.GetChildIndexB()
	bodyA = fixtureA.GetBody()
	bodyB = fixtureB.GetBody()

	// Insert into the world.
	contactData.SetPrev(nil)
	contactData.SetNext(mgr.ContactList)
	if mgr.ContactList != nil {
		mgr.ContactList.Data().SetPrev(c)
	}
	mgr.ContactList = c

	// Connect to island graph.

	// Connect to body A
	// fmt.Printf("getNode(): %p\n", c.GetNodeA())
	// fmt.Printf("getNode(): %p\n", c.GetNodeA())
	// fmt.Printf("getNode(): %p\n", c.GetNodeA())

	contactData.GetNodeA().Contact = c
	contactData.GetNodeA().Other = bodyB

	contactData.GetNodeA().Prev = nil
	contactData.GetNodeA().Next = bodyA.ContactList
	if bodyA.ContactList != nil {
		bodyA.ContactList.Prev = contactData.GetNodeA()
	}
	bodyA.ContactList = contactData.GetNodeA()

	// Connect to body B
	contactData.GetNodeB().Contact = c
	contactData.GetNodeB().Other = bodyA

	contactData.GetNodeB().Prev = nil
	contactData.GetNodeB().Next = bodyB.ContactList
	if bodyB.ContactList != nil {
		bodyB.ContactList.Prev = contactData.GetNodeB()
	}
	bodyB.ContactList = contactData.GetNodeB()

	// Wake up the bodies
	if !fixtureA.IsSensor && !fixtureB.IsSensor {
		bodyA.SetAwake(true)
		bodyB.SetAwake(true)
	}

	mgr.ContactCount++
}
