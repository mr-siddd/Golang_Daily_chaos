package main

import (
	"fmt"
)

func SomeFunc(number string) {
	fmt.Println("The number is ", number)
}

func main() {
	go SomeFunc("2")
	go SomeFunc("4")
	go SomeFunc("6")

	//time.Sleep(1 * time.Second) // Sleep to allow goroutines to finish before the main function exits
	fmt.Println("HEY SID!")
}
