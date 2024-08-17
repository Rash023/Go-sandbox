package main

import "fmt"

//function returning single value
func add(a int, b int) int {
	return a + b
}

//in go functions can return multiple values also
//which makes it very easy for error handling for tightly coupled functions
func calc(a int, b int) (int, int, int) {
	return a + b, a * b, a / b
}

func process(fn func(a int, b int) int) {
	fmt.Println(fn(5, 2))

}

func processIt(ops string) func(a int, b int) int {
	if ops == "add" {
		return func(a, b int) int {
			return a + b
		}
	} else {
		return func(a, b int) int {
			return a - b
		}
	}
}

func main() {
	a := 9
	b := 5
	//if we want to suppress a value returned by the function we can '_' so that go compiler doesnt complain
	sum, mult, div := calc(a, b)
	sum1, mult1, _ := calc(10, 5)

	fmt.Println(sum, mult, div)
	fmt.Println(sum1, mult1)

	//function as parameters inside function
	process(add)

	//function that return a function
	operation := "sub"
	fn := processIt(operation)
	fmt.Println(fn(2, 3))
}
