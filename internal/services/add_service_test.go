package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddService_ValidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "5")
	data.Set("b", "3")

	req, err := http.NewRequest("POST", "/add", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	AddService(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AddService returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]int
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	expected := map[string]int{"result": 8} // 5 + 3 = 8
	if response["result"] != expected["result"] {
		t.Errorf("AddService returned unexpected result: got %v want %v", response["result"], expected["result"])
	}
}

func TestAddService_InvalidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "five")
	data.Set("b", "3")

	req, err := http.NewRequest("POST", "/add", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	AddService(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("AddService returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "Invalid input: 'a' must be an integer\n"
	if rr.Body.String() != expected {
		t.Errorf("AddService returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
