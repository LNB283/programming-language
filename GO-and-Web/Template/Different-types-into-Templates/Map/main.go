/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Pass a structure into template
///					-- Create a variable vartemplate , type : pointer of template
///					-- Create a function inittemplate
///						--- Use template.Must and template.ParseFiles
///					-- Under main, create a vatiable nbaplayer type Map of strings
///				- Create a template to be able to accept slice of string
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"html/template"
	"log"
	"os"
)

var vartemplate *template.Template

func init() {
	vartemplate = template.Must(template.ParseFiles("maptemplate.gohtml"))
}

func main() {

	nbaplayerdata := map[string]string{
		"Jordan":    "Michael",
		"Bryant":    "Kobe",
		"Hachimura": "Rui",
	}

	err := vartemplate.Execute(os.Stdout, nbaplayerdata)

	if err != nil {
		log.Fatalln(err)
	}

}
