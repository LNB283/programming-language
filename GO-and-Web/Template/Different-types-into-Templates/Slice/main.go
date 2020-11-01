/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Pass a slice into template
///					-- Create a variable vartemplate , type : pointer of template
///					-- Create a function inittemplate
///						--- Use template.Must and template.ParseFiles
///					-- Under main, create a variable player, type slice of string
///				- Create a template to be able to accept slice of string
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"os"
	"text/template"
)

var vartemplate *template.Template

func init() {
	vartemplate = template.Must(template.ParseFiles("slicetemplate.gohtml"))
}

func main() {

	nbaplayer := []string{"Jordan", "Kobe", "Rui", "Tony"}

	err := vartemplate.Execute(os.Stdout, nbaplayer)
	if err != nil {
		log.Fatalln(err)
	}
}
