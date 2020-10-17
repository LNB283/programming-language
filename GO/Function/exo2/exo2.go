/////////////////exo2.go///////////////////////////
///objective(s):
///				- Create a function : addition 2 integer
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func addition(a int, b int) int {

	result := a + b
	return result
}

func subtraction(a int, b int) int {
	result := a - b
	return result
}

func multiplication(a int, b int) int {
	result := a * b
	return result
}

func division(a int, b int) int {
	result := a / b
	return result
}

func main() {

	inta := 2
	intb := 3

	resultat := addition(inta, intb)
	fmt.Printf("%d + %d = %d\n", inta, intb, resultat)
	resultat = subtraction(intb, inta)
	fmt.Printf("%d - %d = %d\n", intb, inta, resultat)
	resultat = multiplication(inta, intb)
	fmt.Printf("%d * %d = %d\n", intb, inta, resultat)
	resultat = division(intb, inta)
	fmt.Printf("%v / %v = %v\n", intb, inta, resultat)

}
