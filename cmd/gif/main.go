package main

import (
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

	"github.com/sgreben/yeetgif/pkg/gifcmd"

	"github.com/sgreben/yeetgif/pkg/gifmeta"

	_ "image/jpeg"

	_ "golang.org/x/image/bmp"

	cli "github.com/jawher/mow.cli"
)

type configuration struct {
	Duplicate         int
	Parallelism       int
	Quiet             bool
	DelayMilliseconds func(float64) float64
	Pad               bool
	WriteMeta         bool
	NoOutput          bool
	CliOptions        string
	Raw               bool
}

var config configuration
var version string
var noQuotesRegex = regexp.MustCompile(`^[^ ()\[\]/]+$`)

const appName = "gif"

const (
	commandCat      = "cat"
	commandChop     = "chop"
	commandCompose  = "compose"
	commandCrop     = "crop"
	commandCrowd    = "crowd"
	commandEmoji    = "emoji"
	commandErase    = "erase"
	commandFried    = "fried"
	commandHue      = "hue"
	commandMeta     = "meta"
	commandNoise    = "noise"
	commandNPC      = "npc"
	commandOptimize = "optimize"
	commandPulse    = "pulse"
	commandRain     = "rain"
	commandResize   = "resize"
	commandRoll     = "roll"
	commandScan     = "scan"
	commandShake    = "shake"
	commandText     = "text"
	commandTint     = "tint"
	commandWobble   = "wobble"
	commandWoke     = "woke"
	commandZoom     = "zoom"
)

var app = cli.App(appName, fmt.Sprintf("%v", version))
var images []image.Image
var meta []gifmeta.Extension
var encoded []byte

// Global flags
var (
	duplicate = app.IntOpt("n", 30, "Duplicate a single input image this many times")
	quiet     = app.BoolOpt("q quiet", false, "Disable all log output (stderr)")
	delay     = gifcmd.FloatsCSV{Values: []float64{25}}
	pad       = app.BoolOpt("p pad", true, "Pad images")
	writeMeta = app.BoolOpt("write-meta", true, "Write command line options into output GIF metadata")
	raw       = app.BoolOpt("r raw", false, "Raw (lossless, *not* GIF) image output, for re-piping to yeetgif")
)

func main() {
	app.VarOpt("d delay-ms", &delay, "Frame delay in milliseconds")
	app.Before = func() {
		config.Raw = *raw
		config.Duplicate = *duplicate
		config.Quiet = *quiet
		config.Pad = *pad
		config.DelayMilliseconds = delay.PiecewiseLinear(0, 1)
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
	app.Command(commandRain, "。°。°( ͡° ͜ʖ ͡ °)°。°。°", CommandRain)
	app.Command(commandScan, "( ͡ ⿳ ͜ʖ ͡ ⿳ )", CommandScan)
	app.Command(commandNoise, "·͙*̩̩͙˚̩̥̩̥( ͡▓▒ ͜ʖ ͡█░ )*̩̩͙:͙", CommandNoise)
	app.Command(commandCat, "/ᐠ｡ꞈ｡ᐟ\\", CommandCat)
	app.Command(commandMeta, "(🧠 ͡ಠ ʖ̯ ͡ಠ)┌", CommandMeta)
}

func Output(w io.WriteCloser, images []image.Image, encoded []byte) {
	if len(encoded) > 0 {
		_, err := w.Write(encoded)
		if err != nil {
			log.Fatalf("write: %v", err)
		}
		err = w.Close()
		if err != nil {
			log.Fatalf("close output: %v", err)
		}
		return
	}
	if config.Pad {
		Pad(images)
	}
	if config.Raw {
		err := EncodeRaw(w, images)
		if err != nil {
			log.Fatalf("encode (raw): %v", err)
		}
		err = w.Close()
		if err != nil {
			log.Fatalf("close output: %v", err)
		}
		return
	}
	err := Encode(w, images)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	err = w.Close()
	if err != nil {
		log.Fatalf("close output: %v", err)
	}
	os.Stderr.WriteString("\n")
}
