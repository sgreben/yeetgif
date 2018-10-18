package parse

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/sgreben/yeetgif/pkg/rendersvg/svg"
)

var (
	matchCommandCode = regexp.MustCompile(`[amlhvzcsqtAMLHVZCSQT]`)
	isCommandCode    = map[rune]bool{'a': true, 'm': true, 'l': true, 'h': true, 'v': true, 'z': true, 'c': true, 's': true, 'q': true, 't': true}
	expSpace         = regexp.MustCompile(`[\s]+`)
	expFloat         = regexp.MustCompile(`^-?([0-9]+|[0-9]*\.[0-9]+)$`)
)

type rawCommand struct {
	Code      rune
	Arguments []float64
}

func rawCommands(r io.Reader) ([]rawCommand, error) {
	var commands []rawCommand
	scanner := bufio.NewScanner(r)
	scanner.Split(scanCommand)
	for scanner.Scan() {
		commandString := scanner.Text()
		commandString = strings.TrimSpace(commandString)
		commandString = expSpace.ReplaceAllString(commandString, " ")
		code := []rune(matchCommandCode.FindString(commandString))[0]
		argsString := commandString[1:]
		argsString = strings.Replace(argsString, ",", " ", -1)
		command := rawCommand{Code: code}
		argsScanner := bufio.NewScanner(strings.NewReader(argsString))
		argsScanner.Split(scanFloat)
		for argsScanner.Scan() {
			arg, err := strconv.ParseFloat(argsScanner.Text(), 64)
			if err != nil {
				return nil, err
			}
			command.Arguments = append(command.Arguments, arg)
		}
		if err := argsScanner.Err(); err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return commands, nil
}

func args0(f []float64) error {
	if len(f) != 0 {
		return fmt.Errorf("expected 0 arguments, got %d: %v", len(f), f)
	}
	return nil
}

func args1(f []float64) (float64, []float64, error) {
	if len(f) < 1 {
		return 0, f, fmt.Errorf("expected 1 argument, got %d: %v", len(f), f)
	}
	return f[0], f[1:], nil
}

func args2(f []float64) (float64, float64, []float64, error) {
	if len(f) < 2 {
		return 0, 0, f, fmt.Errorf("expected 2 arguments, got %d: %v", len(f), f)
	}
	return f[0], f[1], f[2:], nil
}

func args4(f []float64) (float64, float64, float64, float64, []float64, error) {
	if len(f) < 4 {
		return 0, 0, 0, 0, f, fmt.Errorf("expected 4 arguments, got %d: %v", len(f), f)
	}
	return f[0], f[1], f[2], f[3], f[4:], nil
}

func args6(f []float64) (float64, float64, float64, float64, float64, float64, []float64, error) {
	if len(f) < 6 {
		return 0, 0, 0, 0, 0, 0, f, fmt.Errorf("expected 6 arguments, got %d: %v", len(f), f)
	}
	return f[0], f[1], f[2], f[3], f[4], f[5], f[6:], nil
}

func args7(f []float64) (float64, float64, float64, float64, float64, float64, float64, []float64, error) {
	if len(f) < 7 {
		return 0, 0, 0, 0, 0, 0, 0, f, fmt.Errorf("expected 7 arguments, got %d: %v", len(f), f)
	}
	return f[0], f[1], f[2], f[3], f[4], f[5], f[6], f[7:], nil
}

// PathCommands parses a sequence of path commands from the given reader
func PathCommands(r io.Reader) ([]svg.PathCommand, error) {
	raw, err := rawCommands(r)
	if err != nil {
		return nil, err
	}
	var commands []svg.PathCommand
	for _, r := range raw {
		var c svg.PathCommand
		upper := unicode.IsUpper(r.Code)
		args := r.Arguments
		var err error
		switch unicode.ToLower(r.Code) {
		case 'a':
			first := true
			for first || len(args) > 0 {
				first = false
				c.Arc = &svg.PathCommandArc{Upper: upper}
				c.Arc.Rx, c.Arc.Ry, c.Arc.XAxisRotation, c.Arc.LargeArcFlag, c.Arc.SweepFlag, c.Arc.X, c.Arc.Y, args, err = args7(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'm':
			first := true
			for first || len(args) > 0 {
				first = false
				c.MoveTo = &svg.PathCommandMoveTo{Upper: upper}
				c.MoveTo.X, c.MoveTo.Y, args, err = args2(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'l':
			first := true
			for first || len(args) > 0 {
				first = false
				c.LineTo = &svg.PathCommandLineTo{Upper: upper}
				c.LineTo.X, c.LineTo.Y, args, err = args2(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'h':
			first := true
			for first || len(args) > 0 {
				first = false
				c.H = &svg.PathCommandH{Upper: upper}
				c.H.X, args, err = args1(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'v':
			first := true
			for first || len(args) > 0 {
				first = false
				c.V = &svg.PathCommandV{Upper: upper}
				c.V.Y, args, err = args1(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'z':
			c.Z = &svg.PathCommandZ{Upper: upper}
			err = args0(args)
		case 'c':
			first := true
			for first || len(args) > 0 {
				first = false
				c.CubicBezier = &svg.PathCommandCubicBezier{Upper: upper}
				c.CubicBezier.X1, c.CubicBezier.Y1, c.CubicBezier.X2, c.CubicBezier.Y2, c.CubicBezier.X, c.CubicBezier.Y, args, err = args6(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 's':
			first := true
			for first || len(args) > 0 {
				first = false
				c.CubicBezierShortcut = &svg.PathCommandCubicBezierShortcut{Upper: upper}
				c.CubicBezierShortcut.X2, c.CubicBezierShortcut.Y2, c.CubicBezierShortcut.X, c.CubicBezierShortcut.Y, args, err = args4(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 'q':
			first := true
			for first || len(args) > 0 {
				first = false
				c.QuadraticBezier = &svg.PathCommandQuadraticBezier{Upper: upper}
				c.QuadraticBezier.X1, c.QuadraticBezier.Y1, c.QuadraticBezier.X, c.QuadraticBezier.Y, args, err = args4(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		case 't':
			first := true
			for first || len(args) > 0 {
				first = false
				c.QuadraticBezierShortcut = &svg.PathCommandQuadraticBezierShortcut{Upper: upper}
				c.QuadraticBezierShortcut.X, c.QuadraticBezierShortcut.Y, args, err = args2(args)
				if err != nil {
					break
				}
				commands = append(commands, c)
			}
		default:
			return nil, fmt.Errorf("unrecognized path command: %q", r.Code)
		}
		if err != nil {
			return nil, fmt.Errorf("command %q: %v", r.Code, err)
		}
	}
	return commands, nil
}
