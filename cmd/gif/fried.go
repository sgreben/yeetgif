package main

import (
	"bytes"
	"image"
	"image/color"
	"math/rand"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandFried(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS]"
	var (
		a          = gifcmd.Float{Value: 0.33}
		b          = gifcmd.Float{Value: 0.2}
		c          = gifcmd.Float{Value: 0.9}
		clip       = cmd.BoolOpt("clip", true, "")
		q          = cmd.IntOpt("j jpeg", 84, "[0,100]")
		w          = cmd.IntOpt("w walk", 10, "🌀")
		t          = gifcmd.Float{Value: 0.4}
		n          = gifcmd.Float{Value: 1.0}
		n1         = gifcmd.Float{Value: 0.02}
		n2         = gifcmd.Float{Value: 0.5}
		n3         = gifcmd.Float{Value: 0.1}
		saturation = gifcmd.Float{Value: 3.0}
		contrast   = gifcmd.Float{Value: 6.0}
		iterations = cmd.IntOpt("i iterations", 1, "")
	)
	cmd.VarOpt("a", &a, "🅰️")
	cmd.VarOpt("b", &b, "🅱️")
	cmd.VarOpt("c", &c, "🆑")
	cmd.VarOpt("n noise", &n, "🌀️")
	cmd.VarOpt("noise1", &n1, "🌀️")
	cmd.VarOpt("noise2", &n2, "🌀️")
	cmd.VarOpt("noise3", &n3, "🌀")
	cmd.VarOpt("u saturation", &saturation, "")
	cmd.VarOpt("o contrast", &contrast, "")
	cmd.VarOpt("t tint", &t, "tint")
	cmd.Action = func() {
		n1.Value *= n.Value
		n2.Value *= n.Value
		n3.Value *= n.Value
		if *q > 100 {
			*q = 100
		}
		if *q < 0 {
			*q = 0
		}
		for i := 0; i < *iterations; i++ {
			Fried(images, t.Value, a.Value, b.Value, c.Value, *q, *w, saturation.Value, contrast.Value, n1.Value, n2.Value, n3.Value, *clip)
		}
	}
}

// Fried meme
func Fried(images []image.Image, tint, a, b, c float64, loss, step int, saturation, contrast, noise1, noise2, noise3 float64, clip bool) {
	if loss < 0 {
		loss = 0
	}
	if loss > 100 {
		loss = 100
	}
	jpeg := func(i, quality int) {
		buf := &bytes.Buffer{}
		imaging.Encode(buf, images[i], imaging.JPEG, imaging.JPEGQuality(quality))
		images[i], _, _ = image.Decode(buf)
	}
	orange := color.RGBA{
		R: 255,
		G: 30,
		B: 0,
	}
	bounds := images[0].Bounds()
	explodePoint := image.Point{
		X: int(rand.Float64() * float64(bounds.Dx())),
		Y: int(rand.Float64() * float64(bounds.Dy())),
	}
	n := len(images)
	explodePoints := make([]image.Point, n)
	for i := 0; i <= n/2; i++ {
		explodePoints[i] = explodePoint
		explodePoints[n-1-i] = explodePoint
		explodePoint.X += int(rand.Float64()*2*float64(step)) - step
		explodePoint.Y += int(rand.Float64()*2*float64(step)) - step
	}
	fry := func(i int) {
		explodePoint := explodePoints[i]
		original := images[i]
		images[i] = imaging.Ripples(images[i], explodePoint, a, b, c)
		exploded := images[i]
		images[i] = imaging.AdjustTint(images[i], tint, orange)
		images[i] = imaging.AdjustNoiseHSL(images[i], noise1, noise2, noise3)
		jpeg(i, 100-loss)
		images[i] = imaging.AdjustSaturation(images[i], saturation)
		images[i] = imaging.AdjustSigmoid(images[i], 0.5, contrast)
		jpeg(i, 100-(loss/2))
		if clip {
			images[i] = imaging.OverlayWithOp(images[i], original, image.ZP, imaging.OpReplaceAlpha)
		}
		images[i] = imaging.OverlayWithOp(images[i], exploded, image.ZP, imaging.OpMinAlpha)
	}
	parallel(len(images), fry)
}
