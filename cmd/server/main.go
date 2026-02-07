package main

import (
	"fmt"
	"net/http"

	"go-alert-backend/internal/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method=%s path=%s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/alerts", handler.Alerts)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
