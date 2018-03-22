package main

import (
	"flag"
	"sync"
	"context"
	"golang.org/x/sync/semaphore"
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

var concurrentConnections int
var numRequests int

func init() {
	flag.IntVar(&concurrentConnections, "c", 1, "Max Concurrent Connections")
	flag.IntVar(&numRequests, "nr", 10, "Total calls To make")
	flag.Parse()
}

func getInsult() string {
	resp, err := http.Get("http://localhost:8080/insult")
	if err != nil {
		log.Printf("Error getting insult: %v", err)
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error parsing body: %v", err)
		return err.Error()
	}

	return string(body)
}

func main() {

	ctx := context.TODO()

	var (
		sem = semaphore.NewWeighted(int64(concurrentConnections))
		out = make([]string, numRequests)
		workerGroup sync.WaitGroup
	)

	for i := range out {
		// blocks if we can't acquire samaphore.
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		fmt.Printf("%v", i)
		// Go get an insult
		workerGroup.Add(1)
		go func(i int) {
			defer sem.Release(1)
			defer workerGroup.Done()
			out[i] = getInsult()
		}(i)
	}

	// Wait for all insult requests to finish
	workerGroup.Wait()
	fmt.Print("\n")
	// Show the Insults
	for i := range(out) {
		fmt.Println(out[i])
	}
}