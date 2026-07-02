package main

import (
	"fmt"
)

type OrderStatus interface {
	OrderValidation() bool
	GetStatus() string
}

type Order struct {
	OrderID      int
	OrderStatus  string
	Transcode    int
	ExchangeFlag string
}

func (o Order) OrderValidation() bool {

	counter := 0

	if o.OrderID > 9999999 {
		fmt.Println("Order ID is invalid")
		counter++

	}
	if o.OrderStatus != "Completed" {
		fmt.Println("Order Status is invalid")
		counter++
	}
	if o.Transcode != 2073 && o.ExchangeFlag != "N" {
		fmt.Println("Transcode is not Updated, Eventhough order is coming from Exchange")
		counter++
	}

	if o.ExchangeFlag != "Y" {
		fmt.Println("Exchange Flag is not Updated")
		counter++
	}

	if counter > 0 {
		fmt.Println("Order is invalid, Please check the order details")
		return false
	}

	return true

}

func (o Order) GetStatus() string {
	return o.OrderStatus
}

func OrderStatusCheck(o OrderStatus) {
	fmt.Println("Order Validation: ", o.OrderValidation())
	fmt.Println("Current Order Status: ", o.GetStatus())

}

func main() {
	fmt.Println("Order Management System")

	OrderStatusCheck(Order{OrderID: 12345689, OrderStatus: "Completed", Transcode: 2000, ExchangeFlag: "Y"})

}
