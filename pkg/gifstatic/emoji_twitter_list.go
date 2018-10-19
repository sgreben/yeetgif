package gifstatic

import (
	"bytes"
	"fmt"
	"image"
	"strings"

	"golang.org/x/text/unicode/runenames"
	// PNG support
	_ "image/png"
)

type Emoji struct {
	KeyName      string
	Runes        []rune
	UnicodeNames []string
}

func (e Emoji) Image() image.Image {
	img, _, _ := image.Decode(bytes.NewReader(EmojiPNG[e.KeyName]))
	return img
}

var EmojiList []Emoji

func init() {
	for name := range EmojiPNG {
		var e Emoji
		e.KeyName = name
		nameParts := strings.Split(name, "-")
		for _, part := range nameParts {
			var r rune
			_, err := fmt.Sscanf(part, "%x", &r)
			if err != nil {
				panic(err)
			}
			e.Runes = append(e.Runes, r)
			e.UnicodeNames = append(e.UnicodeNames, strings.ToLower(runenames.Name(r)))
		}
		EmojiList = append(EmojiList, e)
	}
}
