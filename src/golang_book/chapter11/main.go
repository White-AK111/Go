//Пакеты и повторное использование кода
//Тестирование
package main

import (
	"fmt"
	m "src/golang_book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)

	xMinMax := []float64{11.11, 22.22, 8.88, 12.12}
	min := m.Min(xMinMax)
	fmt.Println(min)

	max := m.Max(xMinMax)
	fmt.Println(max)
}
