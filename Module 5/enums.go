package main

import "fmt"

type orderStatus int

const (
	Recieved orderStatus = iota
	Pending
	Confirmed
	Delivered
)

func updateOrderStatus(status orderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	updateOrderStatus(Recieved)

}
