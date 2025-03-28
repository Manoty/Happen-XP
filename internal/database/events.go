package database

import "database/sql"
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