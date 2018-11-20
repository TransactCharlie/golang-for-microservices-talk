package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 4)

	for i := 0; i < 100; i++ {
		go func(i int) {
			messages <- i
			fmt.Println(i)
		}(i)
	}

	sum := 0
	for i := 0; i < 100; i++ {
		sum += <-messages
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}
