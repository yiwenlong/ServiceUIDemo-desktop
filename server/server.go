package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func Serv(address string) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Printf("Start listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Printf("URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	log.Printf("Count %d\n", count)
}
