/////////////////exo3.go///////////////////////////
///objective(s):
///				- Create Slice of xx integer
///				- Assign a value for each
///				- Use range to browse the slice
///				- Print out the slice content
///				- Add 4 values on the original slice
///				- Browse the slice and print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	originalslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //Declare the original slice
	fmt.Printf("Original Slice\n")
	fmt.Printf("Indice\t\tValue\n")
	for indice, value := range originalslice {
		fmt.Printf("%d\t\t%d\n", indice, value)
	}
	fmt.Printf("\n\n")
	newslicevalues := []int{10, 11, 12, 13, 14, 15} //Declare a new slice of int with additional values
	fmt.Printf("New Slice\n")
	fmt.Printf("Indice\t\tValue\n")
	originalslice = append(originalslice, newslicevalues...) //Add new values to the original slice
	for indice, value := range originalslice {
		fmt.Printf("%d\t\t%d\n", indice, value)
	}

}
