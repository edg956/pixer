package cases

import (
	unitOfWork "github.com/edg956/pixer/internal/infrastructure/uow"
	"github.com/google/uuid"
)

type NewAlbumCommand struct {
	UserID    uuid.UUID
	AlbumName string
}

func createNewAlbum(cmd NewAlbumCommand, uow unitOfWork.UnitOfWork) (uuid.UUID, error) {
	user, err := uow.GetUsers().GetById(cmd.UserID)
	if err != nil {
		return uuid.Nil, err
	}

	album, err := user.CreateNewAlbum(cmd.AlbumName)
	if err != nil {
		return uuid.Nil, err
	}

	// store album in repository
	err = uow.GetAlbums().Save(&album)

	if err != nil {
		return uuid.Nil, err
	}

	return album.ID, nil
}

func CreateNewAlbum(cmd NewAlbumCommand, uow unitOfWork.UnitOfWork) (uuid.UUID, error) {
	return executeInTransaction[
		NewAlbumCommand, uuid.UUID, useCase[NewAlbumCommand, uuid.UUID],
	](createNewAlbum, cmd, uow)
}
