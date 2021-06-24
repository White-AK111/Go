package main

import (
	"fmt"
)

func main() {
	var x string = "Hello World!"
	fmt.Println(x)

	var y string = "Hello"
	y += " World!"
	fmt.Println(y)

	z := "first"
	w := "first"
	fmt.Println(z == w)

	const constVar int = 111
	fmt.Println(constVar)

	fmt.Print("Input a temperature in F: ")
	var tempF float64
	fmt.Scanf("%f", &tempF)
	tempC := ((tempF - 32) * 5 / 9)
	fmt.Println(tempC)

	fmt.Print("Input length in fut: ")
	var lenF float64
	fmt.Scanf("%f", &lenF)
	lenM := (lenF * 0.3048)
	fmt.Println(lenM)
}
