/////////////////main.go/////////////////////////////////////////////////////////////////////////
///objective(s):
///				- Create 3 functions : square , sqaure root , Double
///				- In the template , use these functions
/////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

type val struct {
	Value int
}

func (v val) Double() int {
	return v.Value * 2
}

func (v val) SquareRoot() float64 {
	return math.Sqrt(float64(v.Value))
}

func (v val) Square() int {
	return v.Value * v.Value
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("methodtemplate.gohtml"))
}

func main() {

	v := val{
		Value: 56,
	}

	err := tpl.Execute(os.Stdout, v)
	if err != nil {
		log.Fatalln(err)
	}
}
