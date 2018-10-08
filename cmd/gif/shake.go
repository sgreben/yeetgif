package main

import (
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
		random = gifcmd.Float{Value: 0.5}
		f      = gifcmd.Float{Value: 1.0}
		a      = gifcmd.Float{Value: 8.0}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("r random", &random, "🌀")
	cmd.Action = func() {
		Shake(images, random.Value, f.Value, a.Value)
	}
}

// Shake `images`
func Shake(images []image.Image, random, frequency, amplitude float64) {
	n := len(images)
	phaseY := math.Pi / 2
	move := func(i int) {
		rX, rY := rand.Float64(), rand.Float64()
		offset := image.Point{
			X: int(amplitude * math.Sin(2*math.Pi*frequency*float64(i)/float64(n)+(rX*2*math.Pi))),
			Y: int(amplitude * math.Sin(2*math.Pi*frequency*float64(i)/float64(n)+phaseY+(rY*2*math.Pi))),
		}
		if !config.Pad {
			images[i] = imaging.Paste(image.NewNRGBA(images[i].Bounds()), images[i], offset)
			return
		}
		bounds := images[i].Bounds()
		bounds.Min.X -= int(amplitude)
		bounds.Max.X += int(amplitude)
		bounds.Min.Y -= int(amplitude)
		bounds.Max.Y += int(amplitude)
		images[i] = imaging.Paste(image.NewNRGBA(bounds), images[i], offset)
	}
	parallel(len(images), move)
}
