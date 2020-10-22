/////////////////exo3.go///////////////////////////
///objective(s):
///				- Create a function sqrt
///					-- Parameter: float64
///					-- return float64 and error
///					-- If the value is less than 0 , print out an error and exit the program
//				- Without error, test value:=7.67
///				- With  error, test value:=-7.67
//////////////////////////////////////////////////
package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v", se.err)
}

func main() {
	value := 7.6
	//value := -7.67
	resultat, err := sqrt(value)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sqrt of %v is %v\n", value, resultat)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("Value passed: %v is not correct. We accept only positive float", f)
		return 0, sqrtError{e}
	}
	result := math.Sqrt(f)
	return result, nil
}
