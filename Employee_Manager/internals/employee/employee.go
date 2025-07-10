package employee

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

var employees []Employee
var idCounter int = 1

func ShowMenu() {
	fmt.Println("\n========================")
	fmt.Println("📋 Employee Manager Menu")
	fmt.Println("1. Add Employee")
	fmt.Println("2. List All Employees")
	fmt.Println("3. Update Salary")
	fmt.Println("4. Delete Employee")
	fmt.Println("5. Show Total Employee Count")
	fmt.Println("6. Exit")
	fmt.Println("========================")
}

func AddEmployee() {
	var name string
	var age int
	var salary float64

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter salary: ")
	fmt.Scanln(&salary)

	newEmp := Employee{
		ID:     idCounter,
		Name:   name,
		Age:    age,
		Salary: salary,
	}

	employees = append(employees, newEmp)
	idCounter++

	fmt.Println("✅ Employee added successfully!")
}

func ListEmployees() {
	if len(employees) == 0 {
		fmt.Println("🚫 No employees found.")
		return
	}
	fmt.Println("📄 Current Employees:")
	for _, emp := range employees {
		fmt.Printf("ID: %d | Name: %s | Age: %d | Salary: ₹%.2f\n",
			emp.ID, emp.Name, emp.Age, emp.Salary)
	}
}

func UpdateSalary() {
	var id int
	fmt.Print("Enter employee ID to update salary: ")
	fmt.Scanln(&id)

	for i := range employees {
		if employees[i].ID == id {
			var newSalary float64
			fmt.Print("Enter new salary: ")
			fmt.Scanln(&newSalary)
			employees[i].Salary = newSalary
			fmt.Println("✅ Salary updated successfully!")
			return
		}
	}
	fmt.Println("❌ Employee not found.")
}

func DeleteEmployee() {
	var id int
	fmt.Print("Enter employee ID to delete: ")
	fmt.Scanln(&id)

	for i := range employees {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("🗑️ Employee deleted successfully.")
			return
		}
	}
	fmt.Println("❌ Employee not found.")
}

func ShowTotalCount() {
	fmt.Printf("👥 Total number of employees: %d\n", len(employees))
}
