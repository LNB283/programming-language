/////////////////exo2.go//////////////////////////////////////////////
///objective(s):
///				- Create a variable numuber
///					-- Type integer
///				- Assign a value
///				- Check if the number is odd or even
///					-- calcul the result of number%2
///					-- result = 0 --> even
///					-- result != 0 --> odd
///				- Print out the result
///
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var number, result int

	number = 13
	result = number % 2
	if result == 0 {
		fmt.Printf("the number %d is an even", number)
	} else {
		fmt.Printf("the number %d is an odd", number)
	}

}
