package imaging

import (
	"image/color"
	"math"
)

var byteToFloat [256]float64

func init() {
	for i := 0; i <= 0xFF; i++ {
		byteToFloat[i] = float64(i) / 0xFF
	}
}

func HSLA(c color.RGBA) (h, s, l, a float64) {
	r := byteToFloat[c.R]
	g := byteToFloat[c.G]
	b := byteToFloat[c.B]
	a = byteToFloat[c.A]

	maxByte := c.R
	if c.G > maxByte {
		maxByte = c.G
	}
	if c.B > maxByte {
		maxByte = c.B
	}
	max := byteToFloat[maxByte]
	minByte := c.R
	if c.G < minByte {
		minByte = c.G
	}
	if c.B < minByte {
		minByte = c.B
	}
	min := byteToFloat[minByte]
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
	switch {
	case s < 0:
		s = 0
	case s > 1:
		s = 1
	}
	switch {
	case l < 0:
		l = 0
	case l > 1:
		l = 1
	}
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
