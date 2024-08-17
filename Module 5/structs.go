package main

import (
	"fmt"
	"time"
)

// along with structures we use struct for object oriented programming also
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// acts a constructor for go
func newOrder(id string, amount float32, status string) *Order {
	order := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}

// acts the same as a function inside a class
func (O *Order) changeStatus(status string) {
	O.status = status

}

func main() {
	order1 := Order{
		id:     "1",
		amount: 50,
		status: "Pending",
	}

	order2 := *newOrder("2", 25.00, "Pending")
	fmt.Println(order2)
	order2.changeStatus("Delivered")
	fmt.Println(order2)

	order1.createdAt = time.Now()
	order1.changeStatus("Delivered")
	fmt.Println("Order Struct", order1)

	//shorthand declaration of structs just like objects in ts
	language := struct {
		name   string
		status bool
	}{"Golang", true}

	fmt.Println(language)
}
