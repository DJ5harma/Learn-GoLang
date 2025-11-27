package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Failed
	Delivered
)

type PaymentStatus string

const (
	Done    PaymentStatus = "Done"
	Pending PaymentStatus = "Pending"
	Error   PaymentStatus = "Error"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing status to ", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
