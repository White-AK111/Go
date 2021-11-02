package fibonacci

import (
	"fmt"

	"gitlab.com/White-AK111/gb-go-level1/lesson5/helper"
)

//Print postition num in Fibonacci numbers
func PrintFibonacciRecursive(num int) {

	//Measure the execution time in nanoseconds
	defer helper.Duration(helper.Track("Duration Fibonacci recursive"))

	fmt.Printf("Fibonacci number %d = %d\n", num+1, FibonacciRecursive(num))
}

//Recursive calc Fibonacci number
func FibonacciRecursive(num int) int {

	switch {
	case num == 0:
		return 0
	case num < 2:
		return 1
	default:
		return FibonacciRecursive(num-2) + FibonacciRecursive(num-1)
	}
}

//Print postition num in Fibonacci numbers with map
func PrintFibonacciRecursiveWithMap(num int, fibonacciMap map[int]int) {

	//Measure the execution time in nanoseconds
	defer helper.Duration(helper.Track("Duration Fibonacci recursive with map"))

	fmt.Printf("Fibonacci number %d = %d\n", num+1, FibonacciRecursiveWithMap(num, fibonacciMap))
}

//Calc Fibonacci number with map
func FibonacciRecursiveWithMap(num int, fibonacciMap map[int]int) int {

	//If 2 previous number exist in map, calculate Fibonacci number from map
	_, ok := fibonacciMap[num-2]

	switch {
	case num == 0:
		fibonacciMap[num] = 0
	case num < 2:
		fibonacciMap[num] = 1
	case ok:
		fibonacciMap[num] = fibonacciMap[num-2] + fibonacciMap[num-1]
	default:
		fibonacciMap[num] = FibonacciRecursiveWithMap(num-2, fibonacciMap) + FibonacciRecursiveWithMap(num-1, fibonacciMap)
	}

	return fibonacciMap[num]
}
