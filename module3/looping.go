package main

import "fmt"

//for->only construct for looping in go

func main() {

	//implementing while loop using for construct
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite while loop
	for {
		break
		//logic
	}

	//classic for loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		continue
	}

	//using range keyword
	for i := range 3 {
		fmt.Println(i)
	}
}
