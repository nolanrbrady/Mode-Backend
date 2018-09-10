package main

import (
	"database/sql"
)

// User model in DB
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Declare Crud operations for User model
// Function to Read Data
func (u *User) getUser(db *sql.DB, uid int) error {
	return db.QueryRow("SELECT name, id FROM users WHERE id=$1", u.ID).Scan(&u.Name, &u.ID)
}

//Function to Create User
func (u *User) createUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", u.Name, u.Email, u.ID)
	return err
}

// Function to Delete User
func (u *User) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM user WHERE id=$1", u.ID)
	return err
}

// Update User
func (u *User) updateUser(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", u.Name, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
