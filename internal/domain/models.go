package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email,omitempty"`
}

type Album struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Owner *User     `json:"user"`
}

func CreateNewUser(name string, email string) (*User, error) {
	if id, err := uuid.NewUUID(); err == nil {
		return &User{ID: id, Name: name, Email: email}, nil
	} else {
		return nil, err
	}
}

func (u User) CreateNewAlbum(name string) (Album, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return Album{}, err
	}
	return Album{
		ID:    id,
		Name:  name,
		Owner: &u,
	}, nil
}
