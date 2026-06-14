package main

import "fmt"

func main() {

	myChannel := make(chan string)      // Create a channel of type string
	anotherChannel := make(chan string) // Create another channel of type string

	go func() {
		myChannel <- "Hello from the goroutine!" // Send a message to the channel via
		// goroutine that is child goroutine
	}()

	go func() {
		anotherChannel <- "Hello from another goroutine!" // Send a message to the another channel via
		// another goroutine that is child goroutine
	}()

	select {
	case msgFromAnotherChannel := <-anotherChannel: // Wait for a message from anotherChannel
		fmt.Printf("Received from anotherChannel :%s \n", msgFromAnotherChannel)
	case msgFromMyChannel := <-myChannel: // Wait for a message from myChannel
		fmt.Printf("Received from myChannel: %s\n", msgFromMyChannel)
	}
	// msg := <-myChannel // Receive the message from the channel in
	// // the main function that is main goroutine
	// fmt.Printf(msg)

}
