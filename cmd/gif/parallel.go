package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/schollz/progressbar"
)

func newLabeledProgressBar(n int, label ...string) *progressbar.ProgressBar {
	var option progressbar.Option
	if len(label) == 0 {
		option = progressbar.OptionSetDescription(fmt.Sprintf("%s%d", config.CliOptions, n))
	} else {
		option = progressbar.OptionSetDescription(fmt.Sprintf("%s%d %v", config.CliOptions, n, label))
	}
	return progressbar.NewOptions(n, progressbar.OptionSetWriter(os.Stderr), option)
}

func parallel(n int, f func(int), label ...string) {
	work := make(chan int, 4)
	var wg sync.WaitGroup
	var bar *progressbar.ProgressBar
	if len(label) > 0 {
		bar = newLabeledProgressBar(n, label...)
		bar.Clear()
		bar.RenderBlank()
	}

	wg.Add(config.Parallelism)
	for i := 0; i < config.Parallelism; i++ {
		go func() {
			for i := range work {
				f(i)
				if bar != nil {
					bar.Add(1)
				}
			}
			wg.Done()
		}()
	}
	for i := 0; i < n; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
	if bar != nil {
		bar.Finish()
	}
}

func sequential(n int, f func(int), label ...string) {
	bar := newLabeledProgressBar(n, label...)
	bar.Clear()
	bar.RenderBlank()
	for i := 0; i < n; i++ {
		f(i)
		bar.Add(1)
	}
	bar.Finish()
}
