package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	fmt.Println("Started")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Print(i, ",")
			// Note compiler removes this...
			time.Sleep(2)
		}(i)
	}
	wg.Wait()
	fmt.Println("\nFinished")
}
