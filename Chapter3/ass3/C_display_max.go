package main

import (
	"fmt"
)

// Employee structure
type Employee struct {
	eno    int
	ename  string
	salary float64
}

// Function to find employees with maximum salary
func findMaxSalaryEmployees(employees []Employee) []Employee {
	maxSalary := employees[0].salary
	var maxEmployees []Employee

	// Find max salary
	for _, emp := range employees {
		if emp.salary > maxSalary {
			maxSalary = emp.salary
		}
	}

	// Collect employees with max salary
	for _, emp := range employees {
		if emp.salary == maxSalary {
			maxEmployees = append(maxEmployees, emp)
		}
	}

	return maxEmployees
}

func main() {
	var n int

	// Accept number of employees
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	// Input employee records
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].eno)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].ename)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].salary)
	}

	maxSalaryEmployees := findMaxSalaryEmployees(employees)

	fmt.Println("\nEmployee(s) with Maximum Salary:")
	for _, emp := range maxSalaryEmployees {
		fmt.Printf("Emp No: %d, Name: %s, Salary: %.2f\n", emp.eno, emp.ename, emp.salary)
	}
}
