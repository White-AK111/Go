package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gitlab.com/White-AK111/gb-go-level1/lesson4/sortfunction"
)

func main() {
	//Suppose the initial capacity is 10
	unsortedSlice := make([]int64, 0, 10)

	//Read from Stdin by buffered I/O bufio pkg
	fmt.Printf("For send EOF enter:\t%s\n", "Ctrl + D")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		} else {
			unsortedSlice = append(unsortedSlice, num)
		}
	}

	fmt.Printf("Unsorted slice:\t%v\n", unsortedSlice)
	fmt.Printf("Sorted slice:\t%v\n", sortfunction.InsertionSort(&unsortedSlice))
}
