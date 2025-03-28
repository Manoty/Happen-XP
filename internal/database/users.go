package database

import "database/sql"

// UserModel struct represents the model for users in the database.
type UserModel struct {
	DB *sql.DB

}
// Insert method inserts a new user into the database.
type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"-"`
	Name string `json:"name"`
	
}
