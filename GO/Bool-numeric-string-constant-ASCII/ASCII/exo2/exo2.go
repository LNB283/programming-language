/////////////////exo2.go///////////////////////////
///objective(s):
///				- Print the character corresponding to the ASCII value between 33 to 126
///				- Ideally by using a loop
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	for indice := 33; indice < 127; indice++ {
		charactvalueresult := rune(indice)
		fmt.Printf("%d --> %c\n", indice, charactvalueresult)
	}

}
