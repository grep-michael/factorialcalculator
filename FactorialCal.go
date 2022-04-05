package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	LIMIT = 65
)

func factorialRequestion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Text/html; charset=utf-8")
	fmt.Fprintf(w, "Calculate all factorials from 0-65, with and without memoization<br><br>")
	/*i, e := strconv.ParseUint(req.URL.Query().Get("n"), 10, 64)
	if e != nil {
		fmt.Println(e)
		fmt.Fprintf(w, "Supply get param n for factorial")
		return
	}*/
	runStockFac(w)
	fmt.Fprintf(w, "<br><br>")
	runMemoizedFac(w)

}

var cache = make(map[uint64]uint64)

func memoizedfac(i uint64) uint64 {
	if cache[i] != 0 {
		return cache[i]
	}
	if i <= 0 {
		return 1
	}
	return i * stockfac(i-1)
}

func runMemoizedFac(w http.ResponseWriter) {
	fmt.Fprintf(w, "Memoized recursive factorial<br>")
	start := time.Now()
	for i := uint64(0); i < LIMIT; i++ {
		factorial := memoizedfac(i)
		cache[i] = factorial
	}
	elasped := time.Since(start)
	fmt.Fprintf(w, "Operation time: %v<br>", elasped)
}

func stockfac(i uint64) uint64 {
	if i <= 0 {
		return 1
	}
	return i * stockfac(i-1)
}

func runStockFac(w http.ResponseWriter) {
	fmt.Fprintf(w, "Stock recursive factorial<br>")
	start := time.Now()
	for i := uint64(0); i < LIMIT; i++ {
		stockfac(i)
	}
	elasped := time.Since(start)
	//fmt.Println(factorial)
	fmt.Fprintf(w, "Operation time: %v<br>", elasped)

}

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/factorial", http.HandlerFunc(factorialRequestion))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
