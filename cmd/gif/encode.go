package main

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"time"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	"github.com/ericpauley/go-quantize/quantize"

	"image/gif"

	_ "image/jpeg"
)

func EncodeRaw(w io.Writer, images []image.Image) error {
	tarUname := "yeetgif"
	tarGname := "eggplant"
	tarModTime := time.Now()
	rawBytes := make([][]byte, len(images))
	encodePNG := func(i int) {
		buf := &bytes.Buffer{}
		png.Encode(buf, images[i])
		rawBytes[i] = buf.Bytes()
	}
	parallel(len(images), encodePNG, "encode (raw)")
	tarW := tar.NewWriter(w)
	buf := &bytes.Buffer{}
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
	if err := json.NewEncoder(buf).Encode(meta); err != nil {
		return err
	}
	metadataBytes := buf.Bytes()
	if err := tarW.WriteHeader(&tar.Header{
		Name:    rawMetadataFileName,
		Mode:    0600,
		Size:    int64(len(metadataBytes)),
		Uname:   tarUname,
		Gname:   tarGname,
		ModTime: tarModTime,
	}); err != nil {
		return err
	}
	if _, err := tarW.Write(metadataBytes); err != nil {
		return err
	}
	for i, b := range rawBytes {
		if err := tarW.WriteHeader(&tar.Header{
			Name:    fmt.Sprintf("%d.png", i),
			Mode:    0600,
			Size:    int64(len(b)),
			Uname:   tarUname,
			Gname:   tarGname,
			ModTime: tarModTime,
		}); err != nil {
			return err
		}
		if _, err := tarW.Write(b); err != nil {
			return err
		}
		if err := tarW.Flush(); err != nil {
			return err
		}
	}
	if err := tarW.Close(); err != nil {
		return err
	}
	return nil
}

// Encode the `images` as a GIF to `w`
func Encode(w io.Writer, images []image.Image) error {
	b := Bounds(images)
	out := gif.GIF{
		LoopCount: 0,
		Config: image.Config{
			Width:  b.Dx(),
			Height: b.Dy(),
		},
	}
	quantizer := quantize.MedianCutQuantizer{
		AddTransparent: true,
		Aggregation:    quantize.Mean,
	}
	out.Image = make([]*image.Paletted, len(images))
	out.Delay = make([]int, len(images))
	out.Disposal = make([]byte, len(images))
	n := float64(len(images))
	quantize := func(i int) {
		t := float64(i) / n
		delay := int(config.DelayMilliseconds(t) / 10)
		out.Delay[i] = delay
		out.Disposal[i] = gif.DisposalBackground
		switch img := images[i].(type) {
		case *image.Paletted:
			out.Image[i] = img
		default:
			palette := quantizer.Quantize(make([]color.Color, 0, 256), images[i])
			imgQuantized := image.NewPaletted(images[i].Bounds(), palette)
			out.Image[i] = imgQuantized
		}
	}
	parallel(len(images), quantize, "quantize")
	draw := func(i int) {
		switch images[i].(type) {
		case *image.Paletted:
			// do nothing
		default:
			draw.FloydSteinberg.Draw(out.Image[i], images[i].Bounds(), images[i], image.ZP)
		}
	}
	parallel(len(images), draw, "draw")

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
