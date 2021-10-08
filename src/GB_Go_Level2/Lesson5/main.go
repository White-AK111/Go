package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/White-AK111/GB_Go_Level2/Lesson5/goroutines"
)

func main() {

	// Task#1 and Task#2
	var n int
	flag.IntVar(&n, "count of goroutines", 1000, "count of goroutines")

	calc := goroutines.NewGoroutinePool(n)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, err := calc.CalcSheep(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Total sheep: %d \n", res)
	fmt.Printf("Number of runnable goroutines: %d \n", runtime.NumGoroutine())
}
