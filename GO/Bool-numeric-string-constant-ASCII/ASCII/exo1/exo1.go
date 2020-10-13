/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create 4 variables
///					-- charactvalue
///					-- asciivalueresult
///					-- asciivalue
///					-- charactvalueresult
///				- assign a value to charactvalue and asciivalue
///				- convert charactvalue to ASCII and assign the result to asciivalueresult
///				- convert asciivalue to charact and assign the result to charactvalueresult
///				- Print out the result
///					-- asciivalueresult
///					-- charactvalueresult
//////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	charactvalue := 'L'
	asciivalue := 72

	asciivalueresult := int(charactvalue)
	fmt.Printf("Value of %c in ASCII is:%d\n", charactvalue, asciivalueresult)
	charactvalueresult := rune(asciivalue)
	fmt.Printf("In ASCII, value %d --> %c\n", asciivalue, charactvalueresult)
}
