package main

import (
	"fmt"
	"time"
)

func sleepThenHello() {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello")
}

func main() {
	fmt.Println("Going hello")
	go sleepThenHello()
	fmt.Println("Sleeping 10")
	time.Sleep(time.Second * 10)
	fmt.Println("Finished")
}
