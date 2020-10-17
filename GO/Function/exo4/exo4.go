package main

import (
	"fmt"
)

var cel float32
var far float32

func convertceltofar(celsius float32) float32 {

	fahrenheit := (9*celsius + (32 * 5)) / 5
	return fahrenheit

}

func convertfartocel(fahrenheit float32) float32 {

	celsius := (5 * (fahrenheit - 32)) / 9
	return celsius

}

func main() {

	cel = -20
	far = 68
	resultconvertctof := convertceltofar(cel)
	fmt.Printf("%f celsius is equal to %f farenheit\n", cel, resultconvertctof)
	resultconvertftoc := convertfartocel(far)
	fmt.Printf("%f fahrnheit is equal to %f celsius\n", far, resultconvertftoc)

}
