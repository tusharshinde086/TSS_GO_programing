package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func WriteXML(filename string, person Person) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding XML:", err)
		return
	}

	fmt.Println("XML file written successfully.")
}

func main() {
	person := Person{Name: "John Doe", Age: 30}
	WriteXML("data.xml", person)
}
