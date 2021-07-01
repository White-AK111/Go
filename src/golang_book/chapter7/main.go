package main

import (
	"fmt"
	"os"
)

//Функции

//функция среднего
func average(xs []float64) float64 {
	//panic("Not Implemented")

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//функция возврата двух int
func f() (int, int) {
	return 5, 6
}

//функция суммирования
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//функция генерации чётных чисел
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//функция рекурсии
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

//функция вывода 1st
func first() {
	fmt.Println("1st")
}

//функция вывода 2nd
func second() {
	fmt.Println("2nd")
}

//Функция sum принимает срез чисел и складывает их вместе. Как бы выглядела сигнатура этой функции?
func sumSlice(slice []int) int {
	total := 0
	for _, num := range slice {
		total += num
	}
	return total
}

//Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае, если исходное число чётное, и false, если нечетное. Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).
func half(num int) (byte, bool) {
	if num%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

//Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.
func variableParams(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

//Используя в качестве примера функцию makeEvenGenerator напишите makeOddGenerator, генерирующую нечётные числа.
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Напишите рекурсивную функцию, находящую fib(n).
func fibo(n int) int {
	if n < 2 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x + y)

	fmt.Println(add(1, 2, 3))

	xsl := []int{1, 2, 3}
	fmt.Println(add(xsl...))

	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	fmt.Println(factorial(5))

	defer second()
	first()

	filename := `C:\Users\forma\OneDrive\Документы\Go\src\golang_book\chapter6\main.go`
	f, _ := os.Open(filename)
	defer f.Close()

	//задание1
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(slice1))

	//задание2
	fmt.Println(half(0))
	fmt.Println(half(4))
	fmt.Println(half(5))
	fmt.Println(half(10))

	//задание3
	fmt.Println(variableParams(1, 2, 3, 4, 5, 4, 8, 2))

	//задание4
	nextEven2 := makeOddGenerator()
	fmt.Println(nextEven2()) // 1
	fmt.Println(nextEven2()) // 3
	fmt.Println(nextEven2()) // 5

	//задание5
	fmt.Println(fibo(22))

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
