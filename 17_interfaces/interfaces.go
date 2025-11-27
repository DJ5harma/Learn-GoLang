package main

import "fmt"

type Igateway interface {
	pay(amount float32)
}

type payment struct {
	gateway Igateway
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct {
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway", amount)
}

func (r *stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
	// fakepaymentGw := fakepayment{}
	payment1 := payment{gateway: razorpayPaymentGw}

	payment1.makePayment(1000)
}
