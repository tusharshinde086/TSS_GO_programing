package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.HasSuffix("test", "st")
	fmt.Println("HasSuffix:", result)
}
