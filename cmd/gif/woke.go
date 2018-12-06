package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/gifstatic"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandWoke(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS] POINTS"
	const (
		typeFull   = "full"
		typeCenter = "center"
	)
	var (
		random    = gifcmd.Float{Value: 0.5}
		s         = gifcmd.Float{Value: 0.9}
		hue       = gifcmd.Float{Value: 0.8}
		lightness = gifcmd.Float{Value: 1.0}
		alpha     = gifcmd.Float{Value: 0.8}
		ap        = gifcmd.Float{Value: 2.0}
		at        = gifcmd.Float{Value: 0.15}
		clip      = cmd.BoolOpt("c clip", true, "clip flares to image alpha")
		t         = gifcmd.Enum{
			Choices: []string{typeFull, typeCenter},
			Value:   typeFull,
		}
		flares    = [][2]int{}
		flaresVar = gifcmd.JSON{Value: &flares}
	)
	cmd.VarOpt("t type", &t, "")
	cmd.VarOpt("s scale", &s, "")
	cmd.VarOpt("u hue", &hue, "")
	cmd.VarOpt("l lightness", &lightness, "")
	cmd.VarOpt("a alpha", &alpha, "")
	cmd.VarOpt("p alpha-pow", &ap, "")
	cmd.VarOpt("alpha-threshold", &at, "")
	cmd.VarOpt("r random", &random, "🌀")
	cmd.VarArg("POINTS", &flaresVar, `flare locations, JSON, e.g. "[[123,456],[-100,23]]"`)
	cmd.Action = func() {
		var flarePoints []image.Point
		for _, p := range flares {
			flarePoints = append(flarePoints, image.Point{X: p[0], Y: p[1]})
		}
		changeHue := hue.Text != ""
		var flare image.Image
		switch t.Value {
		case typeFull:
			flare = gifstatic.LensFlare
		case typeCenter:
			flare = gifstatic.LensFlareCenter
		}
		b := flare.Bounds()
		width := int(float64(b.Dx()) * s.Value)
		height := int(float64(b.Dy()) * s.Value)
		flare = imaging.Resize(flare, width, height, imaging.Lanczos)
		alphaT := at.Value
		alphaP := ap.Value
		alphaV := alpha.Value
		flare = imaging.AdjustHSLAFunc(flare, func(_, _ int, h, s, l, a *float64, _ *int) {
			*a = math.Pow(*a, alphaP) * alphaV
			if *a < alphaT {
				*a = 0
				return
			}
			*a = (*a - alphaT) / (1.0 - alphaT)
			*l = (*l) * lightness.Value
			if changeHue {
				*h += hue.Value
			}
		})
		Woke(images, flare, random.Value, flarePoints, *clip)
	}
}

// Woke flares
func Woke(images []image.Image, flare image.Image, random float64, flares []image.Point, clip bool) {
	woke := func(i int) {
		b := images[i].Bounds()
		layer := imaging.New(b.Dx(), b.Dy(), color.Transparent)
		flip := false // (i % 2) == 0
		for _, p := range flares {
			flare := flare
			if random > 0 {
				flare = imaging.Rotate(flare, rand.Float64()*random*360, color.Transparent)
			}
			if flip {
				flare = imaging.FlipV(imaging.FlipH(flare))
			}
			b := flare.Bounds()
			layer = imaging.OverlayWithOp(layer, flare, image.Point{
				X: (p.X - b.Dx()/2),
				Y: (p.Y - b.Dy()/2),
			}, imaging.OpLighten)
		}
		woke := imaging.Overlay(images[i], layer, image.ZP, 1.0)
		if clip {
			images[i] = imaging.OverlayWithOp(
				woke,
				images[i],
				image.ZP,
				imaging.OpReplaceAlpha,
			)
			return
		}
		images[i] = woke
	}
	parallel(len(images), woke, "woke")
}
