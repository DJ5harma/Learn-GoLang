package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}
	var x = post{}

	for range 100 {
		wg.Go(func() {
			defer x.mu.Unlock()
			x.mu.Lock()
			x.views++
		})
	}

	wg.Wait()
	fmt.Println("Final val:", x.views)
}
