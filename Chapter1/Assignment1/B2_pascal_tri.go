package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of rows for Pascal's Triangle:")
	fmt.Scan(&n)

	for line := 0; line < n; line++ 
	
	{

		// spaces 

		for space := 0; space < n-line-1; space++ {
			fmt.Print(" ")
		}

		//  values 	
		
                val := 1
	            	for i := 0; i <= line; i++ {
	            		fmt.Print(val, " ")
	            		val = val + 1
	            	}
	            	fmt.Println()
	}
}
