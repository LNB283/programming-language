/////////////////exo6.go///////////////////////////
///objective(s):
///				- Use function time to print out time in different format
///				- Check information about package time (https://golang.org/pkg/time/)
//////////////////////////////////////////////////
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("With time.now():\t%v\n", time.Now())
	fmt.Printf("With time.Kitchen:\t%v\n", time.Kitchen)
	fmt.Printf("With time.RFC850:\t%v\n", time.RFC850)
	fmt.Printf("With time.RFC3339:\t%v\n", time.RFC3339)
}
