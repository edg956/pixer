package application

import (
	"github.com/edg956/pixer/internal/domain"
	"github.com/edg956/pixer/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

func CreateNewAlbum(userId uuid.UUID, albumName string) (uuid.UUID, error) {
	// TODO: Start transaction boundary at some point (UoW?)
	userRepo, err := repositories.GetUserRepository()
	if err != nil {
		return uuid.Nil, err
	}

	albumRepo, err := repositories.GetAlbumRepository()
	if err != nil {
		return uuid.Nil, err
	}

	var user domain.User
	user, err = userRepo.GetById(userId)
	if err != nil {
		return uuid.Nil, err
	}
	album, err := user.CreateNewAlbum(albumName)

	if err != nil {
		return uuid.Nil, err
	}

	// store album in repository
	err = albumRepo.Create(album)

	if err != nil {
		return uuid.Nil, err
	}

	return album.Id, nil
}
