package main

import "fmt"

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}

type stack[T string | int] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"go", "ts", "js", "py"}
	booleans := []bool{true, false, false, true}

	printSlice(nums)
	println()
	printSlice(names)
	println()
	printSlice(booleans)

	// stack1 := stack[int]{elements: []int{1, 2, 3}}

	// fmt.Println(stack1)

}
