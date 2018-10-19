package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"os/exec"

	cli "github.com/jawher/mow.cli"
)

func CommandOptimize(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS]"
	var (
		k = cmd.IntOpt("kb", 128, "target file size (KB)")
		w = cmd.IntOpt("x width", 128, "target width (pixels)")
		h = cmd.IntOpt("y height", 128, "target height (pixels)")
	)
	cmd.Action = func() {
		Optimize(images, int64(*k), *w, *h)
	}
}

func Optimize(images []image.Image, kb int64, w, h int) {
	maxLoss := 400
	maxColors := 256
	lossStep := 10
	colorsStep := 16
	var resizeArg *string
	switch {
	case w > 0 && h == 0:
		arg := fmt.Sprintf("--resize-width=%d", w)
		resizeArg = &arg
	case w == 0 && h > 0:
		arg := fmt.Sprintf("--resize-height=%d", w)
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
	for colors := maxColors + colorsStep; colors > 0; colors -= colorsStep {
		var colorsArg *string
		if colors <= maxColors {
			arg := fmt.Sprintf("--colors=%d", colors)
			colorsArg = &arg
		}
		for loss := 0; loss <= maxLoss; loss += lossStep {
			outputBuf.Reset()
			r.Seek(0, io.SeekStart)
			args := []string{fmt.Sprintf("--lossy=%d", loss)}
			if resizeArg != nil {
				args = append(args, *resizeArg)
			}
			if colorsArg != nil {
				args = append(args, *colorsArg)
			}
			cmd := exec.Command("gifsicle", args...)
			in, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}
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
				encoded = outputBuf.Bytes()
				images = images[:0]
				return
			}
		}
	}
	log.Printf("could not get size below %dKB", kb)
}
