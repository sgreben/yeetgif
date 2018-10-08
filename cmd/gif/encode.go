package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"
	"time"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	"github.com/ericpauley/go-quantize/quantize"

	"image/gif"

	_ "image/jpeg"
)

// Encode the `images` as a GIF to `w`
func Encode(w io.Writer, images []image.Image) error {
	var maxWidth, maxHeight int
	for _, img := range images {
		width := img.Bounds().Dx()
		if width > maxWidth {
			maxWidth = width
		}
		height := img.Bounds().Dy()
		if height > maxHeight {
			maxHeight = height
		}
	}

	out := gif.GIF{
		LoopCount: 0,
		Config: image.Config{
			Width:  maxWidth,
			Height: maxHeight,
		},
	}
	quantizer := quantize.MedianCutQuantizer{
		AddTransparent: true,
		Aggregation:    quantize.Mean,
	}
	out.Image = make([]*image.Paletted, len(images))
	out.Delay = make([]int, len(images))
	out.Disposal = make([]byte, len(images))
	quantize := func(i int) {
		b := images[i].Bounds()
		palette := quantizer.Quantize(make([]color.Color, 0, 256), images[i])
		imgQuantized := image.NewPaletted(b, palette)
		draw.FloydSteinberg.Draw(imgQuantized, b, images[i], image.ZP)
		out.Image[i] = imgQuantized
		out.Delay[i] = config.DelayMilliseconds / 10
		out.Disposal[i] = gif.DisposalBackground
	}
	parallel(len(images), quantize)

	if !config.WriteMeta {
		return gif.EncodeAll(w, &out)
	}
	buf := &bytes.Buffer{}
	gif.EncodeAll(buf, &out)
	metaJSONBytes, _ := json.Marshal(metaEntry{
		AppName:   appName,
		Timestamp: time.Now().Format(time.RFC3339),
		Args:      os.Args[1:],
		Version:   version,
	})
	meta = append(meta, gifmeta.Extension{
		Type:   gifmeta.Comment,
		Blocks: gifmeta.Blocks(metaJSONBytes),
	})
	return gifmeta.Append(w, buf, meta...)
}
