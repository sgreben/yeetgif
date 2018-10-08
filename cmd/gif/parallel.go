package main

import (
	"os"
	"github.com/schollz/progressbar"
	"sync"
)

func parallel(n int, f func(int)) {
	work := make(chan int, 4)
	var wg sync.WaitGroup
	bar := progressbar.NewOptions(n, progressbar.OptionSetWriter(os.Stderr), progressbar.OptionSetDescription(config.CliOptions))
	bar.RenderBlank()

	wg.Add(config.Parallelism)
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
