package main

import "fmt"

func main() {
	//declaring empty array
	var nums [3]int

	//array length
	fmt.Println(len(nums))

	nums[0] = 1

	//print the single element
	fmt.Println(nums[0])
	//print the whole array
	fmt.Println(nums)
	var vals [4]bool
	fmt.Println(vals)

	//declaring array with values
	test := [3]int{1, 2, 3}
	fmt.Println(test)

	//declaring a 2d array with values
	Array := [3][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(Array)

}
