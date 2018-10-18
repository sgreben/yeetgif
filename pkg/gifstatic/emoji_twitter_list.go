package gifstatic

import (
	"bytes"
	"fmt"
	"image"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/sgreben/yeetgif/pkg/rendersvg/parse"
	"github.com/sgreben/yeetgif/pkg/rendersvg/render"

	"golang.org/x/text/unicode/runenames"
)

type Emoji struct {
	KeyName            string
	Runes              []rune
	UnicodeNames       []string
	UnicodeNamesJoined string
	Keywords           []string
	ShortName          string
	Category           string
	ImageMap           map[string][]byte
}

func (e Emoji) Names() string {
	if e.ShortName == "" {
		return fmt.Sprintf("[%s] [%s]", e.UnicodeNamesJoined, e.KeyName)
	}
	return fmt.Sprintf("[%s] [%s] [%s]", e.ShortName, e.UnicodeNamesJoined, e.KeyName)
}

func (e Emoji) Image(size int) image.Image {
	img, _ := parse.Image(bytes.NewReader(e.ImageMap[e.KeyName]))
	out := image.NewRGBA(image.Rect(0, 0, size, size))
	render.Image(*img, out, nil)
	return out
}

var EmojiList []Emoji

func decodeNamePart(namePart string) []uint16 {
	code64, _ := strconv.ParseInt(namePart, 16, 64)
	code := uint32(code64 & 0xFFFFFFFF)
	if code < 0x10000 {
		return []uint16{uint16(code & 0xFFFF)}
	}
	code -= 0x10000
	code1 := 0xD800 + (code >> 10)
	code2 := 0xDC00 + (code & 0x3FF)
	return []uint16{uint16(code1 & 0xFFFF), uint16(code2 & 0xFFFF)}
}

func init() {
	add := func(name string, imageMap map[string][]byte) {
		var e Emoji
		e.KeyName = name
		e.ImageMap = imageMap
		nameParts := strings.Split(name, "-")
		var utf16Bytes []uint16
		for _, part := range nameParts {
			utf16Bytes = append(utf16Bytes, decodeNamePart(part)...)
		}
		e.Runes = utf16.Decode(utf16Bytes)
		for _, r := range e.Runes {
			e.UnicodeNames = append(e.UnicodeNames, strings.ToLower(runenames.Name(r)))
		}
		e.UnicodeNamesJoined = strings.Join(e.UnicodeNames, " ")
		if meta, ok := emojiMetaForRunes[string(e.Runes)]; ok {
			e.Keywords = meta.Keywords
			e.ShortName = strings.Replace(meta.Category+" "+meta.Name, "_", " ", -1)
			e.Category = meta.Category
		}
		EmojiList = append(EmojiList, e)
	}
	for name := range EmojiSVG {
		add(name, EmojiSVG)
	}
}
