package main

import "fmt"

func main() {

	age := 16

	if age >= 18 {
		fmt.Println("User is teen adult!")

	} else if age >= 12 {
		fmt.Println("User is an teenagers!")
	} else {
		fmt.Println("User is not an adult!")
	}

	var role = "Admin"
	var permissions = true

	if role == "Admin" && permissions {
		fmt.Println("User authorized!")
	}

	//declaring and adding conditionals all at one go
	if price := 45; price > 30 {
		fmt.Println("The item is costly!", price)
	} else if price >= 100 {
		fmt.Println("Overpriced!")
	} else {
		fmt.Println("Price is under budget!")
	}

	//go does not have ternary operator
}
