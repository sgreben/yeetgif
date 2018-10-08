package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"time"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	_ "image/jpeg"

	"github.com/schollz/progressbar"

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
}

var config configuration
var version string
var cliOptions string
var noQuotesRegex = regexp.MustCompile(`^[^ ()\[\]/]+$`)

const (
	appName         = "gif"
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
	commandMeta     = "meta"
	commandErase    = "erase"
	commandNop      = "nop"
)

var commands = []string{
	commandRoll,
	commandWobble,
	commandPulse,
	commandZoom,
	commandShake,
	commandWoke,
	commandFried,
	commandResize,
	commandHue,
	commandTint,
	commandOptimize,
	commandCrop,
	commandMeta,
}

var app = cli.App(appName, fmt.Sprintf("%v", version))
var images []image.Image
var meta []gifmeta.Extension
var encoded []byte

func main() {
	defer os.Stderr.WriteString("\n")
	rand.Seed(time.Now().Unix())
	app.Run(os.Args)
	if config.NoOutput {
		return
	}
	if len(encoded) > 0 {
		_, err := os.Stdout.Write(encoded)
		if err != nil {
			log.Fatalf("write: %v", err)
		}
		err = os.Stdout.Close()
		if err != nil {
			log.Fatalf("close stdout: %v", err)
		}
		return
	}
	if config.Pad {
		Pad(images)
	}
	err := Encode(os.Stdout, images)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	err = os.Stdout.Close()
	if err != nil {
		log.Fatalf("close stdout: %v", err)
	}
}

func init() {
	cliOptions = fmt.Sprintf("%v ", os.Args[1:])
	log.SetPrefix(cliOptions)
	log.SetOutput(os.Stderr)
	app = cli.App(appName, fmt.Sprintf("%v", version))

	// Global flags
	var (
		duplicate = app.IntOpt("n", 20, "Duplicate a single input image this many times")
		quiet     = app.BoolOpt("q quiet", false, "Disable all log output (stderr)")
		delay     = app.IntOpt("d delay-ms", 20, "Frame delay in milliseconds")
		pad       = app.BoolOpt("p pad", true, "Pad images")
		writeMeta = app.BoolOpt("write-meta", true, "Write command line options into output GIF metadata")
	)

	app.Before = func() {
		config.Duplicate = *duplicate
		config.Parallelism = runtime.NumCPU()
		config.Quiet = *quiet
		config.Pad = *pad
		config.DelayMilliseconds = *delay
		config.WriteMeta = *writeMeta
		if config.Quiet {
			log.SetOutput(ioutil.Discard)
		}
		images = Decode(os.Stdin)
		if len(images) == 0 {
			log.Fatal("no images read")
		}
		if len(images) == 1 {
			if config.Duplicate > 0 {
				config.Duplicate--
			}
			images = Duplicate(config.Duplicate, images)
		}
	}

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
	app.Command(commandNop, "乁(ᴗ ͜ʖ ᴗ)ㄏ", func(cmd *cli.Cmd) { cmd.Action = func() {} })
	app.Command(commandMeta, "(🧠 ͡ಠ ʖ̯ ͡ಠ)┌", CommandMeta)
}

// Duplicate the `images` n times
func Duplicate(n int, images []image.Image) (out []image.Image) {
	for i := 0; i < n+1; i++ {
		out = append(out, images...)
	}
	return out
}

func newProgressBar(n int, desc string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(n, progressbar.OptionSetWriter(os.Stderr), progressbar.OptionSetDescription(desc))
	return bar
}
