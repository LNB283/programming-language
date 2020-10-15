/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create Slice of xx integer
///				- Assign a value for each
///				- Use Slice capabilities to create new slices from the original slice
///					--create a slice with the first 3 values
///					--create a slice with the values from the position 4
///					--create a slice with the value between the position 4 to 7
///					--create a slice with the last 3 values
///				- Print out the result for each
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	myslice := []int{23, 8, 33, 18, 12, 56, 24, 99, 7, 9, 12, 45}

	fmt.Printf("indice\t\tval\n")
	for indice, val := range myslice {
		fmt.Printf("%d\t\t%d\n", indice, val)
	}
	fmt.Printf("\n")
	fmt.Printf("First 3 values:%v\n", myslice[:3])
	fmt.Printf("Values from the position 4:%v\n", myslice[4:])
	fmt.Printf("Values between position 4 to 7:%v\n", myslice[4:7])
	slicelenght := len(myslice)          //Calculate the slice lenght
	last3valuesindice := slicelenght - 3 //Create a variable to store the indice of the first value of 3 last values
	fmt.Printf("Last 3 values:%v\n", myslice[last3valuesindice:slicelenght])

}
