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

	fmt.Println("Environment architecture:", runtime.GOARCH)
	fmt.Println("Environment OS:", runtime.GOOS)

}
