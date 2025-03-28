package database

import "database/sql"
// UserModel struct represents the model for users in the database.
type Models struct {
	Users UserModel
	Events EventModel
	Attendees AttendeeModel

}
// User struct represents the model for users in the database.
func NewModels(db *sql.DB) Models{
	return Models{
		Users : UserModel{DB : db},
		Events: EventModel{DB : db},
		Attendees: AttendeeModel{DB: db},
		
	}
}