package mocks

import (
	"github.com/ngohoang211020/snippetbox/internal/models"
	"time"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	if id == 1 {
		u := &models.User{
			ID:      1,
			Name:    "Alice",
			Email:   "alice@example.com",
			Created: time.Now(),
		}

		return u, nil
	}

	return nil, models.ErrNoRecord
}
