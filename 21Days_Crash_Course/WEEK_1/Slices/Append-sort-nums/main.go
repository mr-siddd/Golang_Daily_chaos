package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*Append#4 --> Append and sort numbers*/

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("provide a few numbers")
		return
	}

	var numbers []int
	for _, s := range args {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	fmt.Println(numbers)
}
