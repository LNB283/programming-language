/////////////////exo1.go//////////////////////////////////////////////
///objective(s):
///				- Collect information concerning the environment by using:
///					--runtime package (https://godoc.org/runtime / https://golang.org/pkg/runtime/)
///				- Print out the result
///
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("Informations concerning the environment\n")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Go Routine\t", runtime.NumGoroutine())
	fmt.Printf("\n")
}
