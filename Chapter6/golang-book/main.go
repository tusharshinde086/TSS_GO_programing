package main

import (
	"C:/TSS_GO_Programing/Chapter6/golang-book/programgo/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
