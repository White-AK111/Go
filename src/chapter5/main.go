package main

import (
	"fmt"
)

func main() {
	count := 10

	for i := 1; i <= count; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	for i := 1; i <= count; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "even")

		default:
			fmt.Println(i, "odd")
		}
	}

	count = 100
	for i := 1; i <= count; i++ {
		if i%3 == 0 {
			fmt.Println(i, "div by 3")
		}
	}

	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
