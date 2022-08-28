package uow

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/edg956/pixer/internal/infrastructure/repositories"
	"github.com/google/uuid"
	"testing"
)

func TestCommit(t *testing.T) {
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

	uow.Begin()

	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{ID: id, Name: "Test Album"}
	albumMemory[id] = album
	uow.Commit()

	if value, ok := albumMemory[id]; ok {
		if value != album {
			t.Errorf("Expected %v, instead got %v", album, value)
		}
	} else {
		t.Errorf("Expected albumMemory to contain an album with ID %s", id)
	}
}

func TestRollback(t *testing.T) {
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

	uow.Begin()

	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{ID: id, Name: "Test Album"}

	albumMemory[id] = album
	uow.Rollback()

	if value, ok := albumMemory[id]; ok {
		t.Errorf("Expected memory to be empty, instead got %v", value)
	}
}
