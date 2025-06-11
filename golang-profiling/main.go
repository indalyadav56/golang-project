package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func main() {
	server := http.NewServeMux()

	// Your own handler
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Your own handler
	server.HandleFunc("/name", NameHandler)

	// Register pprof handlers manually
	server.HandleFunc("/debug/pprof/", pprof.Index)
	server.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	server.HandleFunc("/debug/pprof/profile", pprof.Profile)
	server.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	server.HandleFunc("/debug/pprof/trace", pprof.Trace)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", server)
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	for range 100_000_000 {
		fmt.Println("hello world")
	}
	w.Write([]byte(" name: John Doe"))
}
