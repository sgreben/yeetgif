package render

import "image/color"

type Options struct {
	Scale        float64
	Stroke, Fill color.RGBA
	Width        float64
}
