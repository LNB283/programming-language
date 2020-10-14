/////////////////exo2.go//////////////////////////////////////////////
///Problem statement:
///						A circular pool 24m away from the conference.
///						Calculate an approximate value of the radius
///objective(s):
///				- Mathematical formula : 2 * Pi * r
///				- Define Pi as a constant (value --> 3.14159265359)
///					-- Create a constant valpi
///					-- Untyped
///				- Create a variable radius
///					-- Type float32
///				- Create a variable Result
///					-- Type float32
///				- Print out the result
//////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

const (
	// valpi is untyped
	valpi = 3.14159265359
)

func main() {
	//Variable declaration
	var radius float32
	var resultat float32
	//Assign a value to r
	radius = 24
	//Calculation
	resultat = 2 * radius * valpi
	//Print out the result
	fmt.Printf("A circular pool 24m away from the conference.\n")
	fmt.Printf("Approximate value of the Radius:%fm", resultat)

}
