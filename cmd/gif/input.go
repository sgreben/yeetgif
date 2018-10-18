package main

import (
	"image"
	"io"
	"log"
	"os"
)

func Input() {
	images = CommandInput(os.Stdin)
}

func InputAndDuplicate() {
	images = CommandInput(os.Stdin)
	CommandDuplicate(config.Duplicate)
}

func CommandInput(r io.Reader) []image.Image {
	images := Decode(r)
	if len(images) == 0 {
		log.Fatal("no images read")
	}
	return images
}
