package main

import (
	"database/sql"
	"errors"
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
func (u *User) getUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

//Function to Create User
func (u *User) createUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

// Function to Delete User
func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

// Update User
func (u *User) updateUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}
