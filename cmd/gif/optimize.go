package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"

	cli "github.com/jawher/mow.cli"
)

func CommandOptimize(cmd *cli.Cmd) {
	cmd.Before = Input
	cmd.Spec = "[OPTIONS]"
	var (
		k = cmd.IntOpt("kb", 128, "target file size (KB)")
		n = cmd.BoolOpt("n no-resize", false, "don't resize the image")
		w = cmd.IntOpt("x width", 0, "target width (pixels)")
		h = cmd.IntOpt("y height", 0, "target height (pixels)")
	)
	cmd.Action = func() {
		if !*n && *w == 0 && *h == 0 {
			*w = 128
			*h = 128
		}
		if *n {
			*w = 0
			*h = 0
		}
		Optimize(images, int64(*k), *w, *h)
	}
}

func Optimize(images []image.Image, kb int64, w, h int) {
	maxLoss := 400
	lossStep := 10
	var resizeArg *string
	switch {
	case w > 0 && h == 0:
		arg := fmt.Sprintf("--resize-width=%d", w)
		resizeArg = &arg
	case w == 0 && h > 0:
		arg := fmt.Sprintf("--resize-height=%d", h)
		resizeArg = &arg
	case w > 0 && h > 0:
		arg := fmt.Sprintf("--resize-fit=%dx%d", w, h)
		resizeArg = &arg
	}
	inputBuf := &bytes.Buffer{}
	err := Encode(inputBuf, images)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	outputBuf := &bytes.Buffer{}
	r := bytes.NewReader(inputBuf.Bytes())
	os.Stderr.WriteString("\n")
	gifsicleParams := func(colorsArg *string) (params [][]string) {
		for loss := -lossStep; loss <= maxLoss; loss += lossStep {
			var args []string
			if resizeArg != nil {
				args = append(args, *resizeArg)
			}
			if colorsArg != nil {
				args = append(args, *colorsArg)
			}
			if loss >= 0 {
				args = append(args, fmt.Sprintf("--lossy=%d", loss))
			}
			params = append(params, args)
		}
		return
	}
	tryOptimize := func(colorsArg *string) bool {
		params := gifsicleParams(colorsArg)
		return len(params) != sort.Search(len(params), func(i int) bool {
			args := params[i]
			cmd := exec.Command("gifsicle", args...)
			in, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}
			outputBuf.Reset()
			r.Seek(0, io.SeekStart)
			out, err := cmd.StdoutPipe()
			if err != nil {
				log.Fatal(err)
			}
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(in, r)
			if err != nil {
				log.Fatal(err)
			}
			err = in.Close()
			if err != nil {
				log.Fatal(err)
			}
			n, err := io.Copy(outputBuf, out)
			if err != nil {
				log.Fatal(err)
			}
			err = cmd.Wait()
			if err != nil {
				log.Fatal(err)
			}
			sizeKB := n / 1024
			log.Printf("%v: %dKB", cmd.Args, sizeKB)
			if sizeKB <= kb {
				encoded = make([]byte, n)
				copy(encoded, outputBuf.Bytes())
				images = images[:0]
				return true
			}
			return false
		})
	}
	if tryOptimize(nil) {
		return
	}
	colorsArg := "--colors=256"
	if tryOptimize(&colorsArg) {
		return
	}
	colorsArg = "--colors=128"
	if tryOptimize(&colorsArg) {
		return
	}
	colorsArg = "--colors=64"
	if tryOptimize(&colorsArg) {
		return
	}
	log.Printf("could not get size below %dKB", kb)
}
