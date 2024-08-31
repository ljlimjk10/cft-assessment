package services

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func SubtractService(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	aStr := r.FormValue("a")
	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Invalid input: 'a' must be an integer", http.StatusBadRequest)
		return
	}

	bStr := r.FormValue("b")
	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Invalid input: 'b' must be an integer", http.StatusBadRequest)
		return
	}

	result := a - b

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"result": result})
}
