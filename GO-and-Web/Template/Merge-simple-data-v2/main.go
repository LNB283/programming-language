/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Modify the main.go from Marge-simple-data
///				- Create automaticaally the index.html
///				- Manage potential Fatal error
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

	firstname := "Rui"
	lastname := "Hachimura"

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
