package main

import "fmt"

//can be declared outside the function also
const price = 10

func main() {
	//once declared values cannot be modified
	const name string = "john"
	const age = 30
	const float = 10.0
	fmt.Println(price)

	//constant grouping (mainly used for configs)
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}
