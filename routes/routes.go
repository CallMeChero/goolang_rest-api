package routes

import (
	middlewares "example/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// user
	server.POST("/signup", signup)
	server.POST("/login", login)
	// events

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// server.POST("events", middlewares.Authenticate, createEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("events", createEvent)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
}
