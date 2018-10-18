package svg

import (
	"image/color"

	"github.com/sgreben/yeetgif/pkg/rendersvg/geom"
)

type Image struct {
	ViewBox       geom.Rectangle
	Width, Height float64
	Shapes        []Shape
}

type Shape struct {
	Path      *Path
	Circle    *Circle
	Ellipse   *Ellipse
	Composite []Shape
	ColorAttributes
}

type Circle struct {
	Center geom.Point
	Radius float64
}

type Ellipse struct {
	Center geom.Point
	Radius geom.Point
}

type Path struct {
	Commands []PathCommand
}

type ColorAttributes struct {
	FillColor     *color.RGBA
	FillOpacity   *float64
	StrokeColor   *color.RGBA
	StrokeOpacity *float64
	StrokeWidth   *float64
}
