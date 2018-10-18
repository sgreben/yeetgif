package main

import (
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandHue(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS] [HUE_OFFSETS_CSV]"
	const (
		hueModeMul = "mul"
		hueModeAdd = "add"
		hueModeSub = "sub"
		hueModePow = "pow"
		hueModeSin = "sin"
	)
	var (
		f    = gifcmd.Float{Value: 1.0}
		a    = gifcmd.Float{Value: 0.1}
		x    = gifcmd.FloatsCSV{Values: []float64{1.0}}
		y    = gifcmd.FloatsCSV{Values: []float64{1.0}}
		mode = gifcmd.Enum{
			Choices: []string{
				hueModeMul,
				hueModeAdd,
				hueModeSub,
				hueModePow,
				hueModeSin,
			},
			Value: hueModeAdd,
		}
		custom = gifcmd.FloatsCSV{}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("x", &x, "")
	cmd.VarOpt("y", &y, "")
	cmd.VarOpt("m mode", &mode, mode.Help())
	cmd.VarArg("HUE_OFFSETS_CSV", &custom, "")
	cmd.Action = func() {
		var hueF func(float64) float64
		switch {
		case len(custom.Texts) > 0:
			hueF = custom.PiecewiseLinear(0, 1)
		default:
			hueF = func(t float64) float64 {
				return a.Value * math.Sin(math.Pi+2*math.Pi*f.Value*t) / 2
			}
		}
		xF := x.PiecewiseLinear(0, 1)
		yF := y.PiecewiseLinear(0, 1)
		var modeF func(float64, float64, float64) float64
		switch mode.Value {
		case hueModeMul:
			modeF = func(a, b, c float64) float64 {
				return a * b * c
			}
		case hueModeAdd:
			modeF = func(a, b, c float64) float64 {
				return a + b + c
			}
		case hueModeSub:
			modeF = func(a, b, c float64) float64 {
				return a - b - c
			}
		case hueModePow:
			modeF = func(a, b, c float64) float64 {
				return math.Pow(a, b*c)
			}
		case hueModeSin:
			modeF = func(a, b, c float64) float64 {
				return a + math.Sin(2*math.Pi*(b+c))
			}
		}
		HuePulse(images, func(x, y, t float64) float64 {
			return modeF(hueF(t), xF(x), yF(y))
		})
	}
}

func HuePulse(images []image.Image, f func(float64, float64, float64) float64) {
	n := float64(len(images))
	hue := func(i int) {
		t := float64(i) / n
		b := images[i].Bounds()
		w := float64(b.Dx())
		h := float64(b.Dy())
		images[i] = imaging.AdjustHueRotate(images[i], func(x, y int) float64 {
			return f(float64(x-b.Min.X)/w, float64(y-b.Min.Y)/h, t)
		})
	}
	parallel(len(images), hue, "hue")
}
