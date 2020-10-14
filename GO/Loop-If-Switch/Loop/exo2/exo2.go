/////////////////exo2.go//////////////////////////////////////////////
///objective(s):
///				- find numbers which are divisible by 7 and multiple of 5, between 100 and 200 (both are included)
///				- Print out the result
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	for indice := 100; indice <= 200; indice++ {
		if indice%7 == 0 && indice%5 == 0 {
			fmt.Printf("number %v is divided by 7 and 5\n", indice)
		}
	}

}
