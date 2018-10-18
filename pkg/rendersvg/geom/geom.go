package geom

import (
	"image"
	"math"
)

// Point is a two-dimensional point
type Point struct{ X, Y float64 }

func (p Point) ReflectAround(other Point) (out Point) {
	dX := other.X - p.X
	dY := other.Y - p.Y
	out.X = other.X + dX
	out.Y = other.Y + dY
	return
}

// Rectangle is a two-dimensional rectangle
type Rectangle struct{ Min, Max Point }

func (r *Rectangle) Dx() float64 { return r.Max.X - r.Min.X }
func (r *Rectangle) Dy() float64 { return r.Max.Y - r.Min.Y }

func (r *Rectangle) ScaleToFit(out image.Rectangle) float64 {
	return math.Min(float64(out.Dx())/r.Dx(), float64(out.Dy())/r.Dy())
}
