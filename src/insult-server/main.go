package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var principals = []string{
	"bampot",
	"bawbag",
	"clipe",
	"dobber",
	"jakey",
}

var adjectives = []string{
	"crazy",
	"steamin",
	"connivin",
	"ugly",
	"howlin",
}

var actions = []string{
	"Awa' n boil yer head",
	"Shut yer gob",
	"Yer aff yer heid",
}

func choose(s []string) string {
	return s[rand.Intn(len(s))]
}

func genInsult() string {
	time.Sleep(time.Second)
	return fmt.Sprintf("%s ya %s %s\n",  choose(actions),  choose(adjectives), choose(principals))
}

func insultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, genInsult())
}

func main() {
	// Init the random seed
	rand.Seed(time.Now().Unix())
	mux := http.NewServeMux()
	mux.HandleFunc("/insult", insultHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
