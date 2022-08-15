package application

import (
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

func CreateNewAlbum(userId uuid.UUID, albumName string) (uuid.UUID, error) {
	// TODO: Start transaction boundary at some point (UoW?)
	// TODO: fetch user from repository
	user := domain.User{Id: userId, Name: "Anonymous"}
	album, err := user.CreateNewAlbum(albumName)

	if err != nil {
		return uuid.Nil, err
	}

	// TODO: store album in repository

	return album.Id, nil
}
