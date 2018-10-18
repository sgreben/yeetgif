package parse

import (
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

var (
	expHex  = regexp.MustCompile("^#[0-9A-Fa-f]+$")
	expHex3 = regexp.MustCompile("^#[0-9A-Fa-f]{3}$")
	expHex4 = regexp.MustCompile("^#[0-9A-Fa-f]{4}$")
	expHex6 = regexp.MustCompile("^#[0-9A-Fa-f]{6}$")
	expHex8 = regexp.MustCompile("^#[0-9A-Fa-f]{8}$")
	expRgb  = regexp.MustCompile("rgb([^,)]+,[^,)]+,[^,)])")
	expRgba = regexp.MustCompile("rgba([^,)]+,[^,)]+,[^,)])")
	expHsl  = regexp.MustCompile("hsl([^,)]+,[^,)]+,[^,)])")
	expHsla = regexp.MustCompile("hsla([^,)]+,[^,)]+,[^,)])")
)

func Color(s string) (*color.RGBA, error) {
	s = strings.TrimSpace(s)
	keywordColor, isKeyword := keywords[s]
	switch {
	case isKeyword:
		return &keywordColor, nil
	case expHex.MatchString(s):
		return HexColor(s)
	case expRgb.MatchString(s):
		fallthrough // TODO: parse expRgb
	case expRgba.MatchString(s):
		fallthrough // TODO: parse expRgba
	case expHsl.MatchString(s):
		fallthrough // TODO: parse expHsl
	case expHsla.MatchString(s):
		fallthrough // TODO: parse expHsla
	default:
		return nil, fmt.Errorf("can not parse color: %q", s)
	}
}

func HexColor(s string) (*color.RGBA, error) {
	dup := func(x uint64) uint8 { return uint8(x | (x << 4)) }
	switch {
	case expHex3.MatchString(s):
		var rgb uint64
		_, err := fmt.Sscanf(s, "#%x", &rgb)
		c := color.RGBA{R: dup((rgb >> 8) & 0xF), G: dup((rgb >> 4) & 0xF), B: dup(rgb & 0xF), A: 0xFF}
		return &c, err
	case expHex4.MatchString(s):
		var rgba uint64
		_, err := fmt.Sscanf(s, "#%x", &rgba)
		c := color.RGBA{R: dup((rgba >> 12) & 0xF), G: dup((rgba >> 8) & 0xF), B: dup((rgba >> 4) & 0xF), A: dup(rgba & 0xF)}
		return &c, err
	case expHex6.MatchString(s):
		var rgb uint64
		_, err := fmt.Sscanf(s, "#%x", &rgb)
		c := color.RGBA{R: uint8((rgb >> 16) & 0xFF), G: uint8((rgb >> 8) & 0xFF), B: uint8(rgb & 0xFF), A: 0xFF}
		return &c, err
	case expHex8.MatchString(s):
		var rgba uint64
		_, err := fmt.Sscanf(s, "#%x", &rgba)
		c := color.RGBA{R: uint8((rgba >> 24) & 0xFF), G: uint8((rgba >> 16) & 0xFF), B: uint8((rgba >> 8) & 0xFF), A: uint8(rgba & 0xFF)}
		return &c, err
	default:
		return nil, fmt.Errorf("can not parse hex color: %q", s)
	}
}

func hslToRGB(h, s, l float64) (r, g, b float64) {
	m2 := 0.0
	if l <= 0.5 {
		m2 = l * (s + 1)
	} else {
		m2 = l + s - l*s
	}
	m1 := l*2 - m2
	r = hueToRGB(m1, m2, h+1/3)
	g = hueToRGB(m1, m2, h)
	b = hueToRGB(m1, m2, h-1/3)
	return
}

func hueToRGB(m1, m2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	if h*6 < 1 {
		return m1 + (m2-m1)*h*6
	}
	if h*2 < 1 {
		return m2
	}
	if h*3 < 2 {
		return m1 + (m2-m1)*(2/3-h)*6
	}
	return m1
}

var keywords = map[string]color.RGBA{
	"aliceblue":            color.RGBA{R: 240, G: 248, B: 255, A: 255},
	"antiquewhite":         color.RGBA{R: 250, G: 235, B: 215, A: 255},
	"aqua":                 color.RGBA{R: 0, G: 255, B: 255, A: 255},
	"aquamarine":           color.RGBA{R: 127, G: 255, B: 212, A: 255},
	"azure":                color.RGBA{R: 240, G: 255, B: 255, A: 255},
	"beige":                color.RGBA{R: 245, G: 245, B: 220, A: 255},
	"bisque":               color.RGBA{R: 255, G: 228, B: 196, A: 255},
	"black":                color.RGBA{R: 0, G: 0, B: 0, A: 255},
	"blanchedalmond":       color.RGBA{R: 255, G: 235, B: 205, A: 255},
	"blue":                 color.RGBA{R: 0, G: 0, B: 255, A: 255},
	"blueviolet":           color.RGBA{R: 138, G: 43, B: 226, A: 255},
	"brown":                color.RGBA{R: 165, G: 42, B: 42, A: 255},
	"burlywood":            color.RGBA{R: 222, G: 184, B: 135, A: 255},
	"cadetblue":            color.RGBA{R: 95, G: 158, B: 160, A: 255},
	"chartreuse":           color.RGBA{R: 127, G: 255, B: 0, A: 255},
	"chocolate":            color.RGBA{R: 210, G: 105, B: 30, A: 255},
	"coral":                color.RGBA{R: 255, G: 127, B: 80, A: 255},
	"cornflowerblue":       color.RGBA{R: 100, G: 149, B: 237, A: 255},
	"cornsilk":             color.RGBA{R: 255, G: 248, B: 220, A: 255},
	"crimson":              color.RGBA{R: 220, G: 20, B: 60, A: 255},
	"cyan":                 color.RGBA{R: 0, G: 255, B: 255, A: 255},
	"darkblue":             color.RGBA{R: 0, G: 0, B: 139, A: 255},
	"darkcyan":             color.RGBA{R: 0, G: 139, B: 139, A: 255},
	"darkgoldenrod":        color.RGBA{R: 184, G: 134, B: 11, A: 255},
	"darkgray":             color.RGBA{R: 169, G: 169, B: 169, A: 255},
	"darkgreen":            color.RGBA{R: 0, G: 100, B: 0, A: 255},
	"darkgrey":             color.RGBA{R: 169, G: 169, B: 169, A: 255},
	"darkkhaki":            color.RGBA{R: 189, G: 183, B: 107, A: 255},
	"darkmagenta":          color.RGBA{R: 139, G: 0, B: 139, A: 255},
	"darkolivegreen":       color.RGBA{R: 85, G: 107, B: 47, A: 255},
	"darkorange":           color.RGBA{R: 255, G: 140, B: 0, A: 255},
	"darkorchid":           color.RGBA{R: 153, G: 50, B: 204, A: 255},
	"darkred":              color.RGBA{R: 139, G: 0, B: 0, A: 255},
	"darksalmon":           color.RGBA{R: 233, G: 150, B: 122, A: 255},
	"darkseagreen":         color.RGBA{R: 143, G: 188, B: 143, A: 255},
	"darkslateblue":        color.RGBA{R: 72, G: 61, B: 139, A: 255},
	"darkslategray":        color.RGBA{R: 47, G: 79, B: 79, A: 255},
	"darkslategrey":        color.RGBA{R: 47, G: 79, B: 79, A: 255},
	"darkturquoise":        color.RGBA{R: 0, G: 206, B: 209, A: 255},
	"darkviolet":           color.RGBA{R: 148, G: 0, B: 211, A: 255},
	"deeppink":             color.RGBA{R: 255, G: 20, B: 147, A: 255},
	"deepskyblue":          color.RGBA{R: 0, G: 191, B: 255, A: 255},
	"dimgray":              color.RGBA{R: 105, G: 105, B: 105, A: 255},
	"dimgrey":              color.RGBA{R: 105, G: 105, B: 105, A: 255},
	"dodgerblue":           color.RGBA{R: 30, G: 144, B: 255, A: 255},
	"firebrick":            color.RGBA{R: 178, G: 34, B: 34, A: 255},
	"floralwhite":          color.RGBA{R: 255, G: 250, B: 240, A: 255},
	"forestgreen":          color.RGBA{R: 34, G: 139, B: 34, A: 255},
	"fuchsia":              color.RGBA{R: 255, G: 0, B: 255, A: 255},
	"gainsboro":            color.RGBA{R: 220, G: 220, B: 220, A: 255},
	"ghostwhite":           color.RGBA{R: 248, G: 248, B: 255, A: 255},
	"gold":                 color.RGBA{R: 255, G: 215, B: 0, A: 255},
	"goldenrod":            color.RGBA{R: 218, G: 165, B: 32, A: 255},
	"gray":                 color.RGBA{R: 128, G: 128, B: 128, A: 255},
	"green":                color.RGBA{R: 0, G: 128, B: 0, A: 255},
	"greenyellow":          color.RGBA{R: 173, G: 255, B: 47, A: 255},
	"grey":                 color.RGBA{R: 128, G: 128, B: 128, A: 255},
	"honeydew":             color.RGBA{R: 240, G: 255, B: 240, A: 255},
	"hotpink":              color.RGBA{R: 255, G: 105, B: 180, A: 255},
	"indianred":            color.RGBA{R: 205, G: 92, B: 92, A: 255},
	"indigo":               color.RGBA{R: 75, G: 0, B: 130, A: 255},
	"ivory":                color.RGBA{R: 255, G: 255, B: 240, A: 255},
	"khaki":                color.RGBA{R: 240, G: 230, B: 140, A: 255},
	"lavender":             color.RGBA{R: 230, G: 230, B: 250, A: 255},
	"lavenderblush":        color.RGBA{R: 255, G: 240, B: 245, A: 255},
	"lawngreen":            color.RGBA{R: 124, G: 252, B: 0, A: 255},
	"lemonchiffon":         color.RGBA{R: 255, G: 250, B: 205, A: 255},
	"lightblue":            color.RGBA{R: 173, G: 216, B: 230, A: 255},
	"lightcoral":           color.RGBA{R: 240, G: 128, B: 128, A: 255},
	"lightcyan":            color.RGBA{R: 224, G: 255, B: 255, A: 255},
	"lightgoldenrodyellow": color.RGBA{R: 250, G: 250, B: 210, A: 255},
	"lightgray":            color.RGBA{R: 211, G: 211, B: 211, A: 255},
	"lightgreen":           color.RGBA{R: 144, G: 238, B: 144, A: 255},
	"lightgrey":            color.RGBA{R: 211, G: 211, B: 211, A: 255},
	"lightpink":            color.RGBA{R: 255, G: 182, B: 193, A: 255},
	"lightsalmon":          color.RGBA{R: 255, G: 160, B: 122, A: 255},
	"lightseagreen":        color.RGBA{R: 32, G: 178, B: 170, A: 255},
	"lightskyblue":         color.RGBA{R: 135, G: 206, B: 250, A: 255},
	"lightslategray":       color.RGBA{R: 119, G: 136, B: 153, A: 255},
	"lightslategrey":       color.RGBA{R: 119, G: 136, B: 153, A: 255},
	"lightsteelblue":       color.RGBA{R: 176, G: 196, B: 222, A: 255},
	"lightyellow":          color.RGBA{R: 255, G: 255, B: 224, A: 255},
	"lime":                 color.RGBA{R: 0, G: 255, B: 0, A: 255},
	"limegreen":            color.RGBA{R: 50, G: 205, B: 50, A: 255},
	"linen":                color.RGBA{R: 250, G: 240, B: 230, A: 255},
	"magenta":              color.RGBA{R: 255, G: 0, B: 255, A: 255},
	"maroon":               color.RGBA{R: 128, G: 0, B: 0, A: 255},
	"mediumaquamarine":     color.RGBA{R: 102, G: 205, B: 170, A: 255},
	"mediumblue":           color.RGBA{R: 0, G: 0, B: 205, A: 255},
	"mediumorchid":         color.RGBA{R: 186, G: 85, B: 211, A: 255},
	"mediumpurple":         color.RGBA{R: 147, G: 112, B: 219, A: 255},
	"mediumseagreen":       color.RGBA{R: 60, G: 179, B: 113, A: 255},
	"mediumslateblue":      color.RGBA{R: 123, G: 104, B: 238, A: 255},
	"mediumspringgreen":    color.RGBA{R: 0, G: 250, B: 154, A: 255},
	"mediumturquoise":      color.RGBA{R: 72, G: 209, B: 204, A: 255},
	"mediumvioletred":      color.RGBA{R: 199, G: 21, B: 133, A: 255},
	"midnightblue":         color.RGBA{R: 25, G: 25, B: 112, A: 255},
	"mintcream":            color.RGBA{R: 245, G: 255, B: 250, A: 255},
	"mistyrose":            color.RGBA{R: 255, G: 228, B: 225, A: 255},
	"moccasin":             color.RGBA{R: 255, G: 228, B: 181, A: 255},
	"navajowhite":          color.RGBA{R: 255, G: 222, B: 173, A: 255},
	"navy":                 color.RGBA{R: 0, G: 0, B: 128, A: 255},
	"oldlace":              color.RGBA{R: 253, G: 245, B: 230, A: 255},
	"olive":                color.RGBA{R: 128, G: 128, B: 0, A: 255},
	"olivedrab":            color.RGBA{R: 107, G: 142, B: 35, A: 255},
	"orange":               color.RGBA{R: 255, G: 165, B: 0, A: 255},
	"orangered":            color.RGBA{R: 255, G: 69, B: 0, A: 255},
	"orchid":               color.RGBA{R: 218, G: 112, B: 214, A: 255},
	"palegoldenrod":        color.RGBA{R: 238, G: 232, B: 170, A: 255},
	"palegreen":            color.RGBA{R: 152, G: 251, B: 152, A: 255},
	"paleturquoise":        color.RGBA{R: 175, G: 238, B: 238, A: 255},
	"palevioletred":        color.RGBA{R: 219, G: 112, B: 147, A: 255},
	"papayawhip":           color.RGBA{R: 255, G: 239, B: 213, A: 255},
	"peachpuff":            color.RGBA{R: 255, G: 218, B: 185, A: 255},
	"peru":                 color.RGBA{R: 205, G: 133, B: 63, A: 255},
	"pink":                 color.RGBA{R: 255, G: 192, B: 203, A: 255},
	"plum":                 color.RGBA{R: 221, G: 160, B: 221, A: 255},
	"powderblue":           color.RGBA{R: 176, G: 224, B: 230, A: 255},
	"purple":               color.RGBA{R: 128, G: 0, B: 128, A: 255},
	"red":                  color.RGBA{R: 255, G: 0, B: 0, A: 255},
	"rosybrown":            color.RGBA{R: 188, G: 143, B: 143, A: 255},
	"royalblue":            color.RGBA{R: 65, G: 105, B: 225, A: 255},
	"saddlebrown":          color.RGBA{R: 139, G: 69, B: 19, A: 255},
	"salmon":               color.RGBA{R: 250, G: 128, B: 114, A: 255},
	"sandybrown":           color.RGBA{R: 244, G: 164, B: 96, A: 255},
	"seagreen":             color.RGBA{R: 46, G: 139, B: 87, A: 255},
	"seashell":             color.RGBA{R: 255, G: 245, B: 238, A: 255},
	"sienna":               color.RGBA{R: 160, G: 82, B: 45, A: 255},
	"silver":               color.RGBA{R: 192, G: 192, B: 192, A: 255},
	"skyblue":              color.RGBA{R: 135, G: 206, B: 235, A: 255},
	"slateblue":            color.RGBA{R: 106, G: 90, B: 205, A: 255},
	"slategray":            color.RGBA{R: 112, G: 128, B: 144, A: 255},
	"slategrey":            color.RGBA{R: 112, G: 128, B: 144, A: 255},
	"snow":                 color.RGBA{R: 255, G: 250, B: 250, A: 255},
	"springgreen":          color.RGBA{R: 0, G: 255, B: 127, A: 255},
	"steelblue":            color.RGBA{R: 70, G: 130, B: 180, A: 255},
	"tan":                  color.RGBA{R: 210, G: 180, B: 140, A: 255},
	"teal":                 color.RGBA{R: 0, G: 128, B: 128, A: 255},
	"thistle":              color.RGBA{R: 216, G: 191, B: 216, A: 255},
	"tomato":               color.RGBA{R: 255, G: 99, B: 71, A: 255},
	"turquoise":            color.RGBA{R: 64, G: 224, B: 208, A: 255},
	"violet":               color.RGBA{R: 238, G: 130, B: 238, A: 255},
	"wheat":                color.RGBA{R: 245, G: 222, B: 179, A: 255},
	"white":                color.RGBA{R: 255, G: 255, B: 255, A: 255},
	"whitesmoke":           color.RGBA{R: 245, G: 245, B: 245, A: 255},
	"yellow":               color.RGBA{R: 255, G: 255, B: 0, A: 255},
	"yellowgreen":          color.RGBA{R: 154, G: 205, B: 50, A: 255},
}
