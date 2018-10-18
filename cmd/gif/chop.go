package main

import (
	"image"
	"math/rand"

	"github.com/sgreben/yeetgif/pkg/gifcmd"

	cli "github.com/jawher/mow.cli"
)

func CommandChop(cmd *cli.Cmd) {
	cmd.Before = Input
	cmd.Command("shuffle", "", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] [NUM_CHUNKS] [OPTIONS]"
		var (
			nChunks = cmd.IntArg("NUM_CHUNKS", 5, "shuffle chunks of frames instead of single frames")
			reverse = gifcmd.Float{Value: 0.0}
		)
		cmd.VarOpt("r reverse", &reverse, "probability of reversing a chunk")
		cmd.Action = func() {
			switch {
			case *nChunks < 0:
				*nChunks = 0
			case *nChunks >= len(images):
				*nChunks = 0
			}
			if *nChunks == 0 {
				rand.Shuffle(len(images), func(i, j int) {
					images[i], images[j] = images[j], images[i]
				})
				return
			}
			chunks := make([][]image.Image, *nChunks)
			chunkLength := len(images) / len(chunks)
			for i := range images {
				iChunk := i / chunkLength
				if iChunk >= len(chunks) {
					iChunk = len(chunks) - 1
				}
				chunks[iChunk] = append(chunks[iChunk], images[i])
			}
			rand.Shuffle(len(chunks), func(i, j int) {
				chunks[i], chunks[j] = chunks[j], chunks[i]
			})
			images = nil
			for _, chunk := range chunks {
				if rand.Float64() <= reverse.Value {
					for i := 0; i <= len(chunk)/2; i++ {
						j := len(chunk) - 1 - i
						chunk[i], chunk[j] = chunk[j], chunk[i]
					}
				}
				images = append(images, chunk...)
			}
		}
	})
	cmd.Command("dup duplicate", "", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] [START_FRAME [END_FRAME]] [OPTIONS]"
		i := cmd.IntArg("START_FRAME", 0, "")
		j := cmd.IntArg("END_FRAME", -1, "")
		rev := cmd.BoolOpt("r reverse", false, "reverse the order of the duplicated frames")
		n := cmd.IntOpt("n times", 1, "how often to duplicate the frames")
		cmd.Action = func() {
			min, max := *i, *j
			if max < 0 || max >= len(images) {
				max = len(images) - 1
			}
			if min > max {
				min = max
			}
			if min < 0 {
				min = 0
			}
			frames := make([]image.Image, max+1-min)
			for i := min; i <= max; i++ {
				frames[i-min] = images[i]
			}
			if *rev {
				for i := len(frames)/2 - 1; i >= 0; i-- {
					j := len(frames) - 1 - i
					frames[i], frames[j] = frames[j], frames[i]
				}
			}
			dups := make([]image.Image, 0, len(frames)**n)
			for j := 0; j < *n; j++ {
				dups = append(dups, frames...)
			}
			var tail []image.Image
			if max+1 < len(images) {
				tail = images[max+1:]
			}
			imagesNew := images[:max]
			imagesNew = append(imagesNew, dups...)
			imagesNew = append(imagesNew, tail...)
			images = imagesNew
		}
	})
	cmd.Command("drop", "", func(cmd *cli.Cmd) {
		cmd.Command("every", "", func(cmd *cli.Cmd) {
			n := cmd.IntArg("NTH", 2, "")
			cmd.Action = func() {
				var decimated []image.Image
				for i := range images {
					if i%*n == 0 {
						continue
					}
					decimated = append(decimated, images[i])
				}
				images = decimated
			}
		})
		cmd.Command("first", "", func(cmd *cli.Cmd) {
			n := cmd.IntArg("N", len(images)/2, "default: n/2")
			cmd.Action = func() {
				if *n >= len(images) {
					*n = len(images) - 1
				}
				images = images[*n:]
			}
		})
		cmd.Command("last", "", func(cmd *cli.Cmd) {
			n := cmd.IntArg("N", len(images)/2, "default: n/2")
			cmd.Action = func() {
				if *n > len(images) {
					*n = len(images)
				}
				images = images[:len(images)-*n]
			}
		})
	})
	cmd.Command("rev reverse", "", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			for i := 0; i <= len(images)/2; i++ {
				j := len(images) - 1 - i
				images[i], images[j] = images[j], images[i]
			}
		}
	})
}
