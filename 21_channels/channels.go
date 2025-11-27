package main

import (
	"fmt"
	"time"
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

func task(doneChannel chan bool) {
	defer func() { doneChannel <- true }()

	fmt.Println("processing...")
}

func main() {
	// wg := sync.WaitGroup{}

	// numChan := make(chan int)

	// wg.Go(func() { processNum(numChan) })

	// for {
	// 	time.Sleep(time.Second)
	// 	numChan <- rand.Intn(100)
	// }

	// ------------------------------

	// sumChannel := make(chan int)

	// go sumGoRoutine(sumChannel, 5, 4)

	// result := <-sumChannel

	// fmt.Println("result", result)

	// ------------------------------

	// doneChannel := make(chan bool)
	// go task(doneChannel)

	// received := <-doneChannel
	// fmt.Println("Complete", received)

	// // -------------------------------
	// // Buffered Channel (limited amount of data without blocking)
	// doneChannel := make(chan bool)

	// emailChannel := make(chan string, 100) // buffered channel )has size, can send 100 elements without blocking

	// go emailSender(emailChannel, doneChannel)
	// for i := range 6 {
	// 	emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// close(emailChannel)

	// <-doneChannel

	// ------------------------------------------ multi channel listening

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "pong"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan 2", chan2Val)
		}
	}

}

// also type safety in channel
func emailSender(emailChan <-chan string, doneChannel chan<- bool) {
	defer func() { doneChannel <- true }()
	for email := range emailChan {
		time.Sleep(time.Second)
		fmt.Println("Sending email to:", email)
	}
}
