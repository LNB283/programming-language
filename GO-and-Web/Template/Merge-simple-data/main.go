/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create 2 files : main.go
///				- main.go
///					-- create a variable firstname and assign a value
///					-- create a variable vartemplate
///						--- This variable contains the skeleton of your htnl webpage
///					-- Print out the result and create a file index.html
///						--- Command line to create the index.html : go run main.go > index.html
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	firstname := "Rui"

	vartemplate := `
	<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>NBA Player data</title>
		</head>
		<body>
		<h1>First Name: ` + firstname + `</h1>
		</body>
		</html>
	`
	fmt.Println(vartemplate)
}
