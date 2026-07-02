package main

import (
	"fmt"
)

type SportsActivity interface {
	Swim() bool
}

type Man struct {
	cansSwim bool
}

func (men Man) Swim() bool {
	return men.cansSwim
}

type Dog struct {
	DogcansSwim bool
}

func (dog Dog) Swim() bool {
	return dog.DogcansSwim
}

func checkSwimmings(s SportsActivity) {
	fmt.Println("Can Swim: ", s.Swim())

}

func main() {

	fmt.Println("Implementation of Interface in Golang")

	checkSwimmings(Man{cansSwim: true})
	checkSwimmings(Dog{DogcansSwim: false})
}
