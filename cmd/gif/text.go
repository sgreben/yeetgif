package main

import (
	"image"
	"image/color"
	"sync"

	"github.com/sgreben/yeetgif/pkg/gifstatic"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

var (
	textFont, _ = truetype.Parse(gifstatic.RobotoTTF)
)

func CommandText(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS] [TEXT]"
	var (
		a    = gifcmd.Float{Value: 0.7}
		s    = gifcmd.Float{Value: 18.5}
		y    = gifcmd.Float{Value: 0.30}
		padY = gifcmd.Float{Value: 3}
	)
	cmd.VarOpt("a background-alpha", &a, "")
	cmd.VarOpt("s font-size", &s, "")
	cmd.VarOpt("y text-y", &y, "")
	cmd.VarOpt("p background-padding", &padY, "")
	text := cmd.StringArg("TEXT", "#yeetgif", "")
	cmd.Action = func() {
		Text(images, a.Value, padY.Value, s.Value, y.Value, *text)
	}
}

func Text(images []image.Image, alpha, padY, sizePoints, textYFactor float64, text string) {
	fontFace := truetype.NewFace(textFont, &truetype.Options{Size: sizePoints})
	var fontMutex sync.Mutex
	padYBottomExtra := 6.0
	write := func(i int) {
		w, h := images[i].Bounds().Dx(), images[i].Bounds().Dy()
		ctx := gg.NewContext(w, h)
		ctx.SetFontFace(fontFace)
		ctx.SetColor(color.White)
		fontMutex.Lock()
		_, textHeight := ctx.MeasureString(text)
		textY := float64(h) * (1 - textYFactor)
		ctx.DrawStringAnchored(text, float64(w)/2, textY, 0.5, 0.5)
		fontMutex.Unlock()
		textBackgroundHeight := textHeight + 2*padY + padYBottomExtra
		textBackgroundY := textY - (textBackgroundHeight-padYBottomExtra)/2
		textBackground := imaging.New(w, int(textBackgroundHeight), color.Black)
		images[i] = imaging.Overlay(images[i], textBackground, image.Point{Y: int(textBackgroundY)}, alpha)
		images[i] = imaging.Overlay(images[i], ctx.Image(), image.ZP, 1.0)
	}
	parallel(len(images), write)
}
