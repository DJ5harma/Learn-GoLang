package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() //like useEffect's return
	fmt.Println("Doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go task(i+1, &wg)
	}
	wg.Wait()

}
