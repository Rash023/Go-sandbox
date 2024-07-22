package main

import "fmt"

func main() {
	//used for iterating over data structures
	nums := []int{1, 2, 4}

	//simple for loop iteration
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//iterating with range keyword
	for i, val := range nums {
		fmt.Println(i, val)
	}

	//iterating over the map using range function
	mp := map[string]int{"price": 100, "phone": 12345}

	for key, val := range mp {
		fmt.Println(key, val)
	}

	//iterating over the string using range function
	//by default the first variable returns the byte address from the starting point of the string
	//second variable returns the unicode code point rune
	for i, c := range "golang" {
		//explicity convert to string to get the char value from unicode point
		fmt.Println(i, string(c))
	}

}
