package main

import (
	"Employee_Manager/internals/employee"
	"fmt"
)

func main() {
	fmt.Println("👨‍💼 Welcome to Employee Manager CLI")
	defer fmt.Println("👋 Thank you for using Employee Manager CLI")

	for {
		employee.ShowMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			employee.AddEmployee()
		case 2:
			employee.ListEmployees()
		case 3:
			employee.UpdateSalary()
		case 4:
			employee.DeleteEmployee()
		case 5:
			employee.ShowTotalCount()
		case 6:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}
