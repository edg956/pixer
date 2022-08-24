package uow

import (
	"github.com/edg956/pixer/helpers"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type InMemoryUoW struct {
	AlbumMemory map[uuid.UUID]domain.Album
	UserMemory  map[uuid.UUID]domain.User
	checkpoint  *InMemoryUoW
}

func (uow *InMemoryUoW) begin() {
	uow.checkpoint = &InMemoryUoW{
		AlbumMemory: *helpers.CopyMap[uuid.UUID, domain.Album](&uow.AlbumMemory),
		UserMemory:  *helpers.CopyMap[uuid.UUID, domain.User](&uow.UserMemory),
	}
}

func (uow *InMemoryUoW) commit() {
	uow.checkpoint = nil
}

func (uow *InMemoryUoW) rollback() {
	uow.AlbumMemory = uow.checkpoint.AlbumMemory
	uow.UserMemory = uow.checkpoint.UserMemory
	uow.checkpoint = nil
}
