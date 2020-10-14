/////////////////exo3.go//////////////////////////////////////////////
///objective(s):
///				- Use only boolean value
///				- 1 == 1 / 1 == 2 / 2 == 2 / 3 == 4
///				- Print out the result
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	switch {
	case (1 == 1):
		fmt.Println("1 == 1 : True")
		fallthrough //use fallthrough to continue to test next case
	case (2 == 2):
		fmt.Println("2 == 2 : True")
	case (3 == 4):
		fmt.Println("3 == 4 : False")
	}

}
