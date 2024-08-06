package main

import "fmt"

func Counter() func() int {
	var Count int = 0
	return func() int {
		Count += 1
		return Count
	}
}

func main() {

	Increment := Counter()
	fmt.Println(Increment())
	fmt.Println(Increment())

}
