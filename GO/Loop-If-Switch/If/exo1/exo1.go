/////////////////exo1.go//////////////////////////////////////////////
///objective(s):
///				- Create 3 variables
///					-- Type integer
///				- Assign 3 values
///				- Check if 2 numbers are equal or not
///				- Print out the result
///
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var number1, number2, number3 int
	var eval bool
	number1 = 23
	number2 = 33
	number3 = 23

	fmt.Printf("number1:%d\tnumber2:%d\tnumber3:%d\n", number1, number2, number3)
	//Use if
	if number1 == number2 {
		fmt.Printf("number1(%d) = number2(%d)", number1, number2)
		eval = true
	}
	if number2 == number3 {
		fmt.Printf("number2(%d) = number3(%d)", number2, number3)
		eval = true
	}
	if number1 == number3 {
		fmt.Printf("number1(%d) = number3(%d)", number1, number3)
		eval = true
	}
	if eval == false {
		fmt.Printf("No combination exist where 2 numbers are equal")
	}

	fmt.Printf("\n")
	//Use if ... else if ... else
	if number1 == number2 {
		fmt.Printf("number1(%d) = number2(%d)", number1, number2)
	} else if number2 == number3 {
		fmt.Printf("number2(%d) = number3(%d)", number2, number3)
	} else if number1 == number3 {
		fmt.Printf("number1(%d) = number3(%d)", number1, number3)
	} else {
		fmt.Printf("No combination exist where 2 numbers are equal")
	}
}
