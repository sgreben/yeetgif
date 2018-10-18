package main

import (
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandPulse(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS] [ZOOM_LEVELS_CSV]"
	var (
		from   = gifcmd.Float{Value: 0.9}
		f      = gifcmd.Float{Value: 1.0}
		ph     = gifcmd.Float{Value: 0.0}
		to     = gifcmd.Float{Value: 1.0}
		custom = gifcmd.FloatsCSV{}
	)
	cmd.VarOpt("0 from", &from, "")
	cmd.VarOpt("1 to", &to, "")
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("p phase", &ph, "")
	cmd.VarArg("ZOOM_LEVELS_CSV", &custom, "")
	cmd.Action = func() {
		if len(custom.Texts) > 0 {
			Pulse(images, custom.PiecewiseLinear(0, 1))
			return
		}
		frequency := f.Value
		phase := ph.Value
		left := from.Value
		right := to.Value
		Pulse(images, func(t float64) float64 {
			weight := math.Sin(2*math.Pi*phase + 2*math.Pi*frequency*t)
			return left*weight + right*(1-weight)
		})
	}
}

// Pulse `images` `frequency` times between scales `from` and `to`
func Pulse(images []image.Image, f func(float64) float64) {
	n := float64(len(images))
	scale := func(i int) {
		scale := f(float64(i) / n)
		bPre := images[i].Bounds()
		width := float64(bPre.Dx()) * scale
		height := float64(bPre.Dy()) * scale
		images[i] = imaging.Resize(images[i], int(width), int(height), imaging.Lanczos)
		if !config.Pad {
			bPost := images[i].Bounds()
			offset := image.Point{
				X: (bPost.Dx() - bPre.Dx()) / 2,
				Y: (bPost.Dy() - bPre.Dy()) / 2,
			}
			bPre.Min = bPre.Min.Add(offset)
			bPre.Max = bPre.Max.Add(offset)
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), scale, "pulse")
}
