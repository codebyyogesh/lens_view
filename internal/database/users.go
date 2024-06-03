package database

import "database/sql"

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserCreator interface {
	Create(name, email string) (*User, error)
}
type UserStore struct {
	DB *sql.DB
}
