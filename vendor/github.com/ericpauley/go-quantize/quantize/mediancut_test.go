package quantize

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"testing"

	_ "image/jpeg"
)

func TestBuildBucket(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Couldn't open test file")
	}
	i, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Couldn't decode test file")
	}

	q := MedianCutQuantizer{Mode, nil, false}

	colors := q.buildBucket(i)
	t.Logf("Naive color map contains %d elements", len(colors))

	for _, p := range colors {
		if p.p == 0 {
			t.Fatal("Bucket had a 0 priority element")
		}
	}

	q = MedianCutQuantizer{Mode, func(i image.Image, x int, y int) uint32 {
		if x < 2 || y < 2 || x > i.Bounds().Max.X-2 || y > i.Bounds().Max.X-2 {
			return 1
		}
		return 0
	}, false}

	colors = q.buildBucket(i)
	t.Logf("Color map contains %d elements", len(colors))
}

func ExampleMedianCutQuantizer() {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		fmt.Println("Couldn't open test file")
		return
	}
	i, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't decode test file")
		return
	}
	q := MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	fmt.Println(p)
}

func TestQuantize(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Couldn't open test file")
	}
	i, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Couldn't decode test file")
	}
	q := MedianCutQuantizer{Mean, nil, false}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	t.Logf("Created palette with %d colors", len(p))

	q = MedianCutQuantizer{Mean, nil, false}
	p2 := q.Quantize(make([]color.Color, 0, 256), i)

	if len(p) != len(p2) {
		t.Fatal("Quantize is not deterministic")
	}

	for i := range p {
		if p[i] != p2[i] {
			t.Fatal("Quantize is not deterministic")
		}
	}

	q = MedianCutQuantizer{Mode, nil, false}
	p = q.Quantize(make([]color.Color, 0, 256), i)
	t.Logf("Created palette with %d colors", len(p))

	q = MedianCutQuantizer{Mean, nil, true}
	p = q.Quantize(color.Palette{color.RGBA{0, 0, 0, 0}}, i)
	t.Logf("Created palette with %d colors", len(p))

	q = MedianCutQuantizer{Mean, nil, true}
	p = q.Quantize(make([]color.Color, 0, 256), i)
	t.Logf("Created palette with %d colors", len(p))
}

func BenchmarkQuantize(b *testing.B) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		b.Fatal("Couldn't open test file")
	}
	m, _, err := image.Decode(file)
	if err != nil {
		b.Fatal("Couldn't decode test file")
	}
	q := MedianCutQuantizer{Mean, nil, false}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Quantize(make([]color.Color, 0, 256), m)
	}
}

func TestRGBAQuantize(t *testing.T) {
	i := image.NewRGBA(image.Rect(0, 0, 1, 1))
	q := MedianCutQuantizer{Mean, nil, false}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	t.Logf("Created palette with %d colors", len(p))
}

// TestOverQuantize ensures that the quantizer can properly handle an image with more space than needed in the palette
func TestOverQuantize(t *testing.T) {
	file, err := os.Open("test_image2.gif")
	if err != nil {
		t.Fatal("Couldn't open test file")
	}
	i, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Couldn't decode test file")
	}
	q := MedianCutQuantizer{Mean, nil, false}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	t.Logf("Created palette with %d colors", len(p))
}

func TestEmptyQuantize(t *testing.T) {
	i := image.NewNRGBA(image.Rect(0, 0, 0, 0))

	q := MedianCutQuantizer{Mean, nil, false}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	if len(p) != 0 {
		t.Fatal("Quantizer returned colors for empty image")
	}
	t.Logf("Created palette with %d colors", len(p))
}

func TestGif(t *testing.T) {
	file, err := os.Open("test_image.jpg")
	if err != nil {
		t.Fatal("Couldn't open test file")
	}
	i, _, err := image.Decode(file)
	if err != nil {
		t.Fatal("Couldn't decode test file")
	}

	q := MedianCutQuantizer{Mode, nil, false}
	f, err := os.Create("test_output.gif")
	if err != nil {
		t.Fatal("Couldn't open output file")
	}

	options := gif.Options{NumColors: 128, Quantizer: q, Drawer: nil}

	w := bufio.NewWriter(f)

	gif.Encode(w, i, &options)
}
