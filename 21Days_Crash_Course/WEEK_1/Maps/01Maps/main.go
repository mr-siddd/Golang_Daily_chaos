package main

import "fmt"

func main() {

	var dict map[string]string

	key := "good"

	value := dict[key]

	fmt.Printf("%q means %#v \n", key, value)
	fmt.Printf("# of Keys: %d\n", len(dict))
}
