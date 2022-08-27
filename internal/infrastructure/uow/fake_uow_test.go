package uow

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCommit(t *testing.T) {
	var uow UnitOfWork = &FakeUoW{
		albumMemory: make(map[uuid.UUID]domain.Album),
		userMemory:  make(map[uuid.UUID]domain.User),
	}

	ptr := uow.(*FakeUoW)
	ptr.begin()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{}

	ptr.albumMemory[id] = album
	ptr.commit()

	if value, ok := ptr.albumMemory[id]; ok {
		if value != album {
			t.Errorf("Expected %v, instead got %v", album, value)
		}
	} else {
		t.Errorf("Expected albumMemory to contain an album with ID %s", id)
	}
}

func TestRollback(t *testing.T) {
	var uow UnitOfWork = &FakeUoW{
		albumMemory: make(map[uuid.UUID]domain.Album),
		userMemory:  make(map[uuid.UUID]domain.User),
	}

	ptr := uow.(*FakeUoW)
	ptr.begin()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating UUID: %s", err))
	}

	album := domain.Album{}

	ptr.albumMemory[id] = album
	ptr.rollback()

	if value, ok := ptr.albumMemory[id]; ok {
		t.Errorf("Expected memory to be empty, instead got %v", value)
	}
}
