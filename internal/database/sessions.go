package database

import (
	"database/sql"
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
			UserID: userID,
		},
		Token: token,
	}
	// TODO: Hash the session token
	// TODO: Implement SessionService.Create
	return &session, nil
}

// Once a session is created we will need a way to query our SessionStore to determine who the user is with that session.
func (ss *SessionStore) UserLookup(token string) (*User, error) {
	// TODO: Implement SessionService.UserLookup
	return nil, nil
}
