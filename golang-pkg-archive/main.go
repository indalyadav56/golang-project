package main

import (
	"fmt"
	"net/http"
)

// Replace with the actual module path

func main() {
	mux := http.NewServeMux()

	fmt.Println("server is runiing on port 8080")
	http.ListenAndServe(":8080", mux)
}
