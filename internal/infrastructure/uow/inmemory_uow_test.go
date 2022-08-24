package uow

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCommit(t *testing.T) {
	var uow UnitOfWork = &InMemoryUoW{
		AlbumMemory: make(map[uuid.UUID]domain.Album),
		UserMemory:  make(map[uuid.UUID]domain.User),
	}

	ptr := uow.(*InMemoryUoW)
	ptr.begin()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{}

	ptr.AlbumMemory[id] = album
	ptr.commit()

	if value, ok := ptr.AlbumMemory[id]; ok {
		if value != album {
			t.Errorf("Expected %s, instead got %s", album, value)
		}
	} else {
		t.Errorf("Expected AlbumMemory to contain an album with ID %s", id)
	}
}

func TestRollback(t *testing.T) {
	var uow UnitOfWork = &InMemoryUoW{
		AlbumMemory: make(map[uuid.UUID]domain.Album),
		UserMemory:  make(map[uuid.UUID]domain.User),
	}

	ptr := uow.(*InMemoryUoW)
	ptr.begin()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{}

	ptr.AlbumMemory[id] = album
	ptr.rollback()

	if value, ok := ptr.AlbumMemory[id]; ok {
		t.Errorf("Expected memory to be empty, instead got %s", value)
	}
}
