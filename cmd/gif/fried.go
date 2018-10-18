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
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS]"
	var (
		a          = gifcmd.FloatsCSV{Values: []float64{0.33}}
		b          = gifcmd.FloatsCSV{Values: []float64{0.2}}
		c          = gifcmd.FloatsCSV{Values: []float64{0.9}}
		clip       = cmd.BoolOpt("clip", true, "")
		q          = gifcmd.FloatsCSV{Values: []float64{84}}
		w          = cmd.IntOpt("w walk", 10, "🌀")
		t          = gifcmd.FloatsCSV{Values: []float64{0.4}}
		n          = gifcmd.FloatsCSV{Values: []float64{1.0}}
		n1         = gifcmd.FloatsCSV{Values: []float64{0.02}}
		n2         = gifcmd.FloatsCSV{Values: []float64{0.5}}
		n3         = gifcmd.FloatsCSV{Values: []float64{0.1}}
		saturation = gifcmd.FloatsCSV{Values: []float64{3.0}}
		contrast   = gifcmd.FloatsCSV{Values: []float64{6.0}}
		iterations = cmd.IntOpt("i iterations", 1, "")
	)
	cmd.VarOpt("a", &a, "🅰️")
	cmd.VarOpt("b", &b, "🅱️")
	cmd.VarOpt("c", &c, "🆑")
	cmd.VarOpt("n noise", &n, "🌀️")
	cmd.VarOpt("noise1", &n1, "🌀️")
	cmd.VarOpt("noise2", &n2, "🌀️")
	cmd.VarOpt("noise3", &n3, "🌀")
	cmd.VarOpt("j jpeg", &q, "[0,100]")
	cmd.VarOpt("u saturation", &saturation, "")
	cmd.VarOpt("o contrast", &contrast, "")
	cmd.VarOpt("t tint", &t, "tint")
	cmd.Action = func() {
		if *iterations > 1 {
			for i, v := range t.Values {
				t.Values[i] = v / float64(*iterations)
			}
		}
		for i := 0; i < *iterations; i++ {
			Fried(
				images,
				t.PiecewiseLinear(0, 1),
				a.PiecewiseLinear(0, 1),
				b.PiecewiseLinear(0, 1),
				c.PiecewiseLinear(0, 1),
				q.PiecewiseLinear(0, 1),
				*w,
				saturation.PiecewiseLinear(0, 1),
				contrast.PiecewiseLinear(0, 1),
				n.PiecewiseLinear(0, 1),
				n1.PiecewiseLinear(0, 1),
				n2.PiecewiseLinear(0, 1),
				n3.PiecewiseLinear(0, 1),
				*clip,
			)
		}
	}
}

// Fried meme
func Fried(images []image.Image, tintF, aF, bF, cF, lossF func(float64) float64, step int, saturationF, contrastF, noiseF, noise1F, noise2F, noise3F func(float64) float64, clip bool) {
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
		t := float64(i) / float64(n)
		explodePoint := explodePoints[i]
		original := images[i]
		images[i] = imaging.FriedDistortion1(images[i], explodePoint, aF(t), bF(t), cF(t))
		exploded := images[i]
		images[i] = imaging.AdjustTint(images[i], tintF(t), orange)
		noise, noise1, noise2, noise3 := noiseF(t), noise1F(t), noise2F(t), noise3F(t)
		images[i] = imaging.AdjustNoiseHSL(images[i], noise*noise1, noise*noise2, noise*noise3)
		loss := int(lossF(t))
		if loss < 0 {
			loss = 0
		}
		if loss > 100 {
			loss = 100
		}
		jpeg(i, 100-loss)
		images[i] = imaging.AdjustSaturation(images[i], saturationF(t))
		images[i] = imaging.AdjustSigmoid(images[i], 0.5, contrastF(t))
		jpeg(i, 100-(loss/2))
		if clip {
			images[i] = imaging.OverlayWithOp(images[i], original, image.ZP, imaging.OpReplaceAlpha)
		}
		images[i] = imaging.OverlayWithOp(images[i], exploded, image.ZP, imaging.OpMinAlpha)
	}
	parallel(len(images), fry, "fry")
}
