package render

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/sgreben/yeetgif/pkg/rendersvg/svg"
)

var DefaultPathOptions = Options{
	Scale:  1.0,
	Width:  0.0,
	Stroke: color.RGBA{A: 255},
	Fill:   color.RGBA{255, 255, 255, 255},
}

func pathCommands(cs []svg.PathCommand, ctx *gg.Context) {
	state := &svg.PathState{}
	for _, c := range cs {
		c, p := state.Command(c)
		switch {
		case c.MoveTo != nil:
			ctx.MoveTo(p.X, p.Y)
		case c.LineTo != nil, c.H != nil, c.V != nil, c.Arc != nil:
			ctx.LineTo(p.X, p.Y)
		case c.Z != nil:
			ctx.ClosePath()
		case c.CubicBezier != nil:
			ctx.CubicTo(c.CubicBezier.X1, c.CubicBezier.Y1, c.CubicBezier.X2, c.CubicBezier.Y2, c.CubicBezier.X, c.CubicBezier.Y)
		case c.CubicBezierShortcut != nil:
			ctx.CubicTo(c.CubicBezier.X1, c.CubicBezier.Y1, c.CubicBezier.X2, c.CubicBezier.Y2, c.CubicBezier.X, c.CubicBezier.Y)
		case c.QuadraticBezier != nil:
			ctx.QuadraticTo(c.QuadraticBezier.X1, c.QuadraticBezier.Y1, c.QuadraticBezier.X, c.QuadraticBezier.Y)
		case c.QuadraticBezierShortcut != nil:
			ctx.QuadraticTo(c.QuadraticBezier.X1, c.QuadraticBezier.Y1, c.QuadraticBezier.X, c.QuadraticBezier.Y)
		}
	}
}
