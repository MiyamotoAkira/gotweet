package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	datastore "github.com/MiyamotoAkira/gotweet/datastore"
)

var users []string
var messages = make(map[string][]string)

func SetupRouter(repo *datastore.SQLiteRepository) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		userName := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"tweets": messages[userName]})
	})

	r.POST("/user/:name/tweet", func(c *gin.Context) {
		userName := c.Param("name")

		var json struct {
			Message string `json:"message" binding:"required"`
		}

		if c.Bind(&json) == nil {
			messages[userName] = append(messages[userName], json.Message)
			c.String(http.StatusCreated, "")
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	//r.GET("/user/:name", func(c *gin.Context) {
	r.POST("/user/register", func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
		}
		bindError := c.Bind(&json)
		if bindError == nil {
			userName := json.Name
			users = append(users, userName)
			messages[userName] = make([]string, 0)
			c.String(http.StatusCreated, "")
		}
	})

	return r
}
