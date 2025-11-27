package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {
	for i := range 10 {
		go task(i + 1)

		go func(i int) {
			fmt.Println("item", i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
