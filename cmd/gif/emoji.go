package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/sgreben/yeetgif/pkg/imaging"

	"github.com/gobwas/glob"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/gifstatic"

	cli "github.com/jawher/mow.cli"
)

func CommandEmoji(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS] EMOJI..."
	var (
		size       = gifcmd.Float{Value: 128}
		queryParts = cmd.StringsArg("EMOJI", []string{"face", "joy"}, "")
		substring  = cmd.BoolOpt("sub", true, "match *EMOJI*")
		overlay    = cmd.BoolOpt("o overlay", false, "read images from stdin (instead of just creating one)")
		list       = cmd.BoolOpt("l list-only", false, "just list matches")
	)
	cmd.VarOpt("s font-size", &size, "")
	cmd.Action = func() {
		if *overlay {
			ProcessInput()
		}
		queryBuffer := &bytes.Buffer{}
		if *substring {
			queryBuffer.WriteRune('*')
		}
		for i, part := range *queryParts {
			queryBuffer.WriteString(part)
			if i < len(*queryParts)-1 {
				queryBuffer.WriteRune('*')
			}
		}
		if *substring {
			queryBuffer.WriteRune('*')
		}
		query := queryBuffer.String()
		queryGlob, err := glob.Compile(query)
		if err != nil {
			log.Fatalf("parse glob %q: %v", query, err)
		}
		matches := EmojiMatches(queryGlob)
		if len(matches) == 0 {
			log.Fatalf("no emoji matches found for %q", query)
		}
		if *list || len(matches) > 1 {
			log.Printf("%d matching emoji found for %q", len(matches), query)
			for _, m := range matches {
				log.Printf("%s\t%s", string(m.Runes), m.UnicodeNames)
			}
		}
		if *list {
			config.NoOutput = true
			return
		}
		emoji := matches[0]
		log.Printf("picked %s (%s)", string(emoji.Runes), emoji.UnicodeNames)
		Emoji(emoji, size.Value)
	}
}

func EmojiMatches(queryGlob glob.Glob) (matches []gifstatic.Emoji) {
	for _, e := range gifstatic.EmojiList {
		if queryGlob.Match(strings.Join(e.UnicodeNames, " ")) {
			matches = append(matches, e)
		}
	}
	return
}

func Emoji(emoji gifstatic.Emoji, size float64) {
	emojiImage := imaging.Resize(emoji.Image(), int(size), int(size), imaging.Lanczos)
	if len(images) == 0 {
		images = append(images, emojiImage)
		return
	}
	write := func(i int) {
		images[i] = imaging.OverlayCenter(images[0], emojiImage, 1.0)
	}
	parallel(len(images), write)
}
