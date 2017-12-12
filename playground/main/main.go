package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	work := make(chan int, 4)
	ctx, canc := context.WithCancel(context.Background())

	go func() {
		for {
			fmt.Println("loop")
			select {
			case work <- 1:
				fmt.Println("working")
			case <-ctx.Done():
				fmt.Println("done")
			}
		}
	}()

	time.Sleep(time.Millisecond)
	canc()
	time.Sleep(time.Millisecond)
}
