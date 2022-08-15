package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email,omitempty"`
}

type Album struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Owner User      `json:"user"`
}

func (u User) CreateNewAlbum(name string) (Album, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return Album{}, err
	}
	return Album{
		Id:    id,
		Name:  name,
		Owner: u,
	}, nil
}
