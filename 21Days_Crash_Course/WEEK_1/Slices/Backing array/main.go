package main

import (
	"fmt"
)

func main() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// ONLY ADD YOUR CODE HERE
	//
	// Ensure that nums slice never changes even though
	// the mine slice changes.
	mine := make([]int, len(nums))
	copy(mine, nums)
	//mine := nums
	// ----------------------------------------

	mine := append([]int(nil), nums[:3]...)

	// DON'T TOUCH THE FOLLOWING CODE
	//
	// This code changes the elements of the nums
	// slice.
	//
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
