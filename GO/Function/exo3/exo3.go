/////////////////exo3.go///////////////////////////
///objective(s):
///				- Create a function :  average
///					--parameters : unlimited integers
///					-- return : float32
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func avg(val ...int) (float32, int, int) {

	size := len(val)
	sum := 0
	for _, value := range val {
		sum = sum + value
	}
	//fmt.Printf("Sum of the serie: %d\n", sum)
	//fmt.Printf("Serie lenght: %d\n", size)
	average := float32(sum) / float32(size)
	return average, sum, size
}

func main() {

	serie, sum, size := avg(1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 10, 1)

	fmt.Printf("Sum of the serie: %d\n", sum)
	fmt.Printf("Serie lenght: %d\n", size)
	fmt.Printf("Average is %d / %d = %f\n", sum, size, serie)

}
