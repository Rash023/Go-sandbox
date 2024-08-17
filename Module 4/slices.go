package main

import (
	"fmt"
	"slices"
)

//slice -> dynamic
// most used construct in go

func main() {

	// var nums []int
	// fmt.Println(nums == nil)

	//make function arguments meaning (type,initial empty array size, dynamic default capacity)
	// var nums = make([]int, 2, 5)
	// fmt.Println(cap(nums))

	// //append a value into the existing array
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)
	// nums = append(nums, 3)

	// //capacity doubles everytime the limit is reached
	// fmt.Println(cap(nums))

	//more simpler and better  way to initialize a dynamic array
	// nums := []int{}
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums[0] = 1
	// fmt.Println(nums)

	// //copy function
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice operator
	var nums = []int{1, 2, 3}
	fmt.Println(nums[0:1])

	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	//checks if both the arrays are equal are not
	fmt.Println(slices.Equal(nums, nums2))

	//creating 2d arrays
	var twoD = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(twoD)
}
