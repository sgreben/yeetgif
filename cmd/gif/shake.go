package main

import (
	"image"
	"math"
	"math/rand"

	"github.com/sgreben/piecewiselinear"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandShake(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS]"
	var (
		random = gifcmd.FloatsCSV{Values: []float64{0.75}}
		f      = gifcmd.FloatsCSV{Values: []float64{1}}
		a      = gifcmd.FloatsCSV{Values: []float64{7}}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("r random", &random, "🌀")
	cmd.Action = func() {
		randomF := piecewiselinear.Function{Y: random.Values}
		randomF.X = piecewiselinear.Span(0, 1, len(randomF.Y))
		amplitudeF := piecewiselinear.Function{Y: a.Values}
		amplitudeF.X = piecewiselinear.Span(0, 1, len(amplitudeF.Y))
		frequencyF := piecewiselinear.Function{Y: f.Values}
		frequencyF.X = piecewiselinear.Span(0, 1, len(frequencyF.Y))
		Shake(images, randomF.At, frequencyF.At, amplitudeF.At)
	}
}

// Shake `images`
func Shake(images []image.Image, random, frequency, amplitude func(float64) float64) {
	n := float64(len(images))
	phaseY := math.Pi / 2
	move := func(i int) {
		t := float64(i) / n
		r := random(t)
		rX, rY := rand.Float64()*r, rand.Float64()*r
		a := amplitude(t)
		f := frequency(t)
		offset := image.Point{
			X: int(a * math.Sin(2*math.Pi*f*t+(rX*2*math.Pi))),
			Y: int(a * math.Sin(2*math.Pi*f*t+phaseY+(rY*2*math.Pi))),
		}
		if !config.Pad {
			images[i] = imaging.Paste(image.NewNRGBA(images[i].Bounds()), images[i], offset)
			return
		}
		bounds := images[i].Bounds()
		bounds.Min.X -= int(a)
		bounds.Max.X += int(a)
		bounds.Min.Y -= int(a)
		bounds.Max.Y += int(a)
		images[i] = imaging.Paste(image.NewNRGBA(bounds), images[i], offset)
	}
	parallel(len(images), move)
}
