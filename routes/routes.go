package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []string

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tweets": make([]string, 0)})
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
		user := c.Params.ByName("name")
		users = append(users, user)
		c.String(http.StatusCreated, "")
	})

	return r
}
