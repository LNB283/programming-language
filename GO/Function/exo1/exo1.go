/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create the fibonacci function
///				- Call it and print out the result
//////////////////////////////////////////////////

package main

import (
	"fmt"
)

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	now := 0
	next := 1
	return func() int {

		result := now
		//fmt.Println(result)
		temp := next
		next = next + now
		//fmt.Println(next)
		now = temp
		//fmt.Println(now)
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		//fmt.Printf("\n")
		//fmt.Printf("Loop %d\n", i)
		fmt.Printf("Fibonacci %v = %v\n", i, f())
		//fmt.Println(f())
	}
}
