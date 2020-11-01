/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Pass a structure into template
///					-- Create a variable vartemplate , type : pointer of template
///					-- Create a structure nbaplayerdata
///						--- First Name as a string
///						--- Last Name as a string
///						--- Age as an int
///					-- Create a function inittemplate
///						--- Use template.Must and template.ParseFiles
///					-- Under main, create 2 or 3 variable type nbaplayerdata and assign values
///				- Create a template to be able to accept slice of string
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"log"
	"os"
)

var vartemplate *template.Template

type nbaplayerdata struct {
	FirstName string
	LastName  string
	Age       int
}

func init() {
	vartemplate = template.Must(template.ParseFiles("structtemplate.gohtml"))
}

func main() {

	nbaplay1 := nbaplayerdata{
		FirstName: "Michael",
		LastName:  "Jordan",
		Age:       57,
	}

	nbaplay2 := nbaplayerdata{
		FirstName: "Kobe",
		LastName:  "Bryant",
		Age:       42,
	}

	err := vartemplate.Execute(os.Stdout, nbaplay1)
	if err != nil {
		log.Fatalln(err)
	}

	err = vartemplate.Execute(os.Stdout, nbaplay2)
	if err != nil {
		log.Fatalln(err)
	}

}
