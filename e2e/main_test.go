package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	datastore "github.com/MiyamotoAkira/gotweet/datastore"
	sut "github.com/MiyamotoAkira/gotweet/routes"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	repo *datastore.SQLiteRepository
}

func (suite *MainTestSuite) SetupTest() {
	suite.repo = datastore.Setup("/tmp/tweet.db")
}

func (suite *MainTestSuite) TestBasicFlow() {
	router := sut.SetupRouter(suite.repo)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("pong", w.Body.String())
}

func (suite *MainTestSuite) TestRegisterUser() {
	router := sut.SetupRouter(suite.repo)
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"name":"vader"}`)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	suite.Equal(201, w.Code)
	suite.Equal("", w.Body.String())
}

func (suite *MainTestSuite) TestGivenNoTweetsProfileShouldReturnEmpty() {
	router := sut.SetupRouter(suite.repo)
	w := httptest.NewRecorder()

	var jsonStr = []byte(`{"name":"vader"}`)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/vader", nil)
	router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("{\"tweets\":[]}", w.Body.String())
}

func (suite *MainTestSuite) TestGivenSentATweetProfileShouldReturnTweet() {
	router := sut.SetupRouter(suite.repo)
	w := httptest.NewRecorder()

	jsonStr := []byte(`{"name":"vader"}`)
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	jsonStr = []byte(`{"message":"I am your subtweeter"}`)
	req, _ = http.NewRequest("POST", "/user/vader/tweet", bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/vader", nil)
	router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal(`{"tweets":["I am your subtweeter"]}`, w.Body.String())
}
