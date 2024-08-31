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

func TestSubtractService_ValidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "10")
	data.Set("b", "4")

	req, err := http.NewRequest("POST", "/subtract", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	SubtractService(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]int
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	expected := map[string]int{"result": 6}
	assert.Equal(t, expected["result"], response["result"])
}

func TestSubtractService_InvalidInput(t *testing.T) {
	data := url.Values{}
	data.Set("a", "ten")
	data.Set("b", "4")

	req, err := http.NewRequest("POST", "/subtract", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	SubtractService(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	expected := "Invalid input: 'a' must be an integer\n"
	assert.Equal(t, expected, rr.Body.String())
}
