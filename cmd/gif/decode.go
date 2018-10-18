package main

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"sync"

	"github.com/sgreben/yeetgif/pkg/gifmeta"
)

const rawMetadataFileName = "metadata.json"

// Decode images from `r`, add their meta to global meta
func Decode(r io.Reader) []image.Image {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("read: %v", err)
	}
	if ShouldDecodeRaw(data) {
		rawImages, rawMeta, err := DecodeRaw(data)
		if err != nil {
			log.Printf("decode (raw): %v", err)
		}
		meta = append(meta, rawMeta...)
		return rawImages
	}
	stdImages, stdMeta, err := DecodeStandard(data)
	if err != nil {
		log.Printf("decode: %v", err)
	}
	meta = append(meta, stdMeta...)
	return stdImages
}

func DecodeStandard(data []byte) ([]image.Image, []gifmeta.Extension, error) {
	var images []image.Image
	var meta []gifmeta.Extension
	var errors []error

	r := bytes.NewReader(data)
	peekBuf := &bytes.Buffer{}
	tee := io.TeeReader(r, peekBuf)
	for r.Len() > 0 {
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
				err = fmt.Errorf("read gif meta: %v", err)
				errors = append(errors, err)
			}
			continue
		}
		r.Seek(-n, io.SeekCurrent)
		img, _, err := image.Decode(tee)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		images = append(images, img)
	}
	var err error
	if len(errors) > 0 {
		err = fmt.Errorf("%v", err)
	}
	return images, meta, err
}

func ShouldDecodeRaw(data []byte) bool {
	_, err := tar.NewReader(bytes.NewReader(data)).Next()
	return err == nil
}

func DecodeRaw(data []byte) ([]image.Image, []gifmeta.Extension, error) {
	var images []image.Image
	var meta []gifmeta.Extension
	var errors []error

	var errorsMu sync.Mutex
	var imageBytes [][]byte
	tarR := tar.NewReader(bytes.NewReader(data))
	for {
		hdr, err := tarR.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, err)
			break
		}
		if hdr.Name == rawMetadataFileName {
			if err := json.NewDecoder(tarR).Decode(&meta); err != nil {
				errors = append(errors, err)
			}
			continue
		}
		b, err := ioutil.ReadAll(tarR)
		if err != nil {
			errors = append(errors, err)
			break
		}
		imageBytes = append(imageBytes, b)
	}
	images = make([]image.Image, len(imageBytes))
	decodeSingle := func(i int) {
		img, _, err := image.Decode(bytes.NewReader(imageBytes[i]))
		if err != nil {
			err := fmt.Errorf("decode: %v", err)
			errorsMu.Lock()
			errors = append(errors, err)
			errorsMu.Unlock()
			return
		}
		images[i] = img
	}
	parallel(len(images), decodeSingle, "decode (raw)")
	var err error
	if len(errors) > 0 {
		err = fmt.Errorf("%v", err)
	}
	return images, meta, err
}
