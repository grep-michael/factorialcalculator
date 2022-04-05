package main

import (
	"fmt"
	"log"
	"net/http"
)

func factorialRequestion(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "die nerd\n")
}

func startServer() {
	mux := http.NewServeMux()
	mux.Handle("/factorial", http.HandlerFunc(factorialRequestion))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
