package main

import (
	"fmt"
)

// Student struct to hold student details
type Student struct {
	RollNo   int
	StudName string
	Mark1    float64
	Mark2    float64
	Mark3    float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	// Accept student details
	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].StudName)
		fmt.Print("Mark 1: ")
		fmt.Scan(&students[i].Mark1)
		fmt.Print("Mark 2: ")
		fmt.Scan(&students[i].Mark2)
		fmt.Print("Mark 3: ")
		fmt.Scan(&students[i].Mark3)
	}

	// Calculate total and average marks
	for _, student := range students {
		total := student.Mark1 + student.Mark2 + student.Mark3
		average := total / 3
		fmt.Printf("Student: %s (Roll No: %d) - Total Marks: %.2f, Average Marks: %.2f\n", student.StudName, student.RollNo, total, average)
	}
}
