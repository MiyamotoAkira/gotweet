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

func TestRegisterUser(t *testing.T) {
	router := sut.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"name":"vader""}`)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestGivenNoTweetsProfileShouldReturnEmpty(t *testing.T) {
	router := sut.SetupRouter()
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"name":"vader""}`)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/vader", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"tweets\":[]}", w.Body.String())
}
