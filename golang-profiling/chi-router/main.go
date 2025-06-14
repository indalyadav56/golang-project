package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users   = make(map[int]User)
	nextID  = 1
	userMux sync.Mutex
)

func main() {
	// Register pprof routes
	registerPprof()

	// CRUD Routes
	http.HandleFunc("/users", handleUsers)     // POST, GET
	http.HandleFunc("/users/", handleUserByID) // GET, PUT, DELETE

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Create or Get All
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userMux.Lock()
		defer userMux.Unlock()
		var list []User
		for _, user := range users {
			list = append(list, user)
		}
		json.NewEncoder(w).Encode(list)

	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		userMux.Lock()
		user.ID = nextID
		nextID++
		users[user.ID] = user
		userMux.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Get, Update, Delete by ID
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userMux.Lock()
	defer userMux.Unlock()

	user, exists := users[id]
	if !exists {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(user)

	case "PUT":
		var updated User
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		updated.ID = id
		users[id] = updated
		json.NewEncoder(w).Encode(updated)

	case "DELETE":
		delete(users, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 🔬 Register pprof endpoints
func registerPprof() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		log.Println("pprof available at :6060")
		log.Println(http.ListenAndServe("localhost:6060", mux))
	}()
}

func testCRUDLoad() {
	client := &http.Client{}

	// 🔸 Create user (POST /users)
	user := map[string]string{"name": "Test User", "email": "test@example.com"}
	body, _ := json.Marshal(user)
	resp, err := client.Post("http://localhost:8080/users", "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println("POST error:", err)
		return
	}
	defer resp.Body.Close()

	var createdUser map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&createdUser)
	id := int(createdUser["id"].(float64))
	fmt.Println("Created user ID:", id)

	// 🔸 Get all users (GET /users)
	resp, err = client.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("GET all error:", err)
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// 🔸 Get single user (GET /users/{id})
	resp, err = client.Get(fmt.Sprintf("http://localhost:8080/users/%d", id))
	if err != nil {
		fmt.Println("GET single error:", err)
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// 🔸 Update user (PUT /users/{id})
	updated := map[string]string{"name": "Updated User", "email": "updated@example.com"}
	body, _ = json.Marshal(updated)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/users/%d", id), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("PUT error:", err)
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// 🔸 Delete user (DELETE /users/{id})
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/users/%d", id), nil)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("DELETE error:", err)
		return
	}
	resp.Body.Close()
}
