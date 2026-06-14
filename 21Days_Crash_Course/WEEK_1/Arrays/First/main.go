package main

import "fmt"

func main() {
	speed := 10

	fmt.Println(speed)

	var FruitList [5]string
	FruitList[0] = "Apple"
	FruitList[1] = "Banana"

	FruitList[2] = "Elderberry"

	fmt.Println(FruitList)
	fmt.Println(len(FruitList))

	var names [3]string
	names[len(names)-1] = "!"
	names[1] = "think" + names[2]
	names[0] = "Don't"
	names[0] += " "

	fmt.Println(names[0] + names[1] + names[2])
}
