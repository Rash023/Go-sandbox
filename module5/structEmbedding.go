package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	Customer
}

func newOrder(id string, amount float32, status string) *Order {
	order1 := Order{
		id:       id,
		amount:   amount,
		status:   status,
		Customer: Customer{name: "Rashid", phone: "7802313210"},
	}
	return &order1
}

func main() {
	order1 := *newOrder("1", 20.0, "Pending")
	fmt.Println(order1)
	order1.Customer.name = "Mazhar"
	fmt.Println(order1)

}
