package main

import (
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandTint(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	var (
		f    = gifcmd.Float{Value: 1.0}
		a    = gifcmd.Float{Value: 0.95}
		from = gifcmd.Float{Value: 0.7}
		to   = gifcmd.Float{Value: 0.9}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("0 from", &from, "")
	cmd.VarOpt("1 to", &to, "")
	cmd.VarOpt("i intensity", &a, "")
	cmd.Action = func() {
		TintPulse(images, f.Value, a.Value, from.Value, to.Value)
	}
}

func TintPulse(images []image.Image, frequency, weight, from, to float64) {
	n := float64(len(images))
	dist := math.Min(to-from, from+(1-to))
	if dist == 0 && to != from {
		dist = 1.0
	}
	mid := from + (dist / 2)
	if mid < 0 {
		mid++
	}
	if mid > 1 {
		mid--
	}
	tint := func(i int) {
		hue := mid + (dist * math.Sin(2*math.Pi*frequency*float64(i)/n) / 2)
		if hue < 0 {
			hue++
		}
		if hue > 1 {
			hue--
		}
		images[i] = imaging.AdjustHue(images[i], weight, hue)
	}
	parallel(len(images), tint)
}
