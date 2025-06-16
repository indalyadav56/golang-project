package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
}

func main() {

}

func handler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
