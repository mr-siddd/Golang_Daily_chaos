package utils

import "fmt"

func CheckOk(ok bool) bool {
	if !ok {
		fmt.Println("❌ Error in operation")
		return false
	}
	return true
}
