package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemp(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sensing", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var d SensingData
	err := json.Unmarshal([]byte(w.Body.String()), &d)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
}

func TestReadDevice(t *testing.T) {
	_, err := readDevice()
	assert.Nil(t, err)
}
