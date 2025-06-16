package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := setupLogger()
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(zapMiddleware(logger)) // custom zap middleware

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Root handler called")
		w.Write([]byte("Hello, Zap Logger with Chi!"))
	})

	r.Delete("/users", func(w http.ResponseWriter, r *http.Request) {
	})

	// GET /users/{userID}
	r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		fmt.Fprintf(w, "User ID: %s\n", userID)
	})

	http.ListenAndServe(":8080", r)
}

func setupLogger() *zap.Logger {
	logFile, err := os.OpenFile("logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// WriteSyncers
	fileWriter := zapcore.AddSync(logFile)
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriter, zapcore.InfoLevel),
		zapcore.NewCore(encoder, consoleWriter, zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}

func zapMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Serve request
			next.ServeHTTP(w, r)

			duration := time.Since(start)
			logger.Info("HTTP Request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote", r.RemoteAddr),
				zap.Duration("duration", duration),
			)
		})
	}
}
