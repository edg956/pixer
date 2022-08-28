package uow

import (
	"github.com/edg956/pixer/helpers"
	"github.com/edg956/pixer/internal/domain"
	"github.com/edg956/pixer/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type FakeUoW struct {
	albumMemory *map[uuid.UUID]domain.Album
	userMemory  *map[uuid.UUID]domain.User
	Users       repositories.UserRepository
	Albums      repositories.AlbumRepository
	checkpoint  *FakeUoW
}

var instance UnitOfWork

func GetFakeUoW() (UnitOfWork, error) {
	if instance == nil {
		albumMemory := make(map[uuid.UUID]domain.Album)
		userMemory := make(map[uuid.UUID]domain.User)
		albums := repositories.NewFakeAlbumRepository(&albumMemory)
		users := repositories.NewFakeUserRepository(&userMemory)

		var uow UnitOfWork = &FakeUoW{
			Albums:      albums,
			albumMemory: &albumMemory,
			Users:       users,
			userMemory:  &userMemory,
		}
		instance = uow
	}
	return instance, nil
}

func (uow *FakeUoW) Begin() {
	uow.checkpoint = &FakeUoW{
		albumMemory: helpers.CopyMap[uuid.UUID, domain.Album](uow.albumMemory),
		userMemory:  helpers.CopyMap[uuid.UUID, domain.User](uow.userMemory),
	}
}

func (uow *FakeUoW) Commit() {
	uow.checkpoint = nil
}

func (uow *FakeUoW) Rollback() {
	*uow.albumMemory = *uow.checkpoint.albumMemory
	*uow.userMemory = *uow.checkpoint.userMemory
	uow.checkpoint = nil
}

func (uow *FakeUoW) GetUsers() repositories.UserRepository {
	return uow.Users
}

func (uow *FakeUoW) GetAlbums() repositories.AlbumRepository {
	return uow.Albums
}
