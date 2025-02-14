package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (s *Student) Show() {
	fmt.Printf("Student Name: %s\nAge: %d\nGrade: %s\n", s.Name, s.Age, s.Grade)

}
func main() {
	student := &Student{Name: "Alice", Age: 20, Grade: "A"}
	student.Show()
}
