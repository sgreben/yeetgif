package box2d

type EdgeAndPolygonContact struct{ Contact }

func (contact *EdgeAndPolygonContact) Data() *Contact {
	return &contact.Contact
}

func EdgeAndPolygonContact_Create(fixtureA *Fixture, indexA int, fixtureB *Fixture, indexB int) ContactInterface {
	res := &EdgeAndPolygonContact{
		Contact: MakeContact(fixtureA, 0, fixtureB, 0),
	}
	return res
}

func (contact *EdgeAndPolygonContact) Evaluate(manifold *Manifold, xfA Transform, xfB Transform) {
	CollideEdgeAndPolygon(
		manifold,
		contact.GetFixtureA().GetShape().(*EdgeShape), xfA,
		contact.GetFixtureB().GetShape().(*PolygonShape), xfB,
	)
}
