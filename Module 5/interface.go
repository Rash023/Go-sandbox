package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type Payment struct {
	gateway paymenter
}

func (p Payment) makePayment(amount float32) {
	p.gateway.pay(100)
}

type Razorpay struct{}

func (r Razorpay) pay(amount float32) {
	fmt.Println("Payment using Razorpay of amount", amount)
}

type Stripe struct{}

func (s Stripe) pay(amount float32) {
	fmt.Println("Payment using Stripe of amount", amount)
}

type MockGateway struct{}

func (m MockGateway) pay(amount float32) {
	fmt.Println("Payment using Mock Gateway of amount", amount)
}
func main() {
	// stripePaymentGw := Stripe{}
	// razorpayPaymentGw := Razorpay{}
	mockPaymentGw := MockGateway{}
	newPayment := Payment{
		gateway: mockPaymentGw,
	}

	newPayment.makePayment(100)
}
