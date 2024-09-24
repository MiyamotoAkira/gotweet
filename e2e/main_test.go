package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	sut "github.com/MiyamotoAkira/gotweet/routes"
	"github.com/stretchr/testify/assert"
)

func TestBasicFlow(t *testing.T) {
	router := sut.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBasicPost(t *testing.T) {
	router := sut.SetupRouter()

	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"value":"weee"}`)
	req, _ := http.NewRequest("POST", "/admin", bytes.NewBuffer(jsonStr))

	req.Header.Add("authorization", "Basic Zm9vOmJhcg==")
	req.Header.Add("content-type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"ok\"}", w.Body.String())
}
