package box2d

type PolygonContact struct {
	Contact
}

func (contact *PolygonContact) Data() *Contact {
	return &contact.Contact
}

func PolygonContact_Create(fixtureA *Fixture, indexA int, fixtureB *Fixture, indexB int) ContactInterface {
	res := &PolygonContact{
		Contact: MakeContact(fixtureA, 0, fixtureB, 0),
	}

	return res
}

func (contact *PolygonContact) Evaluate(manifold *Manifold, xfA Transform, xfB Transform) {
	CollidePolygons(
		manifold,
		contact.GetFixtureA().GetShape().(*PolygonShape), xfA,
		contact.GetFixtureB().GetShape().(*PolygonShape), xfB,
	)
}
