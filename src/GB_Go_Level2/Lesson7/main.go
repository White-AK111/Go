package main

import (
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson7/assignreflect"
	"github.com/White-AK111/GB_Go_Level2/Lesson7/astparser"
	"log"
)

type MyInt int

func main() {
	// Task#1
	m := map[string]interface{}{
		"StructField": struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}{IntFieldInStruct: 100500, InterfaceFieldInStruct: nil},
		"SliceIntField": []int{1, 3, 5, 7},
		"IntField":      111,
		"FloatField":    222.222,
		"StringField":   "it's a string",
		"BoolField":     true,
	}

	in := assignreflect.NewIn()
	err := assignreflect.AssignmentByReflect(in, m)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Struct in is: %v\n", in)

	// Task#2
	n, err := astparser.CountGoroutines("astparser/astparser.go", "ExecGoroutines")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Count of go func execute: %d\n", n)

	// Task#3
	var one MyInt = 1
	q := NewMyIntQueue()
	q.Insert(one)
}
