package main

import (
	"fmt"
	"time"
)

func solve(i int) {
	fmt.Println(i)
}

func main() {

	//using go routines to run process concurrently

	for i := 0; i <= 10; i++ {
		// go solve(i)
		//using inline function
		go func(i int) {
			fmt.Println(i)
		}(i)

	}

	time.Sleep(time.Second * 2)
}
