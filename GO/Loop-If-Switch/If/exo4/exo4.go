/////////////////exo3.go//////////////////////////////////////////////
///objective(s):
///				- Create 3 variables
///					-- Type integer
///				- Assign a value for each
///				- check if the trianagle  is Equilateral, Isosceles or Scalene
///				- Print out the result
///Additional info:
///				equilateral: three sides equal
///				isocele: two sides equal
///				scalene : all sides are different
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var side1, side2, side3 int
	side1, side2, side3 = 3, 3, 3

	if side1 == side2 && side2 == side3 {
		fmt.Printf("equilateral")
	} else if side1 == side2 || side1 == side3 || side2 == side3 {
		fmt.Printf("isocele")
	} else {
		fmt.Printf("scalene")
	}

}
