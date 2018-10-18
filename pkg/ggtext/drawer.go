// Package ggtext draws text on images.
// It is a modified copy of a portion of github.com/fogleman/gg by Michael Fogleman.
// See LICENSE.md for the original copyright notice.
package ggtext

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
)

// Drawer is a font drawer
type Drawer struct {
	Face  font.Face
	Color color.Color
}

func (td *Drawer) fontHeight() float64 {
	return float64(td.Face.Metrics().Height) / 64
}

// Measure returns the rendered width and height of the specified text
// given the current font face.
func (td *Drawer) Measure(s string) (w, h float64) {
	d := &font.Drawer{Face: td.Face}
	a := d.MeasureString(s)
	return float64(a >> 6), td.fontHeight()
}

// Draw draws the specified text at the specified point.
func (td *Drawer) Draw(img draw.Image, s string, x, y float64) {
	td.DrawAnchored(img, s, x, y, 0, 0)
}

// DrawAnchored draws the specified text at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
func (td *Drawer) DrawAnchored(img draw.Image, s string, x, y, ax, ay float64) {
	w, h := td.Measure(s)
	x -= ax * w
	y += ay * h
	td.draw(img, s, x, y)
}

func (td *Drawer) draw(img draw.Image, s string, x, y float64) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(td.Color),
		Face: td.Face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(y * 64),
		},
	}
	// based on Drawer.DrawString() in golang.org/x/image/font/font.go
	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			continue
		}
		sr := dr.Sub(dr.Min)
		transformer := draw.BiLinear
		fx, fy := float64(dr.Min.X), float64(dr.Min.Y)
		s2d := f64.Aff3{1, 0, fx, 0, 1, fy}
		transformer.Transform(d.Dst, s2d, d.Src, sr, draw.Over, &draw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
}
