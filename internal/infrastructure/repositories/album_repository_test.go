package repositories

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCreate(t *testing.T) {
	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatal("Could not create UUID")
		return
	}

	user, err := domain.CreateNewUser("Test users", "test@server.com")
	if err != nil {
		t.Fatal(fmt.Sprintf("Could not create user %s", err))
		return
	}

	album := domain.Album{ID: id, Name: "Test Album", Owner: user}
	memory := make(map[uuid.UUID]domain.Album)

	repository := InMemoryAlbumRepository{memory: memory}

	err = repository.Create(&album)

	if err != nil {
		t.Fatal("Error creating album")
		return
	}

	if value, ok := memory[id]; ok {
		if value != album {
			t.Errorf("Expected %v, got %v instead", album, value)
		}
	} else {
		t.Errorf("Album not persisted")
	}
}
