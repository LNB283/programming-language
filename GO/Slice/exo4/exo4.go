/////////////////exo4.go///////////////////////////
///objective(s):
///				- Create Slice of integer --> {1,2,3,4,5,6,7,8,9}
///				- Use range to browse the slice
///				- Print out the slice content
///				- Add 1 value to this slice --> {10}
///				- Print out the slice content
///				- Add 3 new values --> {11,12,13}
///				- Print out the slice content
///				- Add a new slice --> {14,15,16}
///				- Print out the slice content
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	originalslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original Slice\n%v\n", originalslice)
	originalslice = append(originalslice, 10)
	fmt.Printf("Add value 10 to the original slice\n%v\n", originalslice)
	originalslice = append(originalslice, 11, 12, 13)
	fmt.Printf("Add value 14,15,16 to the original slice\n%v\n", originalslice)
	addslice := []int{14, 15, 16}
	originalslice = append(originalslice, addslice...)
	fmt.Printf("Add value from new slice componed by value 14,15,16 to the original slice\n%v\n", originalslice)

}
