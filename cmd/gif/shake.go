package main

import (
	"github.com/sgreben/yeetgif/pkg/piecewiselinear"
	"image"
	"math"
	"math/rand"

	"github.com/sgreben/yeetgif/pkg/imaging"
	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
)

func CommandShake(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	var (
		random = gifcmd.Float{Value: 0.75}
		f      = gifcmd.FloatsCSV{Values: []float64{1}}
		a      = gifcmd.FloatsCSV{Values: []float64{7}}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("r random", &random, "🌀")
	cmd.Action = func() {
		amplitudeF := piecewiselinear.Function{}
		k := float64(len(a.Values) - 1)
		for i := 0; i < len(a.Values); i++ {
			amplitudeF.X = append(amplitudeF.X, float64(i)/k)
			amplitudeF.Y = append(amplitudeF.Y, float64((a.Values)[i]))
		}
		frequencyF := piecewiselinear.Function{}
		k = float64(len(f.Values) - 1)
		for i := 0; i < len(f.Values); i++ {
			frequencyF.X = append(frequencyF.X, float64(i)/k)
			frequencyF.Y = append(frequencyF.Y, float64((f.Values)[i]))
		}
		Shake(images, random.Value, frequencyF.At, amplitudeF.At)
	}
}

// Shake `images`
func Shake(images []image.Image, random float64, frequency , amplitude func(float64)float64) {
	n := float64(len(images))
	phaseY := math.Pi / 2
	move := func(i int) {
		rX, rY := rand.Float64(), rand.Float64()
		t := float64(i)/n
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
