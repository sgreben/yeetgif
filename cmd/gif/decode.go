package main

import (
	"bytes"
	"image"
	"image/gif"
	"io"
	"log"

	"github.com/sgreben/yeetgif/pkg/gifmeta"
)

// Decode images from `r`
func Decode(r io.Reader) []image.Image {
	var images []image.Image
	input := &bytes.Buffer{}
	_, err := io.Copy(input, r)
	if err != nil {
		log.Fatalf("read: %v", err)
	}
	seekableReader := bytes.NewReader(input.Bytes())
	peekBuf := &bytes.Buffer{}
	tee := io.TeeReader(seekableReader, peekBuf)
	for seekableReader.Len() > 0 {
		peekBuf.Reset()
		gif, err := gif.DecodeAll(tee)
		n := int64(peekBuf.Len())
		if err == nil {
			for _, img := range gif.Image {
				images = append(images, img)
			}
			moreMeta, err := gifmeta.Read(peekBuf, func(e *gifmeta.Extension) bool {
				return e.Type == gifmeta.Comment
			})
			meta = append(meta, moreMeta...)
			if err != nil {
				log.Printf("read gif meta: %v", err)
			}
			continue
		}
		seekableReader.Seek(-n, io.SeekCurrent)
		img, _, err := image.Decode(seekableReader)
		if err != nil {
			continue
		}
		images = append(images, img)
	}
	return images
}
