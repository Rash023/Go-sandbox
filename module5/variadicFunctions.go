package main

import "fmt"

//variadic functions are the functions which can accept n number of parameters

func sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

//variadic function to accept any data type
func all(nums ...interface{}) {
	for _, k := range nums {
		fmt.Println(k)
	}
}
func main() {
	nums := []int{4, 5, 6}
	res := sum(1, 2, 3, 4)
	all(nums)
	fmt.Println(res)
}
