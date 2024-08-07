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

// As of now this interface is not so much useful.
// But maybe it will be useful in the future
type UserCreator interface {
	Create(email, password string) (*User, error)
	Authenticate(email, password string) (*User, error)
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

func (us *UserStore) Authenticate(email, password string) (*User, error) {

	email = strings.ToLower(email)
	user := User{
		Email: email,
	}

	row := us.DB.QueryRow(`
			SELECT id, password_hash 
			FROM users WHERE 
			email = $1`, email)

	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("authenticate user: %w", err)
	}
	return &user, nil

}
