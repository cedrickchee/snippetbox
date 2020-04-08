package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is custom error if database record is not found.
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials is custom error if a user tries to login with an
	// incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail is custom error if a user tries to signup with an email
	// address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet is ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User is ...
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
