package main

import (
	"indal/handlers"
	"indal/utils"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize database (assuming utils.InitDB exists)
	utils.InitDB()

	// Set up Chi router
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user created"))
	})

	router.Post("/create-user", handlers.CreateUserHandler)

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user created"))
	})

	// Create a spinner for a dynamic "starting" effect
	spinner, _ := pterm.DefaultSpinner.
		WithShowTimer(true).
		WithText("Starting the server...").
		Start()

	// Simulate some startup work (optional)
	// time.Sleep(2 * time.Second) // Uncomment to see the spinner in action

	// Server configuration
	serverAddr := ":8080"

	// Start the server in a goroutine to keep the spinner alive
	go func() {
		if err := http.ListenAndServe(serverAddr, router); err != nil {
			spinner.Fail("Failed to start server: ", err)
			return
		}
	}()

	// Wait briefly to ensure the server starts
	// time.Sleep(100 * time.Millisecond)

	// Success message with styled box
	spinner.Success("Server Started Successfully!")
	pterm.DefaultBox.
		WithTitle("🚀 API Server").
		WithTitleTopCenter().
		WithTextStyle(pterm.NewStyle(pterm.FgGreen)).
		Println(pterm.Sprintf("Listening on http://localhost%s\nReady to handle requests!", serverAddr))
}

// func main() {
// 	utils.InitDB()

// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)

// 	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("user created"))
// 	})

// 	router.Post("/create-user", handlers.CreateUserHandler)

// 	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("user created"))
// 	})

// 	http.ListenAndServe(":8080", router)
// }

func Sum(a, b int) int {
	return a + b
}
