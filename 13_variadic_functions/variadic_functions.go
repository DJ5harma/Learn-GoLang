package main

import "fmt"

func sum(args ...int) int {
	sum := 0
	for _, num := range args {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{3, 4, 5, 6}
	fmt.Println(sum(3, 4, 5, 6))
	fmt.Println(sum(nums...))
}
