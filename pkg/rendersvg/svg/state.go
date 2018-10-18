package svg

import "github.com/sgreben/yeetgif/pkg/rendersvg/geom"

// PathState is a path interpreter state
type PathState struct {
	CurrentPosition geom.Point
	Positions       []geom.Point
	LastCommand     PathCommand
}

// Command updates the state by executing the given command
func (state *PathState) Command(c PathCommand) (upper PathCommand, pos geom.Point) {
	defer func() {
		state.LastCommand = upper
	}()
	state.Positions = append(state.Positions, state.CurrentPosition)
	switch {
	case c.MoveTo != nil:
		if !c.MoveTo.Upper {
			c.MoveTo.Upper = true
			c.MoveTo.X += state.CurrentPosition.X
			c.MoveTo.Y += state.CurrentPosition.Y
		}
		state.CurrentPosition.X = c.MoveTo.X
		state.CurrentPosition.Y = c.MoveTo.Y
	case c.LineTo != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.LineTo.Upper {
			c.LineTo.Upper = true
			c.LineTo.X += state.CurrentPosition.X
			c.LineTo.Y += state.CurrentPosition.Y
		}
		state.CurrentPosition.X = c.LineTo.X
		state.CurrentPosition.Y = c.LineTo.Y
	case c.H != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.H.Upper {
			c.H.Upper = true
			c.H.X += state.CurrentPosition.X
		}
		state.CurrentPosition.X = c.H.X
	case c.V != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.V.Upper {
			c.V.Upper = true
			c.V.Y += state.CurrentPosition.Y
		}
		state.CurrentPosition.Y = c.V.Y
	case c.Z != nil:
		if len(state.Positions) > 0 {
			state.CurrentPosition = state.Positions[0]
			state.Positions = state.Positions[:0]
		}
	case c.CubicBezier != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.CubicBezier.Upper {
			c.CubicBezier.Upper = true
			c.CubicBezier.X += state.CurrentPosition.X
			c.CubicBezier.Y += state.CurrentPosition.Y
			c.CubicBezier.X1 += state.CurrentPosition.X
			c.CubicBezier.Y1 += state.CurrentPosition.Y
			c.CubicBezier.X2 += state.CurrentPosition.X
			c.CubicBezier.Y2 += state.CurrentPosition.Y
		}
		state.CurrentPosition.X = c.CubicBezier.X
		state.CurrentPosition.Y = c.CubicBezier.Y
	case c.CubicBezierShortcut != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.CubicBezierShortcut.Upper {
			c.CubicBezierShortcut.Upper = true
			c.CubicBezierShortcut.X += state.CurrentPosition.X
			c.CubicBezierShortcut.Y += state.CurrentPosition.Y
			c.CubicBezierShortcut.X2 += state.CurrentPosition.X
			c.CubicBezierShortcut.Y2 += state.CurrentPosition.Y
		}
		p1 := state.CurrentPosition
		if state.LastCommand.Kind() == PathCommandKindCubicBezier {
			p1 = geom.Point{
				X: state.LastCommand.CubicBezier.X2,
				Y: state.LastCommand.CubicBezier.Y2,
			}.ReflectAround(state.CurrentPosition)
		}
		state.CurrentPosition.X = c.CubicBezierShortcut.X
		state.CurrentPosition.Y = c.CubicBezierShortcut.Y
		return PathCommand{
			CubicBezier: &PathCommandCubicBezier{
				X:  c.CubicBezierShortcut.X,
				Y:  c.CubicBezierShortcut.Y,
				X1: p1.X,
				Y1: p1.Y,
				X2: c.CubicBezierShortcut.X2,
				Y2: c.CubicBezierShortcut.Y2,
			},
		}, state.CurrentPosition
	case c.QuadraticBezier != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.QuadraticBezier.Upper {
			c.QuadraticBezier.Upper = true
			c.QuadraticBezier.X += state.CurrentPosition.X
			c.QuadraticBezier.Y += state.CurrentPosition.Y
			c.QuadraticBezier.X1 += state.CurrentPosition.X
			c.QuadraticBezier.Y1 += state.CurrentPosition.Y
		}
		state.CurrentPosition.X = c.QuadraticBezier.X
		state.CurrentPosition.Y = c.QuadraticBezier.Y
	case c.QuadraticBezierShortcut != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.QuadraticBezierShortcut.Upper {
			c.QuadraticBezierShortcut.Upper = true
			c.QuadraticBezierShortcut.X += state.CurrentPosition.X
			c.QuadraticBezierShortcut.Y += state.CurrentPosition.Y
		}
		p1 := state.CurrentPosition
		if state.LastCommand.Kind() == PathCommandKindQuadraticBezier {
			p1 = geom.Point{
				X: state.LastCommand.QuadraticBezier.X1,
				Y: state.LastCommand.QuadraticBezier.Y1,
			}.ReflectAround(state.CurrentPosition)
		}
		state.CurrentPosition.X = c.QuadraticBezierShortcut.X
		state.CurrentPosition.Y = c.QuadraticBezierShortcut.Y
		return PathCommand{
			QuadraticBezier: &PathCommandQuadraticBezier{
				X:  c.QuadraticBezierShortcut.X,
				Y:  c.QuadraticBezierShortcut.Y,
				X1: p1.X,
				Y1: p1.Y,
			},
		}, state.CurrentPosition
	case c.Arc != nil:
		state.Positions = append(state.Positions, state.CurrentPosition)
		if !c.Arc.Upper {
			c.Arc.Upper = true
			c.Arc.X += state.CurrentPosition.X
			c.Arc.Y += state.CurrentPosition.Y
		}
		state.CurrentPosition.X = c.Arc.X
		state.CurrentPosition.Y = c.Arc.Y
	}
	return c, state.CurrentPosition
}
