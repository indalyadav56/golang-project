package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w)
	})

	http.ListenAndServe(":8080", server)

}
