package main

import (
	"fmt"
)

// communication bw goroutines

func processNum(numChan chan int) {

	for {
		fmt.Println("processing number", <-numChan)
	}
}

func sumGoRoutine(sumChannel chan int, num1 int, num2 int) {
	sumChannel <- num1 + num2
}

func main() {
	// wg := sync.WaitGroup{}

	// numChan := make(chan int)

	// wg.Go(func() { processNum(numChan) })

	// for {
	// 	time.Sleep(time.Second)
	// 	numChan <- rand.Intn(100)
	// }

	sumChannel := make(chan int)

	go sumGoRoutine(sumChannel, 5, 4)

	result := <-sumChannel

	fmt.Println("result", result)

}
