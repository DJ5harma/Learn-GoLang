package main

import "fmt"

func printSlice[T comparable, V string](items []T, v V) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println(v)
}

type stack[T string | int] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"go", "ts", "js", "py"}
	booleans := []bool{true, false, false, true}

	printSlice(nums, "NUMS")
	printSlice(names, "NAMES")
	printSlice(booleans, "BOOLEANS")

	// stack1 := stack[int]{elements: []int{1, 2, 3}}

	// fmt.Println(stack1)

}
