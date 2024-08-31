package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddService_ValidInput(t *testing.T) {
	data := url.Values{}
	data.Set("num1", "5")
	data.Set("num2", "3")

	req, err := http.NewRequest("POST", "/add", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	AddService(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	var response map[string]int
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode JSON response")

	expected := map[string]int{"result": 8}
	assert.Equal(t, expected["result"], response["result"], "Unexpected result in AddService")
}

func TestAddService_InvalidInput(t *testing.T) {
	data := url.Values{}
	data.Set("num1", "five")
	data.Set("num2", "3")

	req, err := http.NewRequest("POST", "/add", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	AddService(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status Bad Request")

	expectedBody := "Invalid input: 'Number 1' must be an integer\n"
	assert.Equal(t, expectedBody, rr.Body.String(), "Unexpected response body")
}
