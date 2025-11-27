package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// communication queue bw goroutines

func processNum(numChan chan int) {

	for {
		fmt.Println("processing number", <-numChan)
	}
}

func main() {
	wg := sync.WaitGroup{}

	numChan := make(chan int)

	wg.Go(func() { processNum(numChan) })

	for {
		time.Sleep(time.Second)
		numChan <- rand.Intn(100)
	}

}
