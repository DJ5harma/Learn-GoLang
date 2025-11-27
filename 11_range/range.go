package main

import "fmt"

func main() {

	var nums = make([]int, 0)
	nums = append(nums, 100)
	nums = append(nums, 200)
	nums = append(nums, 300)

	// var sum = 0

	// for i, num := range nums {
	// 	// sum += nums[i]
	// 	sum += num
	// 	fmt.Println(i, num)
	// }
	// fmt.Println("Sum:", sum)

	// var mp = map[int]string{1: "Sunu", 21: "Oshu", 30: "Pablu"}

	// for key, value := range mp {
	// 	fmt.Println(key, value)
	// }

	// // starting byte of rune, unicode code point rune
	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
}
