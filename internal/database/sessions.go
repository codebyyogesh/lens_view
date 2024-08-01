package database

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/codebyyogesh/lens_view/internal/rand"
)

type Session struct {
	ID        int
	UserID    int
	TokenHash string
}

type SessionStore struct {
	DB *sql.DB
	//SessionBytesPerToken sets the length of generated session tokens. If unset or too small, the minimum length (MinSessionBytesPerToken) is used.
	SessionBytesPerToken int
}

type NewSession struct {
	NEWSession Session
	// Token is only set when creating a new session. When looking up a session this will be left empty, as we only store the hash of a session token in our database and we cannot reverse it into a raw token.
	Token string
}

const (
	// min bytes for each session token
	MinSessionBytesPerToken = 32
)

//  We will both create the session token and hash it inside Create, then we could return the original session token from the Create method. Create will create a new session for the user provided. The session token will be returned as the Token field on the NEWSession type, but only the hashed session token is stored in the database.

func (ss *SessionStore) Create(userID int) (*NewSession, error) {
	// TODO: Create the session token
	sessionBytesPerToken := ss.SessionBytesPerToken
	if sessionBytesPerToken < MinSessionBytesPerToken {
		sessionBytesPerToken = MinSessionBytesPerToken
	}
	token, err := rand.String(sessionBytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := NewSession{
		NEWSession: Session{
			UserID:    userID,
			TokenHash: ss.hash(token),
		},
		Token: token,
	}
	// Avoid duplicate sessions getting created for the same userID as this ID is unique. E.g. Same user trying to login from difference devices. To solve this we presume that session exists for every user and then try to update it. If it doesn't exist we will create it. Benefit of this approach is we avoid querying the db frequently.

	row := ss.DB.QueryRow(` 
		UPDATE sessions 
		SET token_hash = $2 
		WHERE user_id = $1 
		RETURNING id;`, session.NEWSession.UserID, session.NEWSession.TokenHash)
	// get existing session id
	err = row.Scan(&session.NEWSession.ID)
	if err == sql.ErrNoRows {
		// session does not exist, create it
		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash) 
		VALUES ($1, $2) RETURNING id`, session.NEWSession.UserID, session.NEWSession.TokenHash)
		// get the new session ID, error here will be overwritten by a new error or nil
		err = row.Scan(&session.NEWSession.ID)
	}
	// if the error was not sql.ErrNoRows or other error due to new session creation, return the error
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil
}

// Once a session is created we will need a way to query our SessionStore to determine who the user is with that session.
func (ss *SessionStore) UserLookup(token string) (*User, error) {
	// 1. Hash the session token
	tokenHash := ss.hash(token)

	// 2. Use the hash to Query the db for the session
	var user User
	row := ss.DB.QueryRow(`
			SELECT user_id 
			FROM sessions 
			WHERE token_hash = $1`, tokenHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("user lookup: %w", err)
	}
	// 3. Use the userID from the session to query the db for the user
	row = ss.DB.QueryRow(`
			SELECT  email, password_hash
			FROM users WHERE id = $1`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user lookup: %w", err)
	}
	// 4. Return the user
	return &user, nil
}

// Delete the session associated with the provided token
func (ss *SessionStore) Delete(token string) error {

	tokenHash := ss.hash(token)
	// Execute is used instead of QueryRow because we dont care for any data returned by the query
	_, err := ss.DB.Exec(`
		DELETE FROM sessions 
		WHERE token_hash = $1`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionStore) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:]) //[:] needed because tokenHash is an array and not slice, [:] converts it into a slice
}
