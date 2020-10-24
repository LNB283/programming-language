/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create Array with 10 values (Type integer)
///				- Assign a value for each index
///				- Use range to browse the array
///				- Print out the array
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	//Define my array
	myarray := [10]int{23, 8, 33, 18, 12, 56, 24, 99, 7, 22}
	fmt.Printf("indice\t\tval\n")
	//Browse the array
	for indice, val := range myarray {
		fmt.Printf("%d\t\t%d\n", indice, val)
	}
	//Print the type of my array
	fmt.Printf("Type of My Array is:%T", myarray)
}
