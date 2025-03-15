package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println

	// Index: Find the index of the first occurrence of "e" in "test"
	p("Index:", strings.Index("test", "e"))

	// Join: Join two strings with a hyphen
	p("Join:", strings.Join([]string{"a", "b"}, "-"))

	// Repeat: Repeat the string "a" five times
	p("Repeat:", strings.Repeat("a", 5))

	// Replace: Replace "o" with "0" in "foo"
	p("Replace:", strings.Replace("foo", "o", "0", -1))

	// Split: Split the string "a-b-c-d-e" by "-"
	p("Split:", strings.Split("a-b-c-d-e", "-"))

	// ToLower: Convert "TEST" to lowercase
	p("ToLower:", strings.ToLower("TEST"))

	// ToUpper: Convert "test" to uppercase
	p("ToUpper:", strings.ToUpper("test"))

	// Len: Get the length of "hello"
	p("Len:", len("hello"))

	// Char: Get the second character from "hello"
	p("Char:", string("hello"[1]))

	// Contains: Check if "es" is a substring of "test"
	p("Contains:", strings.Contains("test", "es"))

	// Count: Count how many times "t" appears in "test"
	p("Count:", strings.Count("test", "t"))

	// HasPrefix: Check if "test" starts with "te"
	p("HasPrefix:", strings.HasPrefix("test", "te"))

	// HasSuffix: Check if "test" ends with "st"
	p("HasSuffix:", strings.HasSuffix("test", "st"))
}
