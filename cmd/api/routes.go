package main //Package main is the entry point of the application. It contains the main function and the routes for the API.

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// routes function sets up the routes for the application using the Gin framework.
// It defines the API version and the endpoints for creating, retrieving, updating, and deleting events.
func (app *application) routes() http.Handler{
	g := gin.Default()

	// Set up CORS middleware
	v1 :=g.Group("/api/v1")
	{
		v1.POST("/events", app.createEvent) // Endpoint to create a new event
		v1.GET("/events", app.getAllEvents) // Endpoint to retrieve all events
		v1.GET("/events/:id", app.getEventById)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)


		v1.POST("/auth/register", app.registerUser) // Endpoint to register a new user


	}
    
	return g // Return the Gin router as the HTTP handler
}
