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
	if err := c.ShouldBindJSON(&event); err != nil {
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
func (app *application) getEventById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id")) // Convert the ID parameter from string to int.

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := app.models.Events.Get(id)//Retrieve the event from the database using the Get method from the models package.

    if event == nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	
	}
	c.JSON(http.StatusOK, event)                                                            
}

func (app *application) getAllEvents(c *gin.Context){
	event, err := app.models.Events.GetAll() //Retrieve all events from the database using the GetAll method from the models package.

	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive events"})
		return
	}
	c.JSON(http.StatusOK, event) //Respond with a 200 OK status and the list of events in JSON format.
}

func (app *application) updateEvent(c *gin.Context) {
	// Get the event ID from the request URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	// Fetch the existing event from the database
	existingEvent, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	// If the event does not exist, return 404 Not Found
	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Bind the incoming JSON to updatedEvent
	updatedEvent := &database.Event{}
	if err := c.ShouldBindJSON(updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the ID in the updatedEvent matches the path parameter
	updatedEvent.Id = id

	// Update the event in the database
	if err := app.models.Events.Update(updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	// Return the updated event
	c.JSON(http.StatusOK, updatedEvent)
}
func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert the ID parameter from string to int.
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})

	}
	if err := app.models.Events.Delete(id); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
	}
	c.JSON(http.StatusNoContent, nil)
}
func (app *qpplication) addAttendeeToEvent(c *gin.Context) {
	eventId, err :- strconv.Atoi(c.Param("id")) // Convert the event ID parameter from string to int.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userId, err :- strconv.Atoi(c.Param("userid")) // Convert the event ID parameter from string to int.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	event err := app.models.Events.Get(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":  "Failed to retrieve event"})
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}
	userToAdd, err := app.models.Users.Get(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve User"})
		return
	}
	if userToAdd == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	//make sure user is not already an attendee
	existingAttendee, err := app.models.Attendees.GetByEventAndUser(event.Id, userToAdd.Id)
	if existingAttendee  != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "attendee already exists"})
		return
	}
	//create the attendee
	attendee := database.Attendee{
		EventId: event.Id,
		UserId: userToAdd.Id,
	}
	_, err = app.models.Attendees.Insert(&attendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add attendee"})
		return
	}
	c.JSON(http.StatusCreated, attendee) // Respond with a 201 Created status and the created attendee.
}
func (app *application) getAttendeesForEvent(c *gin.Context){
	
}