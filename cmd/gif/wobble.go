package main

import (
	"image/color"
	"github.com/sgreben/yeetgif/pkg/imaging"
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
)

func CommandWobble(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	const (
		wobbleTypeSine   = "sine"
		wobbleTypeSnap   = "snap"
		wobbleTypeSmooth = "smooth"
	)
	var (
		f  = gifcmd.Float{Value: 1.0}
		a  = gifcmd.Float{Value: 20.0}
		ph = gifcmd.Float{Value: 0.0}
		t  = gifcmd.Enum{
			Choices: []string{
				wobbleTypeSine,
				wobbleTypeSnap,
			},
			Value: wobbleTypeSine,
		}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("p phase", &ph, "")
	cmd.VarOpt("t type", &t, "")
	cmd.Action = func() {
		frequency := f.Value
		amplitude := a.Value
		phase := ph.Value
		n := len(images)
		fs := map[string]func(int) float64{
			wobbleTypeSine: func(i int) float64 {
				return amplitude * math.Sin(2*math.Pi*phase+2*math.Pi*frequency*float64(i)/float64(n))
			},
			wobbleTypeSnap: func(i int) float64 {
				t := float64(i) / float64(n)
				y := math.Sin(2*math.Pi*phase + 2*math.Pi*frequency*t)
				y = math.Sin(y)
				return amplitude * y
			},
		}
		Wobble(images, fs[t.Value])
	}
}


// Wobble `images` `frequency` times by `amplitude` degrees
func Wobble(images []image.Image, f func(int) float64) {
	rotate := func(i int) {
		angle := f(i)
		bPre := images[i].Bounds()
		images[i] = imaging.Rotate(images[i], angle, color.Transparent)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		if !config.Pad {
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), rotate)
}
