package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Going Hello")
	msg := "Hello Friends!"
	go func(msg string) {
		time.Sleep(time.Second * 5)
		fmt.Println(msg)
	}(msg)
	msg = "Bye Friends!"
	fmt.Println(msg)
	fmt.Println("Sleeping 10")
	time.Sleep(time.Second * 10)
	fmt.Println("Finished")
}
