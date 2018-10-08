package main

import "image"

func CommandDuplicate(n int) {
	if len(images) == 1 {
		if n > 0 {
			n--
		}
		images = Duplicate(n, images)
	}
}

// Duplicate the `images` n times
func Duplicate(n int, images []image.Image) (out []image.Image) {
	for i := 0; i < n+1; i++ {
		out = append(out, images...)
	}
	return out
}
