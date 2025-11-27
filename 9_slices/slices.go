package main

import "fmt"

func main() {

	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 5, 10)
	// var nums []int
	// nums := []int{}
	// clear(nums)

	// nums = append(nums, 8)
	// nums = append(nums, 8)
	// nums = append(nums, 8)
	// nums = append(nums, 8)
	// nums = append(nums, 8)
	// nums = append(nums, 8)

	// fmt.Println("length:", len(nums))
	// fmt.Println("capacity:", cap(nums))

	// var nums2 = nums
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// nums2[2] = 9
	// fmt.Println(nums)
	// fmt.Println(nums2)

	// var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// var arr2 = arr[2:5]
	// arr2[0] = 100
	// fmt.Println(arr)
	// fmt.Println(arr2)

	// var nums1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// var nums2 = []int{1, 2}

	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{}, {}}

	nums[0] = append(nums[0], 9)
	fmt.Println(nums)
}
