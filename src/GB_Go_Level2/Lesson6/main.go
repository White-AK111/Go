package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson6/manscheduler"
	"github.com/White-AK111/GB_Go_Level2/Lesson6/race"
	"github.com/White-AK111/GB_Go_Level2/Lesson6/semaphores"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

	var n int
	flag.IntVar(&n, "count of goroutines", 100, "count of goroutines")

	// Task#1
	sem := semaphores.NewGoroutinePool(n)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	_, err := sem.Squaring(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	sem.PrintSquaring()
	fmt.Printf("Number of runnable goroutines: %d \n", runtime.NumGoroutine())

	// Task#2
	man := manscheduler.NewGoroutinePool(n)
	ctxMan, cancelFuncMan := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncMan()
	_, err = man.Inc(ctxMan)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Number of runnable goroutines: %d \n", runtime.NumGoroutine())

	// Task#3
	rc := race.NewGoroutinePool(n)
	ctxRc, cancelFuncMan := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncMan()
	_, err = rc.IncRace(ctxRc)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Number of runnable goroutines: %d \n", runtime.NumGoroutine())
}
