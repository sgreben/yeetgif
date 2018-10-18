package gifbounce

import (
	"image"

	"github.com/sgreben/yeetgif/pkg/box2d"
)

type ThingParams struct {
	Images         []image.Image
	Polygons       [][]box2d.Point
	Bounciness     func(float64) float64
	LinearDamping  func(float64) float64
	AngularDamping func(float64) float64
	Friction       func(float64) float64
	Initial        struct {
		LinearVelocity     box2d.Point
		Position           box2d.Point
		Time               float64
		AngularVelocityDeg float64
	}
}

type Params struct {
	NumFrames int
	Gravity   float64
	Worker    func(int, func(int), ...string)
	Things    struct {
		Dynamic []*ThingParams
		Static  []*ThingParams
		Walls   struct {
			Distance   float64
			Bounciness float64
			Friction   float64
			Left       bool
			Right      bool
			Top        bool
			Bottom     bool
		}
	}
	Solver struct {
		TimeStep           func(float64) float64
		VelocityIterations int
		PositionIterations int
	}
}
