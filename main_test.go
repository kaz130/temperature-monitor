package main

import (
	"testing"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestGetTemp(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sensing", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	_, exists := response["temp"]
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
}
