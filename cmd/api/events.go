package main //Package main is the entry point of the application. It contains the main function and the routes for the API.

import (
	"net/http" //Importing the HTTP package for handling HTTP status codes and responses.
	"strconv"

	"github.com/Manoty/Happen-XP/internal/database" //Importing the database package for database operations.

	"github.com/gin-gonic/gin" //Importing the Gin framework for building the web application.
)

// createEvent handles the creation of a new event.
func (app *application) createEvent(c *gin.Context){
	// Bind the incoming JSON request body to the Event struct.
	// If there's an error during binding, respond with a 400 Bad Request status and the error message.
	var event database.Event
	if err := c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//attempt to insert the event into the database using the Insert method from the models package.
	err :=app.models.Events.Insert(&event)

	// If there's an error during insertion, respond with a 500 Internal Server Error status and an error message.
	// If the insertion is successful, respond with a 201 Created status and the created event.

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusCreated, event)
	


}
// getAllEvents retrieves all events from the database and responds with a JSON array of events.
func (app *application) getAllEvents(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id")) // Convert the ID parameter from string to int.

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
	}
	event, err := app.models.Events.Get(id)//Retrieve the event from the database using the Get method from the models package.

    if event == nil{
		c.JSON(http.StatusNotExtended, gin.H{"error": "Event not found"})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
	
	}
	c.JSON(http.StatusOK, event)                                                            
}