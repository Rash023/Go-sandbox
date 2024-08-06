package main

import "fmt"

func changeNum(num *int) {
	*num = 5

}

func changeSlice(nums *[]int) {
	*nums = append(*nums, 4)
}
func main() {
	num := 1
	fmt.Println("Number before change", num)
	changeNum(&num)
	fmt.Println("Number after change", num)

	var nums = []int{1, 2, 3}
	fmt.Println("Befores updating the array", nums)
	changeSlice(&nums)
	fmt.Println("After updating the array", nums)
}
