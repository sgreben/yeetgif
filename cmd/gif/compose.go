package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"sync"

	"github.com/sgreben/yeetgif/pkg/gifmath"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandCompose(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS] INPUT"
	const (
		orderUnder = "under"
		orderOver  = "over"
	)
	const (
		positionCenter   = "center"
		positionLeft     = "left"
		positionRight    = "right"
		positionTop      = "top"
		positionBottom   = "bottom"
		positionAbsolute = "abs"
	)
	var (
		input = cmd.StringArg("INPUT", "", "")
		z     = gifcmd.Enum{
			Choices: []string{orderUnder, orderOver},
			Value:   orderOver,
		}
		p = gifcmd.Enum{
			Choices: []string{
				positionCenter,
				positionLeft,
				positionRight,
				positionTop,
				positionBottom,
				positionAbsolute,
			},
			Value: positionCenter,
		}
		x = gifcmd.FloatsCSV{Values: []float64{0}}
		y = gifcmd.FloatsCSV{Values: []float64{0}}
		s = gifcmd.FloatsCSV{Values: []float64{1.0}}
	)
	cmd.VarOpt("x", &x, "")
	cmd.VarOpt("y", &y, "")
	cmd.VarOpt("z z-order", &z, z.Help())
	cmd.VarOpt("p position", &p, p.Help())
	cmd.VarOpt("s scale", &s, "")
	cmd.Action = func() {
		f, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		layer := Decode(f)
		scale := s.PiecewiseLinear(0, 1)
		var imageAnchor, layerAnchor imaging.Anchor
		switch p.Value {
		case positionAbsolute:
			imageAnchor = imaging.TopLeft
			layerAnchor = imaging.Center
		case positionCenter:
			imageAnchor = imaging.Center
			layerAnchor = imaging.Center
		case positionLeft:
			imageAnchor = imaging.Left
			layerAnchor = imaging.Right
		case positionRight:
			imageAnchor = imaging.Right
			layerAnchor = imaging.Left
		case positionTop:
			imageAnchor = imaging.Top
			layerAnchor = imaging.Bottom
		case positionBottom:
			imageAnchor = imaging.Bottom
			layerAnchor = imaging.Top
		}
		xF := x.PiecewiseLinear(0, 1)
		yF := y.PiecewiseLinear(0, 1)
		switch z.Value {
		case orderOver:
			Compose(images, layer, xF, yF, scale, imageAnchor, layerAnchor)
		case orderUnder:
			Compose(layer, images, xF, yF, scale, layerAnchor, imageAnchor)
		}
	}
}

func Compose(a, b []image.Image, xF, yF, sF func(float64) float64, anchorA, anchorB imaging.Anchor) {
	if len(a) == 0 || len(b) == 0 {
		return
	}
	n := gifmath.LCM(len(a), len(b))
	var (
		boundsMu   sync.Mutex
		bounds     image.Rectangle
		overOffset = make([]image.Point, n)
	)
	bound := func(i int) {
		t := float64(i) / float64(n)
		s := sF(t)
		under := a[i%len(a)]
		overBounds := b[i%len(b)].Bounds()
		over := imaging.New(int(float64(overBounds.Dx())*s), int(float64(overBounds.Dy())*s), color.Transparent)
		x, y := xF(t), yF(t)
		p := image.Point{X: int(x), Y: int(y)}
		overOffset[i] = imaging.AnchorPoint(under, anchorA).
			Sub(imaging.AnchorPoint(over, anchorB)).
			Add(p)
		boundsMu.Lock()
		bounds = bounds.Union(under.Bounds().Union(over.Bounds().Add(overOffset[i])))
		boundsMu.Unlock()
	}
	parallel(n, bound)
	images = make([]image.Image, n)
	compose := func(i int) {
		under := a[i%len(a)]
		t := float64(i) / float64(n)
		s := sF(t)
		over := ResizeScale1(b[i%len(b)], s)
		bg := imaging.New(bounds.Dx(), bounds.Dy(), color.Transparent)
		bg = imaging.Paste(bg, under, image.ZP.Sub(bounds.Min))
		images[i] = imaging.Overlay(bg, over, overOffset[i].Sub(bounds.Min), 1.0)
	}
	parallel(n, compose, "compose")
}
