package database

import (
	"database/sql"
	"time"
	"context"
)

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

func (m *UserModel) Insert(user *User) error {
	//Insert the user into the database 
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel ()

	query := "INSERT INTO users (email, password, name) VALUES ($1, $2, 3$,)RETURNING id"

	return m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)


}
