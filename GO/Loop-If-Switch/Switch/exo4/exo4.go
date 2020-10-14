/////////////////exo4.go//////////////////////////////////////////////
///objective(s):
///				- print the logic gate by using switch
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	//Indicate the logic gate from this list : AND | OR | NOT
	LogicGate := `NOT`

	switch LogicGate {
	case `AND`:
		fmt.Printf("Logical gate AND\n")
		fmt.Printf("Input1\t\tInput2\t\tOutput\n")
		fmt.Printf("0\t\t0\t\t0\n")
		fmt.Printf("0\t\t1\t\t0\n")
		fmt.Printf("1\t\t0\t\t0\n")
		fmt.Printf("1\t\t1\t\t1\n")
	case `OR`:
		fmt.Printf("Logical gate OR\n")
		fmt.Printf("Input1\t\tInput2\t\tOutput\n")
		fmt.Printf("0\t\t0\t\t0\n")
		fmt.Printf("0\t\t1\t\t1\n")
		fmt.Printf("1\t\t0\t\t1\n")
		fmt.Printf("1\t\t1\t\t1\n")
	case `NOT`:
		fmt.Printf("Logical gate NOT\n")
		fmt.Printf("Input1\t\tOutput\n")
		fmt.Printf("0\t\t1\n")
		fmt.Printf("1\t\t0\n")

	}

}
