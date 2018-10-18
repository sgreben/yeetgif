package render

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"

	"github.com/sgreben/yeetgif/pkg/rendersvg/svg"
)

var DefaultImageOptions = Options{
	Scale:  1.0,
	Width:  1.0,
	Stroke: color.RGBA{A: 255},
	Fill:   color.RGBA{255, 255, 255, 255},
}

func shape(s svg.Shape, ctx *gg.Context, o Options) {
	if s.FillColor != nil {
		o.Fill = *s.FillColor
	}
	if s.StrokeColor != nil {
		o.Stroke = *s.StrokeColor
	}
	if s.StrokeWidth != nil {
		o.Width = *s.StrokeWidth
	}
	switch {
	case s.Path != nil:
		p := s.Path
		pathCommands(p.Commands, ctx)
		ctx.SetLineWidth(o.Width)
		ctx.SetColor(o.Stroke)
		ctx.StrokePreserve()
		ctx.SetColor(o.Fill)
		ctx.Fill()
	case s.Circle != nil:
		c := s.Circle
		ctx.DrawCircle(c.Center.X, c.Center.Y, c.Radius)
		ctx.SetLineWidth(o.Width)
		ctx.SetColor(o.Stroke)
		ctx.StrokePreserve()
		ctx.SetColor(o.Fill)
		ctx.Fill()
	case s.Ellipse != nil:
		e := s.Ellipse
		ctx.DrawEllipse(e.Center.X, e.Center.Y, e.Radius.X, e.Radius.Y)
		ctx.SetLineWidth(o.Width)
		ctx.SetColor(o.Stroke)
		ctx.StrokePreserve()
		ctx.SetColor(o.Fill)
		ctx.Fill()
	case s.Composite != nil:
		for _, s := range s.Composite {
			shape(s, ctx, o)
		}
	}

}

func Image(img svg.Image, out *image.RGBA, o *Options) {
	if o == nil {
		o = &DefaultPathOptions
	}
	o.Scale = img.ViewBox.ScaleToFit(out.Bounds())
	ctx := gg.NewContextForRGBA(out)
	ctx.Scale(o.Scale, o.Scale)
	for _, s := range img.Shapes {
		shape(s, ctx, *o)
	}
}
