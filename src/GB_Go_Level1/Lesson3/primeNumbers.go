package main

import (
	"fmt"
	"math"
	"os"
)

//Prime numbers programm
func getPrimeNums() {
	var n int

	fmt.Print("Enter the upper border: ")
	_, err := fmt.Fscanln(os.Stdin, &n)
	if err != nil {
		fmt.Printf("%d is not a number!\n", n)
		os.Exit(0)
	}

	fmt.Printf("Simple numbers is: %v\n\n", atkin(n))
}

// Function returns list of all prime numbers less than n, Atkin sieve is using
//https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%88%D0%B5%D1%82%D0%BE_%D0%90%D1%82%D0%BA%D0%B8%D0%BD%D0%B0
func atkin(limit int) []int {
	// Инициализация решета
	x2, y2, n := 0, 0, 0
	sqrtLim := int(math.Sqrt(float64(limit)))
	arr := make([]bool, limit+1)
	arr[2], arr[3] = true, true

	// Предположительно простые — это целые с нечётным числом
	// представлений в данных квадратных формах.
	// x2 и y2 — это квадраты i и j (оптимизация).
	for i := 1; i <= sqrtLim; i++ {
		x2 = i * i
		for j := 1; j <= sqrtLim; j++ {
			y2 = j * j
			n = 4*x2 + y2

			if (n <= limit) && (n%12 == 1 || n%12 == 5) {
				arr[n] = !arr[n]
			}

			n -= x2
			if (n <= limit) && (n%12 == 7) {
				arr[n] = !arr[n]
			}

			n -= 2 * y2
			if (i > j) && (n <= limit) && (n%12 == 11) {
				arr[n] = !arr[n]
			}
		}
	}

	// Отсеиваем кратные квадратам простых чисел в интервале [5, sqrt(limit)].
	// (основной этап не может их отсеять)
	for i := 5; i <= sqrtLim; i++ {
		if arr[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				arr[j] = false
			}
		}
	}

	// Вывод списка простых чисел
	var primes []int
	for i, isPrime := range arr {
		if i > 1 && isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
