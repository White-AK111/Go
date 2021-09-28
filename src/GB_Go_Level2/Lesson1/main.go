package main

import (
	"flag"
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson1/catchpanic"
	"github.com/White-AK111/GB_Go_Level2/Lesson1/millionfiles"
	"github.com/White-AK111/GB_Go_Level2/Lesson1/selfstudy"
)

func main() {
	// Task #1 and #2
	err := catchpanic.GetAndRecoverPanic()
	if err != nil {
		fmt.Println(err)
	}

	// Task #3
	var dir string
	flag.StringVar(&dir, "directory", "/home/white/MillionFiles", "Path for create 1M files")
	err = millionfiles.CreateMillionFiles(dir)
	if err != nil {
		fmt.Println(err)
	}

	// Task #4
	err = selfstudy.DontPanic()
	if err != nil {
		fmt.Println(err)
	}
}
