package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"log"
)

var actions = []string{
	"Awa' n boil yer head",
	"Shut yer gob",
	"Yer aff yer heid",
}

func choose(s []string) string {
	return s[rand.Intn(len(s))]
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an expensive operation.
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, choose(actions))
}

func main() {
	// Init the random seed
	rand.Seed(time.Now().Unix())
	log.Println("Starting Actions Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/action", actionHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
