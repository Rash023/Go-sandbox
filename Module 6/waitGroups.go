package main

import (
	"fmt"
	"sync"
)

func check(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task:", i)
}

func main() {

	//using go routines to run process concurrently
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go check(i, &wg)
	}

	wg.Wait()
}
