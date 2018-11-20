package main

import (
	"fmt"
	"time"
	"context"
)

func worker(ctx context.Context, msg string) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
		fmt.Println(msg)
	}
}

func main() {
	d := time.Now().Add(time.Second * 2)
	bg := context.Background()
	ctx, _ := context.WithDeadline(bg, d)

	worker(ctx, "Hello")
	time.Sleep(time.Second * 5)
	worker(ctx, "Hello Again!")
}
