package main

import (
	"fmt"
)

//массивы, срезы, карты
func main() {

	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	total = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	total = 0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	z := [5]float64{
		100,
		101,
		102,
		103,
		104,
	}
	fmt.Println(z)

	var slice1 []float64
	fmt.Println(slice1)
	slice2 := make([]float64, 5)
	fmt.Println(slice2)
	slice3 := make([]float64, 5, 10)
	fmt.Println(slice3)
	arr := [5]float64{1, 2, 3, 4, 5}
	slice4 := arr[1:4]
	fmt.Println(slice4)
	fmt.Println(arr[1:])

	sliceAppend1 := []int{1, 2, 3}
	sliceAppend2 := append(sliceAppend1, 4, 5)
	fmt.Println(sliceAppend1, sliceAppend2)

	sliceCopy1 := []int{1, 2, 3}
	sliceCopy2 := make([]int, 2)
	copy(sliceCopy2, sliceCopy1)
	fmt.Println(sliceCopy1, sliceCopy2)

	mapSrtInt := make(map[string]int)
	mapSrtInt["key"] = 10
	fmt.Println(mapSrtInt["key"])

	mapIntStr := make(map[int]string)
	mapIntStr[1] = "one"
	fmt.Println(mapIntStr[1])
	delete(mapIntStr, 1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements2["B"])

	elements3 := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements3["O"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	//Как обратиться к четвертому элементу массива или среза?
	arrEx1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrEx1[3])
	sliceEx1 := arrEx1[3:4]
	fmt.Println(sliceEx1[0])

	//Чему равна длина среза, созданного таким способом: make([]int, 3, 9)?
	sliceEx2 := make([]int, 3, 9)
	fmt.Println(len(sliceEx2))

	//Дан массив: x := [6]string{"a","b","c","d","e","f"} что вернет вам x[2:5]?
	arrEx3 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(arrEx3[2:5])

	//Напишите программу, которая находит самый наименьший элемент в этом списке:

	arrEx4 := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var minElem int = arrEx4[0]
	for i := 1; i < len(arrEx4); i++ {
		if arrEx4[i] < minElem {
			minElem = arrEx4[i]
		}
	}
	fmt.Println(minElem)
}
