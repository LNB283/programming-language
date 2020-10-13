/////////////////exo4.go///////////////////////////
///objective(s):
///				- create 6 variables and assign a value for each:
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- create your own type (outside the main function)
///							-- type : jerseynumber
///							-- jerseynumber is an integer
///				- create a variable jnum by using the type jerseynumber
///				- assign a value
///				- Print the Type of jerseynumber
///				- Print the value of jnumb
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

type jerseynumber int

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
	// create jnumb variable with the new type jerseynumber
	var jnum jerseynumber
	//Variable declaration and values assignement
	lastname = "Hachimura"
	firstname = "Rui"
	age = 22
	height = 203
	weight = 104
	hometown = "Toyama"
	country = "Japan"
	righthand = true
	// Assign a value to jnum
	jnum = 8
	//Print the result
	fmt.Printf("Last Name:%s\nFirst Name:%s\nAge:%d\nHeight:%d\nWeight:%d\nHomeTown:%s\nCountry:%s\nRight Hand:%t\nLeft Hand:%t\nJersey Number:%d\n", lastname, firstname, age, height, weight, hometown, country, righthand, lefthand, jnum)
	fmt.Printf("\n")
	//Print out the type of the variable jnumb
	fmt.Printf("Type of Jersey Number is:%T\n", jnum)

}
