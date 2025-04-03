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

		v1.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)
		v1.GET("/events/:id/attendees", app.getAttendeesForEvent)

		v1.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent) // Endpoint to delete attendees from event)
		v1.GET("/attendees/:id/events", app.getEventByAttendee)

		
		v1.POST("/auth/register", app.registerUser) // Endpoint to register a new user


	}
    
	return g // Return the Gin router as the HTTP handler
}
