package main

import (
	"fmt"
	"os"
)

func main() {
	for stop := false; !stop; {
		switch getTask() {
		case 0:
			fmt.Println("Goodbye!")
			stop = true
		case 1:
			execCalc()
		case 2:
			getPrimeNums()
		default:
			fmt.Println("There is no such task, goodbye!")
			stop = true
		}
	}
}

//Get current task by user
func getTask() int {
	var taskId int
	fmt.Printf(" %s\n %s\n %s\nSelect task: ",
		"1 - Calculator",
		"2 - SimpleNumbers",
		"0 - Exit")

	_, err := fmt.Scan(&taskId)
	//Chech num for Int
	if err != nil {
		fmt.Printf("%v is not a number!\n", taskId)
		os.Exit(0)
	}

	return taskId
}
