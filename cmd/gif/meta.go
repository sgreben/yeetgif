package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	_ "image/jpeg"

	cli "github.com/jawher/mow.cli"
)

type metaEntry struct {
	AppName   string   `json:"appName"`
	Timestamp string   `json:"timestamp"`
	Args      []string `json:"args"`
	Version   string   `json:"version"`
}

func CommandMeta(cmd *cli.Cmd) {
	cmd.Command("show", "show 🧠", func(cmd *cli.Cmd) {
		raw := cmd.BoolOpt("r raw", false, "print raw JSON")
		pipe := cmd.BoolOpt("p pipe", false, "print shell pipe")
		cmd.Action = func() {
			config.NoOutput = true
			if *pipe {
				var shellCommands [][]string
				for _, e := range meta {
					if e.Type != gifmeta.Comment {
						continue
					}
					s := e.String()
					m := metaEntry{AppName: appName}
					err := json.NewDecoder(strings.NewReader(s)).Decode(&m)
					if err != nil {
						continue
					}
					var shellCommand []string
					shellCommand = append(shellCommand, appName)
					shellCommand = append(shellCommand, m.Args...)
					shellCommands = append(shellCommands, shellCommand)
				}
				n := len(shellCommands)
				for i, ss := range shellCommands {
					for _, s := range ss {
						if noQuotesRegex.MatchString(s) {
							fmt.Printf("%s ", s)
							continue
						}
						fmt.Printf("%q ", s)
					}
					if i < n - 1 {
						fmt.Print("| ")
					}
				}
				fmt.Println()
				return
			}
			for _, e := range meta {
				if e.Type == gifmeta.Comment {
					s := e.String()
					m := metaEntry{AppName: appName}
					err := json.NewDecoder(strings.NewReader(s)).Decode(&m)
					printRaw := *raw || err != nil
					if printRaw {
						fmt.Println(s)
						continue
					}
					fmt.Printf("[%s] %s ", m.Timestamp, m.AppName)
					for _, arg := range m.Args {
						if noQuotesRegex.MatchString(arg) {
							fmt.Printf("%s ", arg)
							continue
						}
						fmt.Printf("%q ", arg)
					}
					fmt.Println()
				}
			}
		}
	})
	cmd.Command("add", "add 🧠", func(cmd *cli.Cmd) {
		d := cmd.StringArg("DATA", "", "")
		cmd.Action = func() {
			meta = append(meta, gifmeta.Extension{
				Type:   gifmeta.Comment,
				Blocks: gifmeta.Blocks([]byte(*d)),
			})
		}
	})
	cmd.Command("clear", "remove 🧠", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			meta = nil
			config.WriteMeta = false
		}
	})
}
