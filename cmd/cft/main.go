package main

import (
	"fmt"
	"net/http"

	"cft/internal/services"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		services.AddService(w, r)
	}).Methods("POST", "OPTIONS")

	router.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		services.SubtractService(w, r)
	}).Methods("POST", "OPTIONS")

	http.Handle("/", enableCORS(router))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
