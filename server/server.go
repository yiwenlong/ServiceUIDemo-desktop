package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

var servers map[int]*http.Server

func Boot(address string) int {
	if servers == nil {
		servers = make(map[int]*http.Server, 5)
	}
	srv := http.Server{Addr: address}
	log.Printf("Start listening on %s\n", address)

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	go func() {
		defer func() {
			log.Printf("Server finish")
		}()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	servers[0] = &srv
	return 0
}

func Stop(server int) {
	if srv, ok := servers[server]; ok {
		log.Printf("Stop server.")
		srv.Shutdown(context.TODO())
		log.Printf("Server shut down!!")
		return
	}
	log.Fatalf("No server found.")
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
