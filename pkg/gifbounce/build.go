package gifbounce

import "github.com/sgreben/yeetgif/pkg/box2d"

func (p *Params) New() *World {
	w := box2d.MakeWorld(box2d.Point{X: 0, Y: -p.Gravity})
	w.ContinuousPhysics = true
	w.AllowSleep = true
	world := &World{Params: p, Box2d: &w}
	if world.Worker == nil {
		world.Worker = func(n int, f func(int), _ ...string) {
			for i := 0; i < n; i++ {
				f(i)
			}
		}
	}
	if p.Things.Walls.Top {
		world.SpawnWallH(0)
	}
	if p.Things.Walls.Bottom {
		world.SpawnWallH(-p.Things.Walls.Distance)
	}
	if p.Things.Walls.Left {
		world.SpawnWallV(0)
	}
	if p.Things.Walls.Right {
		world.SpawnWallV(p.Things.Walls.Distance)
	}
	for _, t := range p.Things.Static {
		t.New().SpawnStatic(world)
	}
	for _, t := range p.Things.Dynamic {
		t.New().SpawnDynamic(world)
	}
	return world
}

func (p *ThingParams) New() *Thing {
	shape := box2d.NewPolygonShape()
	if len(p.Polygons) > 0 {
		shape.Set(p.Polygons[0])
	}
	return &Thing{
		ThingParams: p,
		Shape:       shape,
	}
}

func (w *World) SpawnWallH(distance float64) {
	const (
		veryLarge = 1e16
		density   = 1.0
	)
	f := w.Box2d.CreateBody(&box2d.BodyDef{
		Type:     box2d.BodyTypeKinematicBody,
		Awake:    true,
		Active:   true,
		Position: box2d.Point{Y: distance},
	}).CreateFixture(func() *box2d.EdgeShape {
		shape := box2d.NewEdgeShape()
		shape.Set(box2d.Point{X: -veryLarge}, box2d.Point{X: veryLarge})
		return shape
	}(), density)
	f.Friction = w.Params.Things.Walls.Friction
	f.Restitution = w.Params.Things.Walls.Bounciness
}

func (w *World) SpawnWallV(distance float64) {
	const (
		veryLarge = 1e16
		density   = 1.0
	)
	f := w.Box2d.CreateBody(&box2d.BodyDef{
		Type:     box2d.BodyTypeKinematicBody,
		Awake:    true,
		Active:   true,
		Position: box2d.Point{X: distance},
	}).CreateFixture(func() *box2d.EdgeShape {
		shape := box2d.NewEdgeShape()
		shape.Set(box2d.Point{Y: -veryLarge}, box2d.Point{Y: veryLarge})
		return shape
	}(), density)
	f.Friction = w.Params.Things.Walls.Friction
	f.Restitution = w.Params.Things.Walls.Bounciness
}

func (t *Thing) SpawnDynamic(w *World) {
	const (
		gravityScale = 1.0
		density      = 1.0
	)
	t.Fixture = w.Box2d.CreateBody(&box2d.BodyDef{
		Type:            box2d.BodyTypeDynamicBody,
		Awake:           false,
		Active:          false,
		GravityScale:    gravityScale,
		AngularVelocity: t.Initial.AngularVelocityDeg,
		LinearVelocity:  t.Initial.LinearVelocity,
		Position:        t.Initial.Position,
	}).CreateFixture(t.Shape, density)
	w.Things.Dynamic = append(w.Things.Dynamic, t)
}

func (t *Thing) SpawnStatic(w *World) {
	const (
		gravityScale = 1.0
		density      = 1.0
	)
	t.Static = true
	t.Fixture = w.Box2d.CreateBody(&box2d.BodyDef{
		Type:            box2d.BodyTypeKinematicBody,
		Awake:           true,
		Active:          true,
		GravityScale:    gravityScale,
		AngularVelocity: t.Initial.AngularVelocityDeg,
		Position:        t.Initial.Position,
	}).CreateFixture(t.Shape, density)
	t.Fixture.Body.GetContactList()
	w.Things.Static = append(w.Things.Static, t)
}
