package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"log"
)

var principals = []string{
	"bampot",
	"bawbag",
	"clipe",
	"dobber",
	"jakey",
}

func choose(s []string) string {
	return s[rand.Intn(len(s))]
}

func principalHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an expensive operation.
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, choose(principals))
}

func main() {
	// Init the random seed
	rand.Seed(time.Now().Unix())
	log.Println("Starting Principals Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/principal", principalHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
