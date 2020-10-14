/////////////////exo3.go//////////////////////////////////////////////
///objective(s):
///				- Create a variable year
///					-- Type integer
///				- Assign a value
///				- Check if this year is a leap year or not
///					-- calcul the result of number%4 --> if equal to 0  --> leap year
///					-- calcul the result of number%100 --> if equal to 0  --> leap year
///					-- calcul the result of number%400 --> if equal to 0  --> leap year
///				- Print out the result
///
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var year int
	year = 2032
	if year%4 == 0 || year%100 == 0 || year%400 == 0 {
		fmt.Printf("Year %d is a leap year", year)
	} else {
		fmt.Printf("Year %d is not a leap year", year)
	}
}
