/////////////////exo1.go//////////////////////////////////////////////
///objective(s):
///				- create a variable rank
///					-- Type string
///				- test the rank value for , at least, 4 cases (included default case)
///				- Print out the result
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	rank := ' '
	//test the rank value for each case
	switch rank {
	case 'A':
		fmt.Println("Ranking A")
	case 'B':
		fmt.Println("Ranking B")
	case 'C':
		fmt.Println("Ranking C")
	default:
		fmt.Println("No Ranking")

	}

}
