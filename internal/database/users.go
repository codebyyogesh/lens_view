package database

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserCreator interface {
	Create(email, password string) (*User, error)
}
type UserStore struct {
	DB *sql.DB
}

func (us *UserStore) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	// write user info to DB
	row := us.DB.QueryRow(`
	INSERT INTO users (email, password_hash) 
	VALUES ($1, $2) RETURNING id`, email, hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: string(hashedBytes),
	}
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}
