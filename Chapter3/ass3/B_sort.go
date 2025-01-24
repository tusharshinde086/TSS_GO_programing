package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{12, 5, 9, 1, 15, 3}

	sort.Ints(arr)

	fmt.Println("Sorted array in ascending order:", arr)
}
