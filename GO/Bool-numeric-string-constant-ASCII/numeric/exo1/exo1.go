/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a variable test
///				- Type of test is integer
///				- assign a value to test
///				- Print out the result in 3 forms: decimal / binary / hexadecimal (ref: https://godoc.org/fmt)
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var test int
	test = 156
	///Print out the value of test in
	///Decimal : %d
	///Binary  : %b
	///Hexadecimal: %x
	fmt.Printf("Value of test\nDecimal:%d\nBinary:%b\nHexadecimal:%x\n", test, test, test)
}
