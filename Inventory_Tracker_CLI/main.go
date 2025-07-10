/*Inventory Products Tracker CLI Application*/
package main

import "fmt"

func main() {

	fmt.Println("Welcome to Inventory Products Tracker CLI Application")
	var inventory []InventoryStructure // Slice to hold the inventory products
	defer fmt.Println("Thankyou for using Inventory Products Tracker CLI Application")

	DisplayMenu()
	for {
		var choices int

		fmt.Print("Enter your choice : ")
		fmt.Scanln(&choices)
		switch choices {
		case 1:
			AddingProducts(&inventory) // Function to add products to the inventory
		case 2:
			ListProducts(&inventory) // Function to list all products in the inventory
		case 3:
			UpdateProductQuantity(&inventory) // Function to update the quantity of a product
		case 4:
			DeleteProduct(&inventory)
		case 5:
			TotalInventoryValue(&inventory) // Function to calculate the total value of the inventory
		case 6:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

	}

}
