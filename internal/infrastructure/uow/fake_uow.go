package uow

import (
	"github.com/edg956/pixer/helpers"
	"github.com/edg956/pixer/internal/domain"
	"github.com/edg956/pixer/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type FakeUoW struct {
	albumMemory map[uuid.UUID]domain.Album
	userMemory  map[uuid.UUID]domain.User
	Users       *repositories.UserRepository
	Albums      *repositories.AlbumRepository
	checkpoint  *FakeUoW
}

var instance *UnitOfWork

func GetFakeUoW() (UnitOfWork, error) {
	if instance == nil {
		var repo = UnitOfWork(&FakeUoW{})
		instance = &repo
	}
	return *instance, nil
}

func (uow *FakeUoW) begin() {
	uow.checkpoint = &FakeUoW{
		albumMemory: *helpers.CopyMap[uuid.UUID, domain.Album](&uow.albumMemory),
		userMemory:  *helpers.CopyMap[uuid.UUID, domain.User](&uow.userMemory),
	}
}

func (uow *FakeUoW) commit() {
	uow.checkpoint = nil
}

func (uow *FakeUoW) rollback() {
	uow.albumMemory = uow.checkpoint.albumMemory
	uow.userMemory = uow.checkpoint.userMemory
	uow.checkpoint = nil
}
