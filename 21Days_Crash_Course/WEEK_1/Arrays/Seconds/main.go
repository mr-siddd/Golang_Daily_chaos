package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var MoodyArray = [10]string{"Happy", "Sad", "Angry", "Excited", "Bored", "Nervous", "Calm",
		"Anxious", "Confident", "Tired"}

	fmt.Printf("Data is %q\n", MoodyArray)
	rand.Seed(time.Now().UnixNano())
	Index := rand.Intn(len(MoodyArray))

	if len(os.Args) > 1 {
		Name := os.Args[1]
		fmt.Printf("%s feels %s", Name, MoodyArray[Index])
	} else {
		fmt.Println("Your Name")
	}
}

// func main() {
// 	names := [...][3]string{
// 		{"First Name", "Last Name", "Nickname"},
// 		{"Albert", "Einstein", "emc2"},
// 		{"Isaac", "Newton", "apple"},
// 		{"Stephen", "Hawking", "blackhole"},
// 		{"Marie", "Curie", "radium"},
// 		{"Charles", "Darwin", "fittest"},
// 	}

// 	for i := range names {
// 		n := names[i]
// 		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

// 		if i == 0 {
// 			fmt.Println(strings.Repeat("=", 87))
// 		}
// 	}
// }
