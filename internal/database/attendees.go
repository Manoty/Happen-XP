package database //Package database contains the models and database operations for the application.

import "database/sql" //Importing the database/sql package for SQL database operations.

// AttendeeModel struct represents the model for attendees in the database.
type AttendeeModel struct {
	DB * sql.DB 
}

type Attendee struct {
	Id int `json:"id"`
	UserId int `json:"userid"`
	EventId int `json:"event_id"`
}