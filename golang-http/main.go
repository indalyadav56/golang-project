package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println("server is running.... :8080")
	http.ListenAndServe(":8080", mux)
}
