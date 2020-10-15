/////////////////exo5.go///////////////////////////
///objective(s):
///				- Create Slice of integer --> {1,2,3,4,5,6,7,8,9}
///				- Use range to browse the slice
///				- Print out the slice content
///				- Search the index value of 6
///				- Delete 6 from the original slice
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
	"sort"
)

func main() {

	originalslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original Slice\n%v\n", originalslice)
	fmt.Printf("position of value 6 is %v\n", sort.SearchInts(originalslice, 6))
	originalslice = append(originalslice[:5], originalslice[6:]...)
	fmt.Printf("New Slice after removal of value 6\n%v\n", originalslice)
}
