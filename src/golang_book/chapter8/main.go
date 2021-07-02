package main

import (
	"fmt"
)

//Указатели

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	//z := *x
	//*x = *y
	//*y = z

	*x, *y = *y, *x
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	//Какое будет значение у переменной x после выполнения программы:
	y := 1.5
	square(&y)
	fmt.Println(y)

	//Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1).
	xs, ys := 1, 2
	swap(&xs, &ys)
	fmt.Println(xs, ys)
}
