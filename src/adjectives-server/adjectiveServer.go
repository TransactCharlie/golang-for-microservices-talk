package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"log"
)

var adjectives = []string{
	"crazy",
	"steamin",
	"connivin",
	"ugly",
	"howlin",
}

func choose(s []string) string {
	return s[rand.Intn(len(s))]
}

func adjectivesHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an expensive operation.
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, choose(adjectives))
}

func main() {
	// Init the random seed
	rand.Seed(time.Now().Unix())
	log.Println("Starting Adjectives Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/adjective", adjectivesHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
