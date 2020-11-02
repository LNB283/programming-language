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
///					-- Under main, create 3 nbaplayer
///					-- Create a variable varslicenbaplayer , slice of structure
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
	vartemplate = template.Must(template.ParseFiles("slicestructtemplate.gohtml"))
}

func main() {

	MJ23 := nbaplayerdata{
		FirstName: "Michael",
		LastName:  "Jordan",
		Age:       57,
	}

	KB24 := nbaplayerdata{
		FirstName: "Kobe",
		LastName:  "Bryant",
		Age:       42,
	}

	nbaplayerslice := []nbaplayerdata{MJ23, KB24}

	err := vartemplate.Execute(os.Stdout, nbaplayerslice)
	if err != nil {
		log.Fatalln(err)
	}

}
