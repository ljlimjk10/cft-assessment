package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSubtractService_ValidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "10")
	data.Set("b", "4")

	req, err := http.NewRequest("POST", "/subtract", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	SubtractService(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("SubtractService returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]int
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	expected := map[string]int{"result": 6}
	if response["result"] != expected["result"] {
		t.Errorf("SubtractService returned unexpected result: got %v want %v", response["result"], expected["result"])
	}
}

func TestSubtractService_InvalidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "ten")
	data.Set("b", "4")

	req, err := http.NewRequest("POST", "/subtract", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	SubtractService(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("SubtractService returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "Invalid input: 'a' must be an integer\n"
	if rr.Body.String() != expected {
		t.Errorf("SubtractService returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
