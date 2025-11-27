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
	createdAt time.Time //nanosecond
	customer  Customer
}

func newOrder(id string, amount float32, status string) *Order {
	newOrder := Order{id: id, amount: amount, status: status}
	return &newOrder
}

func (O *Order) changeStatus(status string) {
	O.status = status
}
func (O Order) getAmount() float32 {
	return O.amount
}

func main() {

	// order1 := Order{id: "1", amount: 1000, status: "received"}

	// fmt.Println(order1)

	// order1.changeStatus("confirmed")
	// order1.createdAt = time.Now()
	// fmt.Println(order1)
	// fmt.Println(order1.getAmount())

	// order2 := newOrder("2", 30.50, "received")
	// fmt.Println(order2)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{name: "Sunu", isGood: true}

	// fmt.Println(language)

	order3 := Order{
		id:     "3",
		amount: 30,
		status: "error",
		customer: Customer{
			name:  "Sunu",
			phone: "213423",
		},
	}
	fmt.Println(order3)

}
