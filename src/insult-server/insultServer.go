package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)


func getResouce(target string) (string, error) {
	resp, err := http.Get(target)
	if err != nil {
		log.Printf("Error getting target: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error parsing body: %v", err)
		return "", err
	}

	return string(body), nil
}

func insultHandler(w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	start := time.Now()
	var action, adjective, principal string
	go func() {
		defer wg.Done()
		action, _ = getResouce("http://actions-server:8080/action")
	}()
	go func() {
		defer wg.Done()
		adjective, _ = getResouce("http://adjectives-server:8080/adjective")
	}()
	go func() {
		defer wg.Done()
		principal, _ = getResouce("http://principals-server:8080/principal")
	}()
	wg.Wait()
	fmt.Fprintf(w, fmt.Sprintf("%s ya %s %s\n",  action,  adjective, principal))
	log.Printf("Time Taken: %v", time.Since(start))
}

func main() {
	mux := http.NewServeMux()
	log.Println("Starting Insult Server")
	mux.HandleFunc("/insult", insultHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
