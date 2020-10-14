/////////////////exo1.go//////////////////////////////////////////////
///objective(s):
///				- Print out all numbers betwwen 100 an 200 (both included)
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	for indice := 100; indice <= 200; indice++ {
		fmt.Println(indice)
	}

	// other method
	fmt.Printf("\n\n\n")
	var indice2 int
	indice2 = 201
	for indice2 <= 300 {
		fmt.Println(indice2)
		indice2++
	}

}
