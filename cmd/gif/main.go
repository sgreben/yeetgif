package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"time"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	"image/gif"
	_ "image/jpeg"

	cli "github.com/jawher/mow.cli"
)

type configuration struct {
	Duplicate         int
	Parallelism       int
	Quiet             bool
	DelayMilliseconds int
	Pad               bool
	WriteMeta         bool
	NoOutput          bool
	CliOptions        string
}

var config configuration
var version string
var noQuotesRegex = regexp.MustCompile(`^[^ ()\[\]/]+$`)

const appName = "gif"

const (
	commandRoll     = "roll"
	commandWobble   = "wobble"
	commandPulse    = "pulse"
	commandZoom     = "zoom"
	commandShake    = "shake"
	commandWoke     = "woke"
	commandFried    = "fried"
	commandResize   = "resize"
	commandHue      = "hue"
	commandTint     = "tint"
	commandOptimize = "optimize"
	commandCrop     = "crop"
	commandCompose  = "compose"
	commandCrowd    = "crowd"
	commandChop     = "chop"
	commandText     = "text"
	commandMeta     = "meta"
	commandEmoji    = "emoji"
	commandNPC      = "npc"
	commandErase    = "erase"
	commandNop      = "nop"
)

var app = cli.App(appName, fmt.Sprintf("%v", version))
var images []image.Image
var meta []gifmeta.Extension
var encoded []byte

// Global flags
var (
	duplicate = app.IntOpt("n", 20, "Duplicate a single input image this many times")
	quiet     = app.BoolOpt("q quiet", false, "Disable all log output (stderr)")
	delay     = app.IntOpt("d delay-ms", 20, "Frame delay in milliseconds")
	pad       = app.BoolOpt("p pad", true, "Pad images")
	writeMeta = app.BoolOpt("write-meta", true, "Write command line options into output GIF metadata")
)

func main() {
	app.Before = func() {
		config.Duplicate = *duplicate
		config.Quiet = *quiet
		config.Pad = *pad
		config.DelayMilliseconds = *delay
		config.WriteMeta = *writeMeta
		if config.Quiet {
			log.SetOutput(ioutil.Discard)
		}
	}
	app.Run(os.Args)
	if !config.NoOutput {
		Output(os.Stdout, images, encoded)
	}
}

func init() {
	rand.Seed(time.Now().Unix())
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	config.CliOptions = fmt.Sprintf("%v ", os.Args[1:])
	log.SetPrefix(config.CliOptions)
	config.Parallelism = runtime.NumCPU()
	app.Command(commandRoll, "(☭ ͜ʖ ☭)", CommandRoll)
	app.Command(commandWobble, "🍆( ͡° ͜ʖ ͡°)🍆", CommandWobble)
	app.Command(commandPulse, "( ͡◉ ͜ʖ ͡◉)", CommandPulse)
	app.Command(commandZoom, "(⌐▀͡ ̯ʖ▀)", CommandZoom)
	app.Command(commandShake, "˵(˵ ͡⚆ ͜ʖ ͡⚆˵)˵", CommandShake)
	app.Command(commandWoke, "💯  W O K E F L A R E S ( ͡ 🅱️ ͜ʖ ͡ 🅱️ ) 💯", CommandWoke)
	app.Command(commandFried, "fr͍͈i̗̟̲̻e͕̗d̬ m̷͔͊e̶̪̿m̷̙̈́é̵̤s̷̺͒", CommandFried)
	app.Command(commandHue, "( ͡☆ ͜ʖ ͡☆)", CommandHue)
	app.Command(commandTint, "🎨༼ຈل͜ຈ༽", CommandTint)
	app.Command(commandResize, "(° ͜ʖ°)¯\\_( ͡☉ ͜ʖ ͡☉)_/¯", CommandResize)
	app.Command(commandCrop, "┬┴┬┴┤ ͜ʖ ͡°)", CommandCrop)
	app.Command(commandOptimize, "👌( ͡ᵔ ͜ʖ ͡ᵔ )👌", CommandOptimize)
	app.Command(commandCompose, "(ﾉ ͡° ͜ʖ ͡°)ﾉ*:･ﾟ✧", CommandCompose)
	app.Command(commandCrowd, "(⟃ ͜ʖ ⟄) ͜ʖ ͡°)( ° ͜ʖ( ° ͜ʖ °)", CommandCrowd)
	app.Command(commandErase, "( ͡° ͜ʖ ͡°)=ε/̵͇̿̿/'̿̿ ̿ ̿ ̿ ̿ ̿", CommandErase)
	app.Command(commandChop, "✂️( ͡°Ĺ̯ ͡° )🔪", CommandChop)
	app.Command(commandText, "🅰️乁(˵ ͡☉ ͜ʖ ͡☉˵)┌🅱️", CommandText)
	app.Command(commandEmoji, "╰( ͡° ͜ʖ ͡° )つ──☆*🤔", CommandEmoji)
	app.Command(commandNPC, "•L•", CommandNPC)
	app.Command(commandNop, "乁(ᴗ ͜ʖ ᴗ)ㄏ", func(cmd *cli.Cmd) { cmd.Action = func() {} })
	app.Command(commandMeta, "(🧠 ͡ಠ ʖ̯ ͡ಠ)┌", CommandMeta)
}

func Input(r io.Reader) []image.Image {
	images := Decode(os.Stdin)
	if len(images) == 0 {
		log.Fatal("no images read")
	}
	return images
}

// Decode images from `r`
func Decode(r io.Reader) []image.Image {
	var images []image.Image
	input := &bytes.Buffer{}
	_, err := io.Copy(input, r)
	if err != nil {
		log.Fatalf("read: %v", err)
	}
	seekableReader := bytes.NewReader(input.Bytes())
	peekBuf := &bytes.Buffer{}
	tee := io.TeeReader(seekableReader, peekBuf)
	for seekableReader.Len() > 0 {
		peekBuf.Reset()
		gif, err := gif.DecodeAll(tee)
		n := int64(peekBuf.Len())
		if err == nil {
			for _, img := range gif.Image {
				images = append(images, img)
			}
			moreMeta, err := gifmeta.Read(peekBuf, func(e *gifmeta.Extension) bool {
				return e.Type == gifmeta.Comment
			})
			meta = append(meta, moreMeta...)
			if err != nil {
				log.Printf("read gif meta: %v", err)
			}
			continue
		}
		seekableReader.Seek(-n, io.SeekCurrent)
		img, _, err := image.Decode(seekableReader)
		if err != nil {
			continue
		}
		images = append(images, img)
	}
	return images
}

func Output(w io.WriteCloser, images []image.Image, encoded []byte) {
	if len(encoded) > 0 {
		_, err := w.Write(encoded)
		if err != nil {
			log.Fatalf("write: %v", err)
		}
		err = w.Close()
		if err != nil {
			log.Fatalf("close stdout: %v", err)
		}
		return
	}
	if config.Pad {
		Pad(images)
	}
	err := Encode(w, images)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	err = w.Close()
	if err != nil {
		log.Fatalf("close stdout: %v", err)
	}
	os.Stderr.WriteString("\n")
}
