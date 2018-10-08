package main

import "sync"

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
