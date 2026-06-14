package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ConversionTable = [8]string{"0.84", "0.62", "1.31", "0.39", "0.27", "2.20", "3.28", "1.10"}
	var Currencies = [8]string{"EURO", "GBP", "AUD", "CAD", "CNY", "INR", "MXN", "JPY"}
	fmt.Printf("%-20s %-20s\n", "Currency", "Conversion Rate to USD")
	fmt.Println("========================================")

	var Number float64
	Number, _ = strconv.ParseFloat(os.Args[1], 64)

	if len(os.Args) > 1 {
		for i := 0; i < len(Currencies); i++ {
			rate, _ := strconv.ParseFloat(ConversionTable[i], 64)
			fmt.Printf("%.2f, USD is %.2f %s \n", Number, Number*rate, Currencies[i])
		}
	} else {
		fmt.Println("Please provide number")
	}

}
