/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create Slice of xx integer
///				- Assign a value for each
///				- Use range to browse the slice
///				- Print out the slice content
//////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	//Define my slice. You don't indicate the size for the slice
	myslice := []int{23, 8, 33, 18, 12, 56, 24, 99, 7, 9, 12, 45}
	fmt.Printf("index\t\tval\n")
	//Browse the slice
	for indice, val := range myslice {
		fmt.Printf("%d\t\t%d\n", indice, val)
	}
	//Print the type of my array
	fmt.Printf("Type of My slice is:%T\n", myslice)
	fmt.Printf("The length of our slice is:%v\n", len(myslice))
	fmt.Printf("The capacity of our slice is:%v\n", cap(myslice))

}
