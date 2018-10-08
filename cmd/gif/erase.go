package main

import (
	"image"
	"image/color"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandErase(cmd *cli.Cmd) {
	var (
		t  = gifcmd.Float{Value: 0.2}
		x  = cmd.IntOpt("x sample-x", 3, "")
		y  = cmd.IntOpt("y sample-y", 3, "")
		wh = gifcmd.Float{Value: 1.0}
		ws = gifcmd.Float{Value: 0.5}
		wl = gifcmd.Float{Value: 1.0}
	)
	//sample = cmd.BoolOpt("sample", true, "")
	cmd.VarOpt("t tolerance", &t, "")
	cmd.VarOpt("u", &wh, "")
	cmd.VarOpt("s", &ws, "")
	cmd.VarOpt("l", &wl, "")
	cmd.Action = func() {
		Erase(images, *x, *y, t.Value, wh.Value, ws.Value, wl.Value)
	}
}

func Erase(images []image.Image, x, y int, t, wh, ws, wl float64) {
	sqr := func(x float64) float64 { return x * x }
	erase := func(i int) {
		sample := images[i].At(x, y)
		r, g, b, _ := sample.RGBA()
		sh, ss, sl, _ := imaging.HSLA(color.RGBA{uint8(r / 0xff), uint8(g / 0xff), uint8(b / 0xff), 0})
		images[i] = imaging.AdjustHSLAFunc(images[i], func(h, s, l, a *float64) {
			dist := math.Sqrt((wh*sqr(*h-sh) + ws*sqr(*l-sl) + wl*sqr(*s-ss)) / (wh + ws + wl))
			if dist < t/2 {
				*a = 0
				return
			}
			if dist < t {
				*a = *a * ((dist - t/2) / (t / 2))
			}
		})
	}
	parallel(len(images), erase)
}
