package svg

type commandKind string

// PathCommandKind is a SVG path command kind
type PathCommandKind commandKind

// All supported command kinds
const (
	PathCommandKindMoveTo                  PathCommandKind = "M"
	PathCommandKindLineTo                  PathCommandKind = "L"
	PathCommandKindH                       PathCommandKind = "H"
	PathCommandKindV                       PathCommandKind = "V"
	PathCommandKindZ                       PathCommandKind = "Z"
	PathCommandKindCubicBezier             PathCommandKind = "C"
	PathCommandKindCubicBezierShortcut     PathCommandKind = "S"
	PathCommandKindQuadraticBezier         PathCommandKind = "Q"
	PathCommandKindQuadraticBezierShortcut PathCommandKind = "T"
	PathCommandKindArc                     PathCommandKind = "A"
)

// PathCommand is an SVG path command
type PathCommand struct {
	MoveTo                  *PathCommandMoveTo                  `json:",omitempty"` // Mm
	LineTo                  *PathCommandLineTo                  `json:",omitempty"` // Ll
	H                       *PathCommandH                       `json:",omitempty"` // Hh
	V                       *PathCommandV                       `json:",omitempty"` // Vv
	Z                       *PathCommandZ                       `json:",omitempty"` // Zz
	CubicBezier             *PathCommandCubicBezier             `json:",omitempty"` // Cc
	CubicBezierShortcut     *PathCommandCubicBezierShortcut     `json:",omitempty"` // Ss
	QuadraticBezier         *PathCommandQuadraticBezier         `json:",omitempty"` // Qq
	QuadraticBezierShortcut *PathCommandQuadraticBezierShortcut `json:",omitempty"` // Tt
	Arc                     *PathCommandArc                     `json:",omitempty"` // Aa
}

// Kind returns the command's kind
func (c *PathCommand) Kind() PathCommandKind {
	switch {
	case c.MoveTo != nil:
		return PathCommandKindMoveTo
	case c.LineTo != nil:
		return PathCommandKindLineTo
	case c.H != nil:
		return PathCommandKindH
	case c.V != nil:
		return PathCommandKindV
	case c.Z != nil:
		return PathCommandKindZ
	case c.CubicBezier != nil:
		return PathCommandKindCubicBezier
	case c.CubicBezierShortcut != nil:
		return PathCommandKindCubicBezierShortcut
	case c.QuadraticBezier != nil:
		return PathCommandKindQuadraticBezier
	case c.QuadraticBezierShortcut != nil:
		return PathCommandKindQuadraticBezierShortcut
	case c.Arc != nil:
		return PathCommandKindArc
	default:
		return ""
	}
}

// PathCommandMoveTo is the M/m SVG path command
type PathCommandMoveTo struct {
	Upper bool
	X, Y  float64
}

// PathCommandLineTo is the L/l SVG path command
type PathCommandLineTo struct {
	Upper bool
	X, Y  float64
}

// PathCommandH is the H/h SVG path command
type PathCommandH struct {
	Upper bool
	X     float64
}

// PathCommandV is the V/v SVG path command
type PathCommandV struct {
	Upper bool
	Y     float64
}

// PathCommandZ is the Z/z SVG path command
type PathCommandZ struct {
	Upper bool
}

// PathCommandCubicBezier is the C/c SVG path command
type PathCommandCubicBezier struct {
	Upper                bool
	X1, Y1, X2, Y2, X, Y float64
}

// PathCommandCubicBezierShortcut is the S/s SVG path command
type PathCommandCubicBezierShortcut struct {
	Upper        bool
	X2, Y2, X, Y float64
}

// PathCommandQuadraticBezier is the Q/q SVG path command
type PathCommandQuadraticBezier struct {
	Upper        bool
	X1, Y1, X, Y float64
}

// PathCommandQuadraticBezierShortcut is the T/t SVG path command
type PathCommandQuadraticBezierShortcut struct {
	Upper bool
	X, Y  float64
}

// PathCommandArc is the A/a SVG path command
type PathCommandArc struct {
	Upper                                                bool
	Rx, Ry, XAxisRotation, LargeArcFlag, SweepFlag, X, Y float64
}
