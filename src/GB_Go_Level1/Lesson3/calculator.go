package main

import (
	"fmt"
	"math"
	"os"
)

//Calculator programm
func execCalc() {
	var operand1, operand2, result float64
	var operation string
	operations := []string{"+", "-", "*", "/", "mod", "sqrt", "%"}

	fmt.Print("Enter 1-st operand: ")
	_, err1 := fmt.Scanln(&operand1)
	//Chech num for float
	if err1 != nil {
		fmt.Printf("Incorrect first operand %v entered!\n\n", operand1)
		os.Exit(0)
	}

	fmt.Print("Enter 2-st operand: ")
	_, err2 := fmt.Scanln(&operand2)
	//Chech num for float
	if err2 != nil {
		fmt.Printf("Incorrect second operand %v entered!\n\n", operand2)
		os.Exit(0)
	}

	fmt.Printf("Enter operation %v", operations)
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 != 0 {
			result = operand1 / operand2
		} else {
			fmt.Printf("Division by zero unsupported!\n\n")
		}
	case "mod":
		result = float64(int(operand1) % int(operand2))
	case "sqrt":
		result = math.Sqrt(operand1)
	case "%":
		result = operand1 * operand2 / 100
	default:
		fmt.Printf("Operation %s does not support!\n\n", operation)
	}

	fmt.Printf("Result is: %v\n\n", result)
}
