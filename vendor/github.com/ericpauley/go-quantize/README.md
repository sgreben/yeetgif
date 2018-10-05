# go-quantize
go-quantize is a highly-optimized and memory-efficient palette generator. It currently implements the Median Cut algorithm, including weighted color priority.

## Performance
go-quantize makes exactly two slice allocations per palette generated, the larger of which is efficiently pooled. It also uses performant direct pixel accesses for certain image types, reducing memory footprint and increasing throughput.

## Benchmarks
go-quantize performs significantly faster than existing quantization libraries:

```
# bench/bench_test.go
BenchmarkQuantize-8          	      50	  20070550 ns/op	  122690 B/op	     258 allocs/op
BenchmarkSoniakeysMedian-8   	       3	 465833354 ns/op	 3479624 B/op	     782 allocs/op
BenchmarkSoniakeysMean-8     	       3	 342759921 ns/op	 2755712 B/op	     262 allocs/op
BenchmarkEsimov-8            	       2	 645129392 ns/op	35849608 B/op	 8872273 allocs/op
```

## Example Usage
```go
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
```
