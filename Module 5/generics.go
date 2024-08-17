package main

import "fmt"

func printSlice[T any](nums []T) {
	for _, i := range nums {
		fmt.Println(i)
	}
}

//for more strict generics
func strictSlice[T int | string](nums []T) {
	for _, i := range nums {
		fmt.Println(i)
	}
}

//using generics over struct
type Stack[T any] struct {
	elements []T
}

func main() {
	// names := []string{"golang", "typescript"}
	// nums := []int{1, 2, 3, 4}

	// printSlice(nums)
	// strictSlice(nums)

	newStack := Stack[int]{
		elements: []int{1, 2, 3},
	}

	newStack2 := Stack[string]{
		elements: []string{"golang", "typescript"},
	}
	fmt.Println(newStack)
	fmt.Println(newStack2)

}
