package main

import (
	"context"
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson4/goroutines"
	"github.com/White-AK111/GB_Go_Level2/Lesson4/signalchanel"
	"log"
	"runtime"
	"time"
)

func main() {

	// Task#1
	calc := goroutines.NewGoroutineAdd(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, err := calc.AddWithGoroutines(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Last value in chanell: %d \n", res)
	fmt.Printf("Number of runnable goroutines: %d \n", runtime.NumGoroutine())

	//Task#2
	fmt.Println(signalchanel.TreatmentSignal())
}
