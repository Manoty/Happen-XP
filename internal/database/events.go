package database

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Manoty/Happen-XP/internal/database"
	"github.com/gin-gonic/gin"
)

// AttendeeModel struct represents the model for attendees in the database.
type EventModel struct {
	DB *sql.DB

}
// Insert method inserts a new event into the database.
type Event struct {
	Id int `json:"id"`
	OwnerId int `json:"ownner_id" binding:"required"`
	Name string `json:"name" binding:"required,min=5"`
	Description string `json:"description" binding:"required,min=15"`
	Date string `json:"date" binding:"required, datetime=2006-01-02T15:04:05Z07:00"`
	Location string `json:"location" binding:"required, min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5)"
	

	return m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)

}
func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
		
	}
	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		// Handle any errors encountered during iteration over rows
		return nil, err
	}
	return events, nil
}
func (m *EventModel) Get(id int) (*Event, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel ()

	query := "SELECT * FROM events WHERE id = $1"

	var event Event

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
	if err != nil{
		if err == sql.ErrNoRows {
			return nil, nil // Event not found
		}
		return nil, err // Other error

	}
	return &event, nil // Event found
}
func (m *EventModel) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query:= "UPDATE events SET name = $1, description - $2, date = $3, Location = $4 WHERE id = $5"

	_, err := m.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Location, event.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query:= "DELETE FROM events WHERE id = $1"
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (app *application)addAttendeesToEvent(c *gin.Context){
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "	Invalid event Id"})
		return
	} 

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "	Invalid user Id"})
		return
	}
	event, err := app.models.Events.Get(eventId)
	if err := nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}
	userToAdd, err := app.models.Users.Get(userId)
	if err := nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve User"})
		return
	}
	if userToAdd == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	existingAttendee, err := app.models.attendees.GetByEventAndAttendee(eventId, userToAdd)
	if err := nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attendee"})
		return
	}
	if existingAttendee == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "attendee already exists"})
		return
	}
	attendee := database.Attendee{
		EventId: event.Id, 
		userId: userToAdd.Id,
	}

	_err = app.models.Attendees.insert(&attendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add attendee"})
		return
	}
	c.JSON(http.StatusCreated, attendee)

}
func (app *application) getAttendeesForEvent (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event Id"})
		return

	}
	users, err := app.models.Attendeer.GetAttendeesByEvent(id)
	if err := nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve event attendee"})
		return
	}
	c.JSON(http.StatusOK, users)

}
func (app *application) deleteAttendeeFromEvent (c *gin.Context ){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event Id"})
		return 

	}
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return

	}
	err := app.models.Attendees.Delete(userId, id )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attendee "})
		return
	}
	c.JSON(http.StatusNoContent, nil)
	
}
func (app *application) getEventByAttendee(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendee id"})
		return
	}
	event, err := app.models.Events.GetByAttendee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get event"})
		return
	}
	c.JSON(http.StatusOk, events)
}

