package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sgreben/yeetgif/pkg/gifmeta"
	"github.com/sgreben/yeetgif/pkg/gifstatic"

	"github.com/ericpauley/go-quantize/quantize"

	"image/gif"

	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"

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

type metaEntry struct {
	AppName   string   `json:"appName"`
	Timestamp string   `json:"timestamp"`
	Args      []string `json:"args"`
	Version   string   `json:"version"`
}

var config configuration
var version string
var cliOptions string
var noQuotesRegex = regexp.MustCompile("\\S")

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

var app *cli.Cli
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

	var ( // Global flags
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

	app.Command(commandRoll, "(☭ ͜ʖ ☭)", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			r = gifcmd.Float{Value: 1.0}
			s = gifcmd.Float{Value: 1.0}
		)
		cmd.VarOpt("r revolutions", &r, "")
		cmd.VarOpt("s scale", &s, "")
		cmd.Action = func() {
			Roll(images, r.Value, s.Value)
		}
	})

	app.Command(commandWobble, "🍆( ͡° ͜ʖ ͡°)🍆", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		const (
			wobbleTypeSine   = "sine"
			wobbleTypeSnap   = "snap"
			wobbleTypeSmooth = "smooth"
		)
		var (
			f = gifcmd.Float{Value: 1.0}
			a = gifcmd.Float{Value: 20.0}
			t = gifcmd.Enum{
				Choices: []string{
					wobbleTypeSine,
					wobbleTypeSnap,
				},
				Value: wobbleTypeSine,
			}
		)
		cmd.VarOpt("f frequency", &f, "")
		cmd.VarOpt("a amplitude", &a, "")
		cmd.VarOpt("t type", &t, "")
		cmd.Action = func() {
			frequency := f.Value
			amplitude := a.Value
			n := len(images)
			fs := map[string]func(int) float64{
				wobbleTypeSine: func(i int) float64 {
					return amplitude * math.Sin(2*math.Pi*frequency*float64(i)/float64(n))
				},
				wobbleTypeSnap: func(i int) float64 {
					t := float64(i) / float64(n)
					y := math.Sin(2 * math.Pi * frequency * t)
					y = math.Sin(y)
					return amplitude * y
				},
			}
			Wobble(images, fs[t.Value])
		}
	})

	app.Command(commandPulse, "( ͡◉ ͜ʖ ͡◉)", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			from = gifcmd.Float{Value: 1.0}
			f    = gifcmd.Float{Value: 1.0}
			to   = gifcmd.Float{Value: 1.5}
		)
		cmd.VarOpt("0 from", &from, "")
		cmd.VarOpt("1 to", &to, "")
		cmd.VarOpt("f frequency", &f, "")
		cmd.Action = func() {
			Pulse(images, f.Value, from.Value, to.Value)
		}
	})

	app.Command(commandZoom, "(⌐▀͡ ̯ʖ▀)", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			from = gifcmd.Float{Value: 1.0}
			to   = gifcmd.Float{Value: 1.5}
		)
		cmd.VarOpt("0 from", &from, "")
		cmd.VarOpt("1 to", &to, "")
		cmd.Action = func() {
			Zoom(images, from.Value, to.Value)
		}
	})

	app.Command(commandShake, "˵(˵ ͡⚆ ͜ʖ ͡⚆˵)˵", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			random = gifcmd.Float{Value: 0.5}
			f      = gifcmd.Float{Value: 1.0}
			a      = gifcmd.Float{Value: 8.0}
		)
		cmd.VarOpt("f frequency", &f, "")
		cmd.VarOpt("a amplitude", &a, "")
		cmd.VarOpt("r random", &random, "🌀")
		cmd.Action = func() {
			Shake(images, random.Value, f.Value, a.Value)
		}
	})

	app.Command(commandWoke, "💯  W O K E F L A R E S ( ͡ 🅱️ ͜ʖ ͡ 🅱️ ) 💯", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] POINTS"
		const (
			typeFull   = "full"
			typeCenter = "center"
		)
		var (
			random = gifcmd.Float{Value: 0.5}
			s      = gifcmd.Float{Value: 0.9}
			h      = gifcmd.Float{Value: 0.8}
			l      = gifcmd.Float{Value: 1.0}
			a      = gifcmd.Float{Value: 0.8}
			ap     = gifcmd.Float{Value: 2.0}
			clip   = cmd.BoolOpt("c clip", true, "clip flares to image alpha")
			t      = gifcmd.Enum{
				Choices: []string{typeFull, typeCenter},
				Value:   typeFull,
			}
			flares    = [][2]int{}
			flaresVar = gifcmd.JSON{Value: &flares}
		)
		cmd.VarOpt("t type", &t, "")
		cmd.VarOpt("s scale", &s, "")
		cmd.VarOpt("u hue", &h, "")
		cmd.VarOpt("l lightness", &l, "")
		cmd.VarOpt("a alpha", &a, "")
		cmd.VarOpt("p alpha-pow", &ap, "")
		cmd.VarOpt("r random", &random, "🌀")
		cmd.VarArg("POINTS", &flaresVar, `flare locations, JSON, e.g. "[[123,456],[-100,23]]"`)
		cmd.Action = func() {
			var flarePoints []image.Point
			for _, p := range flares {
				flarePoints = append(flarePoints, image.Point{X: p[0], Y: p[1]})
			}
			changeHue := h.Text != ""
			var flare image.Image
			switch t.Value {
			case typeFull:
				flare = gifstatic.LensFlare
			case typeCenter:
				flare = gifstatic.LensFlareCenter
			}
			if changeHue {
				flare = imaging.AdjustHSLAFunc(flare, func(hue, _, _, _ *float64) {
					*hue += h.Value
				})
			}
			Woke(images, flare, ap.Value, a.Value, l.Value, s.Value, random.Value, flarePoints, *clip)
		}
	})

	app.Command(commandFried, "fr͍͈i̗̟̲̻e͕̗d̬ m̷͔͊e̶̪̿m̷̙̈́é̵̤s̷̺͒", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			a          = gifcmd.Float{Value: 0.33}
			b          = gifcmd.Float{Value: 0.2}
			c          = gifcmd.Float{Value: 0.9}
			clip       = cmd.BoolOpt("clip", true, "")
			q          = cmd.IntOpt("j jpeg", 84, "[0,100]")
			w          = cmd.IntOpt("w walk", 10, "🌀")
			t          = gifcmd.Float{Value: 0.4}
			n          = gifcmd.Float{Value: 1.0}
			n1         = gifcmd.Float{Value: 0.02}
			n2         = gifcmd.Float{Value: 0.5}
			n3         = gifcmd.Float{Value: 0.1}
			saturation = gifcmd.Float{Value: 3.0}
			contrast   = gifcmd.Float{Value: 6.0}
			iterations = cmd.IntOpt("i iterations", 1, "")
		)
		cmd.VarOpt("a", &a, "🅰️")
		cmd.VarOpt("b", &b, "🅱️")
		cmd.VarOpt("c", &c, "🆑")
		cmd.VarOpt("n noise", &n, "🌀️")
		cmd.VarOpt("noise1", &n1, "🌀️")
		cmd.VarOpt("noise2", &n2, "🌀️")
		cmd.VarOpt("noise3", &n3, "🌀")
		cmd.VarOpt("u saturation", &saturation, "")
		cmd.VarOpt("o contrast", &contrast, "")
		cmd.VarOpt("t tint", &t, "tint")
		cmd.Action = func() {
			n1.Value *= n.Value
			n2.Value *= n.Value
			n3.Value *= n.Value
			if *q > 100 {
				*q = 100
			}
			if *q < 0 {
				*q = 0
			}
			for i := 0; i < *iterations; i++ {
				Fried(images, t.Value, a.Value, b.Value, c.Value, *q, *w, saturation.Value, contrast.Value, n1.Value, n2.Value, n3.Value, *clip)
			}
		}
	})

	app.Command(commandHue, "( ͡☆ ͜ʖ ͡☆)", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			f    = gifcmd.Float{Value: 1.0}
			a    = gifcmd.Float{Value: 0.1}
			from = gifcmd.Float{Value: -1.0}
			to   = gifcmd.Float{Value: 1.0}
		)
		cmd.VarOpt("f frequency", &f, "")
		cmd.VarOpt("a amplitude", &a, "")
		cmd.Action = func() {
			HuePulse(images, f.Value, from.Value*a.Value, to.Value*a.Value)
		}
	})

	app.Command(commandTint, "🎨༼ຈل͜ຈ༽", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			f    = gifcmd.Float{Value: 1.0}
			a    = gifcmd.Float{Value: 0.95}
			from = gifcmd.Float{Value: 0.7}
			to   = gifcmd.Float{Value: 0.9}
		)
		cmd.VarOpt("f frequency", &f, "")
		cmd.VarOpt("0 from", &from, "")
		cmd.VarOpt("1 to", &to, "")
		cmd.VarOpt("i intensity", &a, "")
		cmd.Action = func() {
			TintPulse(images, f.Value, a.Value, from.Value, to.Value)
		}
	})
	app.Command(commandResize, "(° ͜ʖ°)¯\\_( ͡☉ ͜ʖ ͡☉)_/¯", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			s = gifcmd.Float{Value: 1.0}
			w = gifcmd.Float{Value: 0}
			h = gifcmd.Float{Value: 0}
		)
		cmd.VarOpt("s scale", &s, "")
		cmd.VarOpt("x width", &w, "width (pixels)")
		cmd.VarOpt("y height", &h, "height (pixels)")
		cmd.Action = func() {
			if s.Text != "" {
				ResizeScale(images, s.Value)
				return
			}
			ResizeTarget(images, w.Value, h.Value)
		}
	})

	app.Command(commandCrop, "┬┴┬┴┤ ͜ʖ ͡°)", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			t = gifcmd.Float{Value: 0.0}
		)
		cmd.VarOpt("t threshold", &t, "")
		cmd.Action = func() {
			AutoCrop(images, t.Value)
		}
	})

	app.Command(commandOptimize, "👌( ͡ᵔ ͜ʖ ͡ᵔ )👌", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS]"
		var (
			k = cmd.IntOpt("kb", 128, "target file size (KB)")
			w = cmd.IntOpt("x width", 128, "target width (pixels)")
			h = cmd.IntOpt("y height", 128, "target height (pixels)")
		)
		cmd.Action = func() {
			Optimize(images, int64(*k), *w, *h)
		}
	})

	app.Command(commandCompose, "(ﾉ ͡° ͜ʖ ͡°)ﾉ*:･ﾟ✧", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] INPUT"
		const (
			orderUnder = "under"
			orderOver  = "over"
		)
		const (
			positionCenter   = "center"
			positionLeft     = "left"
			positionRight    = "right"
			positionTop      = "top"
			positionBottom   = "bottom"
			positionAbsolute = "abs"
		)
		var (
			input = cmd.StringArg("INPUT", "", "")
			z     = gifcmd.Enum{
				Choices: []string{orderUnder, orderOver},
				Value:   orderOver,
			}
			p = gifcmd.Enum{
				Choices: []string{
					positionCenter,
					positionLeft,
					positionRight,
					positionTop,
					positionBottom,
					positionAbsolute,
				},
				Value: positionCenter,
			}
			x = cmd.IntOpt("x", 0, "")
			y = cmd.IntOpt("y", 0, "")
			s = gifcmd.Float{Value: 1.0}
		)
		cmd.VarOpt("z z-order", &z, z.Help())
		cmd.VarOpt("p position", &p, p.Help())
		cmd.VarOpt("s scale", &s, "")
		cmd.Action = func() {
			f, err := os.Open(*input)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			layer := Decode(f)
			if s.Value != 1.0 {
				ResizeScale(layer, s.Value)
			}
			offset := image.Point{*x, *y}
			var imageAnchor, layerAnchor imaging.Anchor
			switch p.Value {
			case positionAbsolute:
				imageAnchor = imaging.TopLeft
				layerAnchor = imaging.Center
			case positionCenter:
				imageAnchor = imaging.Center
				layerAnchor = imaging.Center
			case positionLeft:
				imageAnchor = imaging.Left
				layerAnchor = imaging.Right
			case positionRight:
				imageAnchor = imaging.Right
				layerAnchor = imaging.Left
			case positionTop:
				imageAnchor = imaging.Top
				layerAnchor = imaging.Bottom
			case positionBottom:
				imageAnchor = imaging.Bottom
				layerAnchor = imaging.Top
			}
			switch z.Value {
			case orderOver:
				Compose(images, layer, offset, imageAnchor, layerAnchor)
			case orderUnder:
				Compose(layer, images, offset, layerAnchor, imageAnchor)
			}
		}
	})

	app.Command(commandCrowd, "(⟃ ͜ʖ ⟄) ͜ʖ ͡°)( ° ͜ʖ( ° ͜ʖ °)", func(cmd *cli.Cmd) {
		var (
			n   = cmd.IntOpt("n", 3, "crowd size")
			rpx = gifcmd.Float{Value: 0.5}
			rpy = gifcmd.Float{Value: 0.25}
			rs  = gifcmd.Float{Value: 0.25}
			ra  = gifcmd.Float{Value: 0.0}
			ro  = gifcmd.Float{Value: 1.0}
		)
		cmd.VarOpt("x", &rpx, "random x")
		cmd.VarOpt("y", &rpy, "random y")
		cmd.VarOpt("s scale", &rs, "random scale")
		cmd.VarOpt("a alpha", &ra, "random alpha")
		cmd.VarOpt("o offset", &ro, "random frame offset")
		cmd.Action = func() {
			Crowd(images, *n, rpx.Value, rpy.Value, rs.Value, ra.Value, ro.Value)
			AutoCrop(images, 0.0)
		}
	})

	app.Command(commandNop, "乁(ᴗ ͜ʖ ᴗ)ㄏ", func(cmd *cli.Cmd) {
		cmd.Action = func() {}
	})

	app.Command(commandMeta, "(🧠 ͡ಠ ʖ̯ ͡ಠ)┌", func(cmd *cli.Cmd) {
		cmd.Command("show", "show 🧠", func(cmd *cli.Cmd) {
			raw := cmd.BoolOpt("r raw", false, "print raw JSON")
			cmd.Action = func() {
				config.NoOutput = true
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
	})
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

// Duplicate the `images` n times
func Duplicate(n int, images []image.Image) (out []image.Image) {
	for i := 0; i < n+1; i++ {
		out = append(out, images...)
	}
	return out
}

// Encode the `images` as a GIF to `w`
func Encode(w io.Writer, images []image.Image) error {
	var maxWidth, maxHeight int
	for _, img := range images {
		width := img.Bounds().Dx()
		if width > maxWidth {
			maxWidth = width
		}
		height := img.Bounds().Dy()
		if height > maxHeight {
			maxHeight = height
		}
	}

	out := gif.GIF{
		LoopCount: 0,
		Config: image.Config{
			Width:  maxWidth,
			Height: maxHeight,
		},
	}
	quantizer := quantize.MedianCutQuantizer{
		AddTransparent: true,
		Aggregation:    quantize.Mean,
	}
	out.Image = make([]*image.Paletted, len(images))
	out.Delay = make([]int, len(images))
	out.Disposal = make([]byte, len(images))
	quantize := func(i int) {
		b := images[i].Bounds()
		palette := quantizer.Quantize(make([]color.Color, 0, 256), images[i])
		imgQuantized := image.NewPaletted(b, palette)
		draw.FloydSteinberg.Draw(imgQuantized, b, images[i], image.ZP)
		out.Image[i] = imgQuantized
		out.Delay[i] = config.DelayMilliseconds / 10
		out.Disposal[i] = gif.DisposalBackground
	}
	parallel(len(images), quantize)

	if !config.WriteMeta {
		return gif.EncodeAll(w, &out)
	}
	buf := &bytes.Buffer{}
	gif.EncodeAll(buf, &out)
	metaJSONBytes, _ := json.Marshal(metaEntry{
		AppName:   appName,
		Timestamp: time.Now().Format(time.RFC3339),
		Args:      os.Args[1:],
		Version:   version,
	})
	meta = append(meta, gifmeta.Extension{
		Type:   gifmeta.Comment,
		Blocks: gifmeta.Blocks(metaJSONBytes),
	})
	return gifmeta.Append(w, buf, meta...)
}

// Roll the `images` `rev` times
func Roll(images []image.Image, rev, preScale float64) {
	n := len(images)
	rotate := func(i int) {
		angle := 360 * rev * float64(i) / float64(n)
		bPre := images[i].Bounds()
		if preScale != 1.0 {
			images[i] = imaging.Resize(images[i], int(float64(bPre.Dx())*preScale), int(float64(bPre.Dy())*preScale), imaging.Lanczos)
		}
		images[i] = imaging.Rotate(images[i], angle, color.Transparent)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		if !config.Pad {
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), rotate)
}

// Wobble `images` `frequency` times by `amplitude` degrees
func Wobble(images []image.Image, f func(int) float64) {
	rotate := func(i int) {
		angle := f(i)
		bPre := images[i].Bounds()
		images[i] = imaging.Rotate(images[i], angle, color.Transparent)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		if !config.Pad {
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), rotate)
}

// Pulse `images` `frequency` times between scales `from` and `to`
func Pulse(images []image.Image, frequency, from, to float64) {
	n := len(images)
	scale := func(i int) {
		weight := math.Sin(2 * math.Pi * frequency * float64(i) / float64(n))
		scale := from*weight + to*(1-weight)
		bPre := images[i].Bounds()
		width := float64(bPre.Dx()) * scale
		height := float64(bPre.Dy()) * scale
		images[i] = imaging.Resize(images[i], int(width), int(height), imaging.Lanczos)
		if !config.Pad {
			bPost := images[i].Bounds()
			offset := image.Point{
				X: (bPost.Dx() - bPre.Dx()) / 2,
				Y: (bPost.Dy() - bPre.Dy()) / 2,
			}
			bPre.Min = bPre.Min.Add(offset)
			bPre.Max = bPre.Max.Add(offset)
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), scale)
}

// Zoom `images` once from `from` to `to`
func Zoom(images []image.Image, from, to float64) {
	n := len(images)
	scale := func(i int) {
		weight := float64(i) / float64(n)
		scale := from*(1-weight) + to*weight
		bPre := images[i].Bounds()
		width := float64(bPre.Dx()) * scale
		height := float64(bPre.Dy()) * scale
		images[i] = imaging.Resize(images[i], int(width), int(height), imaging.Gaussian)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		images[i] = imaging.Crop(images[i], bPre)
	}
	parallel(len(images), scale)
}

// Shake `images`
func Shake(images []image.Image, random, frequency, amplitude float64) {
	n := len(images)
	phaseY := math.Pi / 2
	move := func(i int) {
		rX, rY := rand.Float64(), rand.Float64()
		offset := image.Point{
			X: int(amplitude * math.Sin(2*math.Pi*frequency*float64(i)/float64(n)+(rX*2*math.Pi))),
			Y: int(amplitude * math.Sin(2*math.Pi*frequency*float64(i)/float64(n)+phaseY+(rY*2*math.Pi))),
		}
		if !config.Pad {
			images[i] = imaging.Paste(image.NewNRGBA(images[i].Bounds()), images[i], offset)
			return
		}
		bounds := images[i].Bounds()
		bounds.Min.X -= int(amplitude)
		bounds.Max.X += int(amplitude)
		bounds.Min.Y -= int(amplitude)
		bounds.Max.Y += int(amplitude)
		images[i] = imaging.Paste(image.NewNRGBA(bounds), images[i], offset)
	}
	parallel(len(images), move)
}

// Woke flares
func Woke(images []image.Image, flare image.Image, alphaPow, alpha, lightness, scale, random float64, flares []image.Point, clip bool) {
	b := flare.Bounds()
	width := int(float64(b.Dx()) * scale)
	height := int(float64(b.Dy()) * scale)
	flare = imaging.Resize(flare, width, height, imaging.Lanczos)
	flare = imaging.AdjustHSLAFunc(flare, func(h, s, l, a *float64) {
		*a = math.Pow(*a, alphaPow) * alpha
		*l = (*l) * lightness
	})
	woke := func(i int) {
		b := images[i].Bounds()
		layer := imaging.New(b.Dx(), b.Dy(), color.Transparent)
		flip := true
		for _, p := range flares {
			flare := flare
			if random > 0 {
				flare = imaging.Rotate(flare, rand.Float64()*random*360, color.Transparent)
			}
			if flip {
				flare = imaging.FlipV(imaging.FlipH(flare))
			}
			b := flare.Bounds()
			layer = imaging.OverlayWithOp(layer, flare, image.Point{
				X: (p.X - b.Dx()/2),
				Y: (p.Y - b.Dy()/2),
			}, imaging.OpLighten)
			flip = !flip
		}
		woke := imaging.Overlay(images[i], layer, image.ZP, 1.0)
		if clip {
			images[i] = imaging.OverlayWithOp(
				woke,
				images[i],
				image.ZP,
				imaging.OpReplaceAlpha,
			)
			return
		}
		images[i] = woke
	}
	parallel(len(images), woke)
}

// Fried meme
func Fried(images []image.Image, tint, a, b, c float64, loss, step int, saturation, contrast, noise1, noise2, noise3 float64, clip bool) {
	if loss < 0 {
		loss = 0
	}
	if loss > 100 {
		loss = 100
	}
	jpeg := func(i, quality int) {
		buf := &bytes.Buffer{}
		imaging.Encode(buf, images[i], imaging.JPEG, imaging.JPEGQuality(quality))
		images[i], _, _ = image.Decode(buf)
	}
	orange := color.RGBA{
		R: 255,
		G: 30,
		B: 0,
	}
	bounds := images[0].Bounds()
	explodePoint := image.Point{
		X: int(rand.Float64() * float64(bounds.Dx())),
		Y: int(rand.Float64() * float64(bounds.Dy())),
	}
	n := len(images)
	explodePoints := make([]image.Point, n)
	for i := 0; i <= n/2; i++ {
		explodePoints[i] = explodePoint
		explodePoints[n-1-i] = explodePoint
		explodePoint.X += int(rand.Float64()*2*float64(step)) - step
		explodePoint.Y += int(rand.Float64()*2*float64(step)) - step
	}
	fry := func(i int) {
		explodePoint := explodePoints[i]
		original := images[i]
		images[i] = imaging.Explode(images[i], explodePoint, a, b, c)
		exploded := images[i]
		images[i] = imaging.AdjustTint(images[i], tint, orange)
		images[i] = imaging.AdjustNoiseHSL(images[i], noise1, noise2, noise3)
		jpeg(i, 100-loss)
		images[i] = imaging.AdjustSaturation(images[i], saturation)
		images[i] = imaging.AdjustSigmoid(images[i], 0.5, contrast)
		jpeg(i, 100-(loss/2))
		if clip {
			images[i] = imaging.OverlayWithOp(images[i], original, image.ZP, imaging.OpReplaceAlpha)
		}
		images[i] = imaging.OverlayWithOp(images[i], exploded, image.ZP, imaging.OpMinAlpha)
	}
	parallel(len(images), fry)
}

// ResizeScale resizes by a factor
func ResizeScale(images []image.Image, scale float64) {
	resize := func(i int) {
		b := images[i].Bounds()
		width, height := float64(b.Dx()), float64(b.Dy())
		images[i] = imaging.Resize(images[i], int(width*scale), int(height*scale), imaging.Lanczos)
	}
	parallel(len(images), resize)
}

// ResizeTarget resizes to fit in the given bounds
func ResizeTarget(images []image.Image, width, height float64) {
	resize := func(i int) {
		b := images[i].Bounds()
		w, h := float64(b.Dx()), float64(b.Dy())
		scale := 1.0
		rw, rh := math.Abs(width/w-1), math.Abs(height/w-1)
		switch {
		case width > 0 && height > 0:
			if rw < rh {
				scale = width / w
			} else {
				scale = height / h
			}
		case width > 0:
			scale = width / w
		case height > 0:
			scale = height / h
		}
		images[i] = imaging.Resize(images[i], int(w*scale), int(h*scale), imaging.Lanczos)
	}
	parallel(len(images), resize)
}

func TintPulse(images []image.Image, frequency, weight, from, to float64) {
	n := float64(len(images))
	dist := math.Min(to-from, from+(1-to))
	if dist == 0 && to != from {
		dist = 1.0
	}
	mid := from + (dist / 2)
	if mid < 0 {
		mid++
	}
	if mid > 1 {
		mid--
	}
	tint := func(i int) {
		hue := mid + (dist * math.Sin(2*math.Pi*frequency*float64(i)/n) / 2)
		if hue < 0 {
			hue++
		}
		if hue > 1 {
			hue--
		}
		images[i] = imaging.AdjustHue(images[i], weight, hue)
	}
	parallel(len(images), tint)
}

func HuePulse(images []image.Image, frequency, from, to float64) {
	n := float64(len(images))
	mid := (from + to) / 2
	dist := to - from
	hue := func(i int) {
		delta := mid + (dist * math.Sin(math.Pi+2*math.Pi*frequency*float64(i)/n) / 2)
		images[i] = imaging.AdjustHueRotate(images[i], delta)
	}
	parallel(len(images), hue)
}

func Optimize(images []image.Image, kb int64, w, h int) {
	maxLoss := 400
	maxColors := 256
	lossStep := 10
	colorsStep := 16
	var resizeArg *string
	if w > 0 && h > 0 {
		arg := fmt.Sprintf("--resize-fit=%dx%d", w, h)
		resizeArg = &arg
	}
	inputBuf := &bytes.Buffer{}
	err := Encode(inputBuf, images)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	outputBuf := &bytes.Buffer{}
	r := bytes.NewReader(inputBuf.Bytes())
	os.Stderr.WriteString("\n")
	for colors := maxColors + colorsStep; colors > 0; colors -= colorsStep {
		var colorsArg *string
		if colors <= maxColors {
			arg := fmt.Sprintf("--colors=%d", colors)
			colorsArg = &arg
		}
		for loss := 0; loss <= maxLoss; loss += lossStep {
			outputBuf.Reset()
			r.Seek(0, io.SeekStart)
			args := []string{fmt.Sprintf("--lossy=%d", loss)}
			if resizeArg != nil {
				args = append(args, *resizeArg)
			}
			if colorsArg != nil {
				args = append(args, *colorsArg)
			}
			cmd := exec.Command("gifsicle", args...)
			in, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}
			out, err := cmd.StdoutPipe()
			if err != nil {
				log.Fatal(err)
			}
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(in, r)
			if err != nil {
				log.Fatal(err)
			}
			err = in.Close()
			if err != nil {
				log.Fatal(err)
			}
			n, err := io.Copy(outputBuf, out)
			if err != nil {
				log.Fatal(err)
			}
			err = cmd.Wait()
			if err != nil {
				log.Fatal(err)
			}
			sizeKB := n / 1024
			log.Printf("%v: %dKB", cmd.Args, sizeKB)
			if sizeKB <= kb {
				encoded = outputBuf.Bytes()
				images = images[:0]
				return
			}
		}
	}
	log.Printf("could not get size below %dKB", kb)
}

func Pad(images []image.Image) {
	width, height := 0, 0
	for i := range images {
		if w := images[i].Bounds().Dx(); w > width {
			width = w
		}
		if h := images[i].Bounds().Dy(); h > height {
			height = h
		}
	}
	pad := func(i int) {
		padded := imaging.New(width, height, color.Transparent)
		images[i] = imaging.PasteCenter(padded, images[i])
	}
	parallel(len(images), pad)
}

func AutoCrop(images []image.Image, threshold float64) {
	width, height := 0, 0
	for i := range images {
		if w := images[i].Bounds().Dx(); w > width {
			width = w
		}
		if h := images[i].Bounds().Dy(); h > height {
			height = h
		}
	}
	sample := imaging.Clone(images[0])
	for i := range images {
		sample = imaging.OverlayWithOp(sample, images[i], image.ZP, imaging.OpMaxAlpha)
	}
	b := imaging.OpaqueBounds(sample, uint8(threshold*255))
	crop := func(i int) {
		images[i] = imaging.Crop(images[i], b)
	}
	parallel(len(images), crop)
}

func Crowd(images []image.Image, n int, rpx, rpy, rs, ra, ro float64) {
	width, height := 0, 0
	for i := range images {
		if w := images[i].Bounds().Dx(); w > width {
			width = w
		}
		if h := images[i].Bounds().Dy(); h > height {
			height = h
		}
	}
	mid := image.Point{
		X: width / 2,
		Y: height / 2,
	}
	p := make([]image.Point, n)
	s := make([]float64, n)
	a := make([]float64, n)
	o := make([]int, n)
	var b, bOriginal image.Rectangle
	bOriginal.Max.X = width
	bOriginal.Max.Y = height
	b.Max.X = bOriginal.Max.X
	b.Max.Y = bOriginal.Max.Y
	for j := range p {
		s[j] = 1.0 - (rand.Float64() * rs)
		p[j].X = int(s[j] * float64(width) * rpx * 2 * (rand.Float64() - 0.5))
		p[j].Y = int(s[j] * float64(height) * rpy * 2 * (rand.Float64() - 0.5))
		b = b.Union(bOriginal.Add(p[j]))
		o[j] = rand.Intn(int(ro * float64(len(images)-1)))
		a[j] = 1.0 - (rand.Float64() * ra)
	}
	offset := b.Min
	b = b.Sub(offset)
	originals := images
	images = make([]image.Image, len(originals))
	crowd := func(i int) {
		crowded := imaging.New(b.Dx(), b.Dy(), color.Transparent)
		for j := range p {
			iLayer := (o[j] + i) % len(images)
			layer := originals[iLayer]
			bLayer := layer.Bounds()
			w, h := float64(bLayer.Dx())*s[j], float64(bLayer.Dy())*s[j]
			layer = imaging.Resize(layer, int(w), int(h), imaging.Lanczos)
			midLayer := imaging.AnchorPoint(layer, imaging.Center)
			p := p[j].Sub(midLayer).Add(mid).Sub(offset)
			crowded = imaging.Overlay(crowded, layer, p, a[j])
		}
		images[i] = crowded
	}
	overwrite := func(i int) {
		originals[i] = images[i]
	}
	parallel(len(images), crowd)
	parallel(len(images), overwrite)
}

func Compose(a, b []image.Image, p image.Point, anchorA, anchorB imaging.Anchor) {
	compose := func(i int) {
		ai := i % len(a)
		bi := i % len(b)
		under := a[ai]
		over := b[bi]
		overOffset := imaging.AnchorPoint(under, anchorA).Sub(imaging.AnchorPoint(over, anchorB))
		bounds := under.Bounds().Union(over.Bounds().Add(overOffset))
		bg := image.NewNRGBA(bounds.Sub(bounds.Min))
		bg = imaging.Paste(bg, under, bounds.Min.Mul(-1))
		images[i] = imaging.Overlay(bg, over, overOffset.Sub(bounds.Min), 1.0)
	}
	var an, bn, z big.Int
	an.SetInt64(int64(len(a)))
	bn.SetInt64(int64(len(b)))
	n := int(z.Mul(z.Div(&bn, z.GCD(nil, nil, &an, &bn)), &an).Int64())
	images = make([]image.Image, n)
	parallel(n, compose)
}

func newProgressBar(n int, desc string) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(n, progressbar.OptionSetWriter(os.Stderr), progressbar.OptionSetDescription(desc))
	return bar
}

func parallel(n int, f func(int)) {
	work := make(chan int, 4)
	var wg sync.WaitGroup
	wg.Add(config.Parallelism)
	bar := newProgressBar(n, cliOptions)
	bar.RenderBlank()

	for i := 0; i < config.Parallelism; i++ {
		go func() {
			for i := range work {
				f(i)
				bar.Add(1)
			}
			wg.Done()
		}()
	}
	for i := 0; i < n; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
}
