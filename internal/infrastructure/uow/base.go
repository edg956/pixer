package uow

import "github.com/edg956/pixer/internal/infrastructure/repositories"

type UnitOfWork interface {
	GetUsers() repositories.UserRepository
	GetAlbums() repositories.AlbumRepository
	Begin()
	Commit()
	Rollback()
}
