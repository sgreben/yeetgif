package gifbounce

import (
	"math"

	"github.com/sgreben/yeetgif/pkg/box2d"
)

type World struct {
	*Params
	Box2d  *box2d.World
	Things struct {
		Dynamic []*Thing
		Static  []*Thing
	}
}

func (w *World) ContainsDynamicThings(aabb box2d.AABB) bool {
	found := false
	w.Box2d.QueryAABB(func(fixture *box2d.Fixture) bool {
		found = fixture.Body.Type == box2d.BodyTypeDynamicBody
		return !found
	}, aabb)
	return found
}

func (w *World) Step(t float64) {
	for _, thing := range w.Things.Dynamic {
		if thing.Fixture.Body.IsActive() {
			continue
		}
		if t >= thing.Initial.Time {
			thing.Fixture.Body.SetActive(true)
			thing.Fixture.Body.SetAwake(true)
		}
	}
	w.Worker(len(w.Things.Static), func(i int) {
		w.Things.Static[i].Step(t)
		w.Things.Static[i].Record()
	})
	w.Worker(len(w.Things.Dynamic), func(i int) {
		w.Things.Dynamic[i].Step(t)
		w.Things.Dynamic[i].Record()
	})
	w.Box2d.Step(w.Solver.TimeStep(t), w.Solver.VelocityIterations, w.Solver.PositionIterations)
}

type Recording struct {
	Active       []bool
	Frames       []int
	Angles       []float64
	WorldCenters []box2d.Point
	LocalCenters []box2d.Point
	Bounds       []box2d.AABB
}

func (r *Recording) PadRightTo(k int) {
	k -= r.Len()
	if k < 0 {
		return
	}
	r.Active = append(append([]bool(nil), r.Active...), make([]bool, k)...)
	r.Frames = append(append([]int(nil), r.Frames...), make([]int, k)...)
	r.Angles = append(append([]float64(nil), r.Angles...), make([]float64, k)...)
	r.WorldCenters = append(append([]box2d.Point(nil), r.WorldCenters...), make([]box2d.Point, k)...)
	r.LocalCenters = append(append([]box2d.Point(nil), r.LocalCenters...), make([]box2d.Point, k)...)
	r.Bounds = append(append([]box2d.AABB(nil), r.Bounds...), make([]box2d.AABB, k)...)
}

func (r *Recording) Slice(i, j int) Recording {
	return Recording{
		Active:       r.Active[i:j],
		Frames:       r.Frames[i:j],
		Angles:       r.Angles[i:j],
		WorldCenters: r.WorldCenters[i:j],
		LocalCenters: r.LocalCenters[i:j],
		Bounds:       r.Bounds[i:j],
	}
}

func (r *Recording) Len() int {
	return len(r.WorldCenters)
}

func (r *Recording) Record(active bool, frame int, angleDeg float64, worldCenter, localCenter box2d.Point, aabb box2d.AABB) {
	r.Active = append(r.Active, active)
	r.Frames = append(r.Frames, frame)
	r.Angles = append(r.Angles, angleDeg)
	r.WorldCenters = append(r.WorldCenters, worldCenter)
	r.LocalCenters = append(r.LocalCenters, localCenter)
	r.Bounds = append(r.Bounds, aabb)
}

type Thing struct {
	*ThingParams
	Shape   *box2d.PolygonShape
	Fixture *box2d.Fixture
	Frame   int

	Static    bool
	Recording Recording
}

func (t *Thing) Step(time float64) {
	if t.Friction != nil {
		t.Fixture.Friction = t.Friction(time)
	}
	if t.Bounciness != nil {
		t.Fixture.Restitution = t.Bounciness(time)
	}
	if t.LinearDamping != nil {
		t.Fixture.Body.LinearDamping = t.LinearDamping(time)
	}
	if t.AngularDamping != nil {
		t.Fixture.Body.AngularDamping = t.AngularDamping(time)
	}
	t.Frame++
	if t.Frame >= len(t.Polygons) {
		t.Frame = 0
	}
	t.Shape.Set(t.Polygons[t.Frame])
}

func (t *Thing) WorldCenter() box2d.Point {
	if t.Static {
		return t.Fixture.GetAABB(0).GetCenter()
	}
	return t.Fixture.Body.GetWorldCenter()
}

func (t *Thing) LocalCenter() box2d.Point {
	body := t.Fixture.Body
	if t.Static {
		pos := body.GetPosition()
		bounds := t.Fixture.GetAABB(0)
		return box2d.Point{
			X: pos.X - bounds.Min.X,
			Y: pos.Y - bounds.Min.Y,
		}
	}
	return t.Fixture.Body.GetLocalCenter()
}

func (t *Thing) Record() {
	body := t.Fixture.Body
	angleDeg := body.GetAngle() * 180.0 / math.Pi
	worldCenter, localCenter := t.WorldCenter(), t.LocalCenter()
	bounds := t.Fixture.GetAABB(0)
	t.Recording.Record(body.IsActive(), t.Frame, angleDeg, worldCenter, localCenter, bounds.Clone())
}
