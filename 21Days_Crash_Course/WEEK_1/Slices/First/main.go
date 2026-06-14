package main

import (
	"fmt"
)

/*Appending Two Slices and using prettySlice*/

// func main() {

// 	var TODO []string
// 	TODO = append(TODO, "Gym at 09:45 PM")
// 	TODO = append(TODO, "Read a book", "Go for a walk", "Cook dinner")
// 	s.Show("todo", TODO)

// 	tomorrowTasks := []string{"See Mom", "Learn Golang"}
// 	TODO = append(TODO, tomorrowTasks...)

// 	s.Show("todo", TODO)

// }

/*Append #1 --> Append and compare byte slices*/
// func main() {
// 	PNG, header := []byte{'P', 'N', 'G'}, []byte{}
// 	Result := append(header, PNG...)

// 	s.Show("Result", Result)
// 	if bytes.Equal(PNG, Result) {
// 		fmt.Printf("They are Equal")
// 	} else {
// 		fmt.Printf("They are Not Equal")
// 	}

// }

/*Append #3 --> Fix problem from the code*/
func main() {

	// toppings := []int{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, extra cheese)

	// fmt.Printf("pizza       : %s\n", pizza)
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(toppings, "onions", "extra cheese")

	fmt.Printf("pizza       : %s\n", pizza)
}
