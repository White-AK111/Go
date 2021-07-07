//Тестирование
package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {

	if len(xs) != 0 {
		total := float64(0)
		for _, x := range xs {
			total += x
		}
		return total / float64(len(xs))
	} else {
		return 0
	}

}

// Найти минимум в массиве чисел.
func Min(slice []float64) float64 {
	if len(slice) != 0 {
		min := slice[0]
		for i := 1; i < len(slice); i++ {
			if slice[i] < min {
				min = slice[i]
			}
		}
		return min
	} else {
		return 0
	}

}

// Найти максимум в массиве чисел.
func Max(slice []float64) float64 {
	if len(slice) != 0 {
		max := slice[0]
		for i := 1; i < len(slice); i++ {
			if slice[i] > max {
				max = slice[i]
			}
		}
		return max
	} else {
		return 0
	}
}
