package mock

import (
	"time"

	"github.com/cedrickchee/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@foo.bar",
	Created: time.Now(),
}

// UserModel is user model.
type UserModel struct{}

// Insert method adds a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@foo.bar":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate method verifies whether a user exists with the provided email
// address and password. This will return the relevant user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@foo.bar":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get method fetches details for a specific user based on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
