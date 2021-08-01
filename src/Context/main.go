package main

import (
	"context"
	"fmt"
	"time"
)

func do(ctx context.Context, id int) {
	fmt.Printf("%d done\n", id)
}

func main() {
	for id := 0; id < 1000; id++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		do(ctx, id)
	}
}
