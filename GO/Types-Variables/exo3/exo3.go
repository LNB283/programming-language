/////////////////exo3.go///////////////////////////
///objective(s):
///				- create 6 variables:
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- Create a variable "result"
///				- Use the function sprintf to assign the result to the variable "result"
///				- Print the result:
///							-- with Println
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	//Variable declaration
	var lastname string
	var firstname string
	var age int
	var height int
	var weight int
	var hometown string
	var country string
	var righthand bool
	var lefthand bool
	var result string //variable is a string. We assign the result of Sprintf function

	righthand = true
	//Assign the result of Sprintf function to the variable result
	result = fmt.Sprintf("Last Name:%s\nFirst Name:%s\nAge:%d\nHeight:%d\nWeight:%d\nHomeTown:%s\nCountry:%s\nRight Hand:%t\nLeft Hand:%t\n", lastname, firstname, age, height, weight, hometown, country, righthand, lefthand)
	//Print the result by using the variable result
	fmt.Println(result)
}
