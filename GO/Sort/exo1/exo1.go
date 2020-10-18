/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create 2 variable , type slice of integers and slice of strings
///				- Use the functions:
///					--sort.Ints([]int)
///					--sort.Strings([]strings)
///				- Print out the slices before and after Sort
//////////////////////////////////////////////////
package main

import (
	"fmt"
	"sort"
)

func main() {

	intslice := []int{23, 33, 45, 12, 18, 9, 24, 99}
	stringslice := []string{"Michael", "Rui", "Kobe", "James"}
	fmt.Printf("Slice of integer\n")
	fmt.Printf("Original:%v\n", intslice)
	sort.Ints(intslice)
	fmt.Printf("After Sort:%v\n", intslice)
	fmt.Printf("Slice of strings\n")
	fmt.Printf("Original:%v\n", stringslice)
	sort.Strings(stringslice)
	fmt.Printf("After Sort:%v\n", stringslice)

}
