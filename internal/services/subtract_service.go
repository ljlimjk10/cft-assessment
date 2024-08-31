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

	numOneStr := r.FormValue("num1")
	num1, err := strconv.Atoi(numOneStr)
	if err != nil {
		http.Error(w, "Invalid input: 'Number 1' must be an integer", http.StatusBadRequest)
		return
	}

	numTwoStr := r.FormValue("num2")
	num2, err := strconv.Atoi(numTwoStr)
	if err != nil {
		http.Error(w, "Invalid input: 'Number 2' must be an integer", http.StatusBadRequest)
		return
	}

	result := num1 - num2

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"result": result})
}
