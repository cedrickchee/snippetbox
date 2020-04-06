package models

import (
	"errors"
	"time"
)

// ErrNoRecord is ...
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
