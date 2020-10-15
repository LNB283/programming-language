/////////////////exo6.go///////////////////////////
///objective(s):
///				- Create 2 slices of strings
///				- Print out the 2 slices contents
///				- create a multi dimensional slice of strings that contains the 2 previous slices
///				- Print out the multi dimensional slice of strings content
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	//Create 2 slices of strings with values
	slice1 := []string{`Rui`, `Hachimura`, `Japan`}
	slice2 := []string{`Kobey`, `Bryant`, `USA`}
	//Print out the raw information to validate 2 slices
	fmt.Println(slice1)
	fmt.Println(slice2)
	//Create a multi-dimensional slice of string
	finalslice := [][]string{slice1, slice2}
	//Print out the raw information
	fmt.Println(finalslice)
	//First loop to collect the index position of each records
	for indice, intermediateslice := range finalslice {
		fmt.Printf("Index %d\n", indice)
		//Second loop to browse the content of each records and print out
		for indice2, values := range intermediateslice {
			fmt.Printf("Index position : %d \t Values: %s\n", indice2, values)
		}
	}
}
