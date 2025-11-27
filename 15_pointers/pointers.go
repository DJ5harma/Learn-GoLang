package main

import "fmt"

func changeNum(num *int) {
	*num++
	fmt.Println(*num)
}

func main() {

	num := 10
	changeNum(&num)
	changeNum(&num)
	changeNum(&num)
	fmt.Println("after change", num)
}
