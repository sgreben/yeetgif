package bench

import (
	"image"
	"image/color"
	"os"
	"testing"

	_ "image/jpeg"

	"github.com/ericpauley/go-quantize/quantize"
	"github.com/esimov/colorquant"
	"github.com/soniakeys/quant/mean"
	"github.com/soniakeys/quant/median"
)

func getImage(b *testing.B) (m image.Image) {
	file, err := os.Open("../test_image.jpg")
	if err != nil {
		b.Fatal("Couldn't open test file")
	}
	m, _, err = image.Decode(file)
	if err != nil {
		b.Fatal("Couldn't decode test file")
	}
	b.ReportAllocs()
	b.ResetTimer()
	return
}

func BenchmarkQuantize(b *testing.B) {
	m := getImage(b)
	q := quantize.MedianCutQuantizer{quantize.Mean, nil, false}
	for i := 0; i < b.N; i++ {
		q.Quantize(make([]color.Color, 0, 256), m)
	}
}

func BenchmarkSoniakeysMedian(b *testing.B) {
	m := getImage(b)
	q := median.Quantizer(256)
	for i := 0; i < b.N; i++ {
		q.Palette(m)
	}
}

func BenchmarkSoniakeysMean(b *testing.B) {
	m := getImage(b)
	q := mean.Quantizer(256)
	for i := 0; i < b.N; i++ {
		q.Palette(m)
	}
}

func BenchmarkEsimov(b *testing.B) {
	m := getImage(b)
	q := colorquant.Quant{}
	for i := 0; i < b.N; i++ {
		q.Quantize(m, 256)
	}
}
