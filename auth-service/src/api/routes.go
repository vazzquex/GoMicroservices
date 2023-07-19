package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() http.Handler {

	router := gin.Default()

	// specify who is allowed to connect
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Link"}
	config.AllowCredentials = true
	config.MaxAge = 300
	router.Use(cors.New(config))

	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/ping" && c.Request.Method == "GET" {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
			c.Abort()
			return
		}
		c.Next()
	})

	router.POST("/auth", func(c *gin.Context) {
		app.Auth()
	})

	return router

}
