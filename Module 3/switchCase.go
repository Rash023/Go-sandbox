package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 3
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	//multiple conditions of switch

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")

	}

	//type switch

	whoAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")

		case string:
		case bool:
			fmt.Println("Boolean")

		default:
			fmt.Println("Other", t)
		}
	}
	whoAmi(true)
}
