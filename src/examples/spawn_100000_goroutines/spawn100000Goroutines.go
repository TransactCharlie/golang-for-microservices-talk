package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting 100 000 goroutines!")
	for i := 0; i < 100000; i++ {
		go func(i int) {
			fmt.Print(i, ", ")
			time.Sleep(10)
		}(i)
	}
	fmt.Println("Started")
	time.Sleep(time.Second * 10)
	fmt.Println("Finished")
}
