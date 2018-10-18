package main

import (
	"log"
	"os"

	cli "github.com/jawher/mow.cli"
)

func CommandCat(cmd *cli.Cmd) {
	cmd.Spec = "[INPUT...]"
	inputs := cmd.StringsArg("INPUT", nil, "<filename>")
	cmd.Action = func() {
		if len(*inputs) == 0 {
			InputAndDuplicate()
		}
		for _, input := range *inputs {
			f, err := os.Open(input)
			if err != nil {
				log.Fatal(err)
			}
			images = append(images, Decode(f)...)
			f.Close()
		}
	}
}
