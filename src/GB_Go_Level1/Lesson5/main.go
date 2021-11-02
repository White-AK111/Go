package main

import (
	"fmt"
	"os"

	"gitlab.com/White-AK111/gb-go-level1/lesson5/fibonacci"
)

func main() {

	//0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, …

	//Read num
	num := 0
	fmt.Print("Enter number for calculate Fibonacci position:")
	_, err := fmt.Scanln(&num)

	//Chech num for Int
	if err != nil {
		fmt.Printf("%d is not a number!\n", num)
		os.Exit(1)
	}

	//With recurcive
	fmt.Printf("\nFibonacci recursive\n")
	fibonacci.PrintFibonacciRecursive(num - 1)

	//With recurcive and map
	//Suppose the initial map is 10
	fmt.Printf("\nFibonacci recursive with map\n")
	fibonacciMap := make(map[int]int, 10)
	fibonacci.PrintFibonacciRecursiveWithMap(num-1, fibonacciMap)
}
