/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create 3 different Templates
///				- Use ParseGlob (https://golang.org/src/html/template/template.go?s=14740:14789#L442)
///					-- Parse the template in the forlder designated
///				- Use ExecuteTemplate (https://golang.org/src/html/template/template.go?s=4079:4164#L122)
///					-- Applies the template identified
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {

	//Parse all files on the folder Templates
	vartemplate, err := template.ParseGlob("Templates/*")
	//Manage potential fatal error
	if err != nil {
		log.Fatalln(err)
	}

	err = vartemplate.ExecuteTemplate(os.Stdout, "Template1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n")
	err = vartemplate.ExecuteTemplate(os.Stdout, "Template2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n")
	err = vartemplate.ExecuteTemplate(os.Stdout, "Template3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n")

}
