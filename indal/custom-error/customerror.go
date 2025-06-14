package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// CustomError represents a custom error with additional context
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"-"` // Not serialized in JSON
}

// Error implements the error interface
func (e *CustomError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap allows unwrapping to get the underlying cause
func (e *CustomError) Unwrap() error {
	return e.Cause
}

// NewCustomError creates a new custom error
func NewCustomError(code int, message string, cause error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// Response is the structure for the API response
type Response struct {
	Message string `json:"message"`
}

// Handler is the interface for the HTTP handler
type Handler interface {
	Serve(w http.ResponseWriter, r *http.Request)
}

// RealHandler is the actual handler implementation
type RealHandler struct{}

func (h *RealHandler) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := NewCustomError(http.StatusMethodNotAllowed, "Method not allowed", nil)
		sendError(w, err)
		return
	}

	// Simulate a database operation that fails
	if err := simulateDatabaseOperation(); err != nil {
		customErr := NewCustomError(http.StatusInternalServerError, "Database operation failed", err)
		sendError(w, customErr)
		return
	}

	resp := Response{Message: "Hello from the Go HTTP server!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		customErr := NewCustomError(http.StatusInternalServerError, "Failed to encode response", err)
		sendError(w, customErr)
	}
}

// sendError writes the custom error to the response
func sendError(w http.ResponseWriter, customErr *CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(customErr.Code)
	json.NewEncoder(w).Encode(customErr)
}

// simulateDatabaseOperation simulates a database failure
func simulateDatabaseOperation() error {
	return errors.New("database connection timeout")
}

func main() {
	handler := &RealHandler{}
	http.HandleFunc("/hello", handler.Serve)
	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// type CustomError2 struct {
// 	ErrorCode string
// 	ErrorMsg  string
// 	Cause     error `json:"-"` // Not serialized in JSON
// }

// func (c *CustomError) Error2() string {
// 	return c.ErrorMsg
// }

// // Unwrap allows unwrapping to get the underlying cause
// func (e *CustomError) Unwrap() error {
// 	return e.Cause
// }

// type User struct {
// 	UserID    int    `json:"userId"`
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// func main() {
// 	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer resp.Body.Close()

// 	// if resp.StatusCode != http.StatusOK {
// 	// 	panic("Failed to fetch data: " + resp.Status)
// 	// }

// 	// byteData, err := io.ReadAll(resp.Body)

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Println(string(byteData))

// 	// var userObj User
// 	// if err := json.Unmarshal(byteData, &userObj); err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println(userObj.UserID)
// 	// fmt.Println(userObj.ID)
// 	// fmt.Println(userObj.Title)
// 	// fmt.Println(userObj.Completed)

// 	// fmt.Println(GetData().Unwrap())

// 	fmt.Println(ReverseString("indal"))

// }

// func GetData() *CustomError {
// 	return &CustomError{
// 		ErrorCode: "20",
// 		ErrorMsg:  "hello this is custom error",
// 		Cause:     errors.New("this is error cause for cusotm use"),
// 	}
// }

// func ReverseString(str string) string {
// 	r := []rune(str)
// 	left, right := 0, len(r)-1

// 	for left < right {

// 		r[left], r[right] = r[right], r[left]

// 		left++
// 		right++
// 	}

// 	return string(r)

// }
