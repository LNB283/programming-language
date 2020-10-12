/////////////////exo2.go///////////////////////////
///objective(s):
///				- create 6 variables:
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- Don't assign a value
///				- Print the type and the result:
///							-- with Println
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	//Variable declaration and values assignement
	var lastname string
	var firstname string
	var age int
	var height int
	var weight int
	var hometown string
	var country string
	var righthand bool
	var lefthand bool

	// Print Type of each variables with Printlf
	//Use Printf and %T
	fmt.Printf("Type of each variables\n")
	fmt.Printf("Last Name:%T\nFirst Name:%T\nAge:%T\nHeight:%T\nWeight:%T\nHomeTown:%T\nCountry:%T\nRight Hand:%T\nLeft Hand:%T\n", lastname, firstname, age, height, weight, hometown, country, righthand, lefthand)
	fmt.Printf("\n")
	// Print result with Println
	//By default, if we don't assign any values, for string you have blank , int you have 0 , boolean you have false
	fmt.Println("Print result without vvalues assigned")
	fmt.Println(" Last Name: ", lastname, "\n", "First Name: ", firstname, "\n", "Age: ", age, "\n", "Height: ", height, "\n", "Weight: ", weight, "\n", "HomeTown: ", hometown, "\n", "Country: ", country, "\n", "Right Hand: ", righthand, "\n", "Left Hand: ", lefthand)

}
