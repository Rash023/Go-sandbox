package main

import (
	"fmt"
	"time"
)

// channel helps sending data from one go routine to another go routine

// send data in the channel
func PrintChannel(numChan chan int) {
	for num := range numChan {
		fmt.Println("Number Recieved", num)
		time.Sleep(time.Second)
	}
}

// receive data in the channel
func Calc(numChan chan int, num1 int, num2 int) {
	sum := num1 + num2
	numChan <- sum

}

// channel as waitgroup
func Process(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}

// buffered channel
// increasing type safety we use array notations
// marking that the chan can only recieve or send
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println(email)
		time.Sleep(time.Second)
	}
}
func main() {

	// //channel to send data
	// numChan := make(chan int)
	// go PrintChannel(numChan)
	// for {
	// 	numChan <- rand.Intn(100)

	// }

	// numChan := make(chan int)

	// go Calc(numChan, 4, 5)

	// res := <-numChan
	// fmt.Println(res)

	//using channel as wait groups
	// done := make(chan bool)
	// go Process(done)

	// <-done //block

	//using the above implementation of channel we can only process one data at a time
	//as it blocks the channel once recieved and you cannot process the channel further
	//so to overcome this we use buffered channels which can recieve a given amount of
	//data specified without blocking the channel, below is the implementation
	emailChan := make(chan string, 10)
	done := make(chan bool)
	go emailSender(emailChan, done)
	for i := 0; i <= 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)

	}
	close(emailChan)
	<-done

}
