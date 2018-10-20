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
		alpha      = gifcmd.Float{Value: 1.0}
		queryParts = cmd.StringsArg("EMOJI", nil, "one or more glob expressions")
		exact      = cmd.BoolOpt("e exact", false, "match the query exactly")
		pipe       = cmd.BoolOpt("p pipe", false, "overlay the emoji over input images (instead of just creating one)")
		list       = cmd.BoolOpt("l list-only", false, "just list matches")
	)
	cmd.VarOpt("s size", &size, "")
	cmd.VarOpt("a pipe-alpha", &alpha, "")
	cmd.Action = func() {
		if *pipe {
			ProcessInput()
		}
		queryBuffer := &bytes.Buffer{}
		if !*exact {
			queryBuffer.WriteRune('*')
		}
		for i, part := range *queryParts {
			queryBuffer.WriteString(part)
			if i < len(*queryParts)-1 {
				queryBuffer.WriteRune('*')
			}
		}
		if !*exact {
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
		Emoji(emoji, size.Value, alpha.Value)
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

func Emoji(emoji gifstatic.Emoji, size, alpha float64) {
	emojiImage := imaging.Resize(emoji.Image(), int(size), int(size), imaging.Lanczos)
	if len(images) == 0 {
		images = append(images, emojiImage)
		return
	}
	write := func(i int) {
		images[i] = imaging.OverlayCenter(images[0], emojiImage, alpha)
	}
	parallel(len(images), write)
}
