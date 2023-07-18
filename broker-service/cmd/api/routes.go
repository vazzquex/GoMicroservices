package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	router := gin.Default()

	//specify who is allowed to connect
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"https://*", "http://*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Tocken"}
	config.ExposeHeaders = []string{"Link"}
	config.MaxAge = 300

	router.Use(cors.New(config))

	router.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
