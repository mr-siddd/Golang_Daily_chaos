package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Finder")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a Directory")
		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Type().IsRegular() {
			name := file.Name()
			fmt.Println(name)
		}
		//fmt.Println(file.Name())
		//fmt.Println(file.Mode())
	}
}
