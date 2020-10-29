/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Modify the main.go from Marge-simple-data-v2
///				- pass manually the first name and last name on the command
///					-- From terminal:  go run main.go "Rui" "Hachimura"
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//Give the first name
	firstname := os.Args[1]
	//Give the Last name
	lastname := os.Args[2]
	//Give the age of the player
	ageplayer := os.Args[3]

	//Program executes
	fmt.Println(os.Args[0])
	//The additional input
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	vartemplate := fmt.Sprintf(`
	<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>NBA Player data</title>
		</head>
		<body>
		<h1>First Name: ` + firstname + `</h1>
		<h1>Last Name: ` + lastname + `</h1>
		<h1>Age: ` + ageplayer + `</h1>
		</body>
		</html>`)

	//Create  and open the index.html file
	reqcreate, err := os.Create("index.html")
	//Manage potential fatal error
	if err != nil {
		log.Fatal("Create an Error file", err)
	}
	// Copy the content from vartemplate to reqcreate
	io.Copy(reqcreate, strings.NewReader(vartemplate))
	//Don't forget to close the file after the creation process is completed
	defer reqcreate.Close()

}
