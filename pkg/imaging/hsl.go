package imaging

import (
	"image/color"
	"math"
)

func HSLA(c color.RGBA) (h, s, l, a float64) {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0
	a = float64(c.A) / 255.0

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	l = (max + min) / 2
	delta := max - min
	if delta != 0 {
		if l < 0.5 {
			s = delta / (max + min)
		} else {
			s = delta / (2 - max - min)
		}
		r2 := (((max - r) / 6) + (delta / 2)) / delta
		g2 := (((max - g) / 6) + (delta / 2)) / delta
		b2 := (((max - b) / 6) + (delta / 2)) / delta
		switch {
		case r == max:
			h = b2 - g2
		case g == max:
			h = (1.0 / 3.0) + r2 - b2
		case b == max:
			h = (2.0 / 3.0) + g2 - r2
		}
	}

	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}
	return
}

func RGBA(h, s, l, a float64) (c color.RGBA) {
	switch {
	case h < 0:
		h = -(h - math.Ceil(h))
	case h > 1:
		h = h - math.Floor(h)
	}
	s = math.Max(0, math.Min(s, 1))
	l = math.Max(0, math.Min(l, 1))
	c.A = uint8(255 * a)
	if s == 0 {
		c.R = uint8(255 * l)
		c.G = uint8(255 * l)
		c.B = uint8(255 * l)
		return
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	c.R = uint8(255 * hrgb(v1, v2, h+(1.0/3.0)))
	c.G = uint8(255 * hrgb(v1, v2, h))
	c.B = uint8(255 * hrgb(v1, v2, h-(1.0/3.0)))

	return
}

func hrgb(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}
