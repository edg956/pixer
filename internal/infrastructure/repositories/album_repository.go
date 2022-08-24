package repositories

import (
	"errors"
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type AlbumRepository interface {
	Create(album domain.Album) error
}

type InMemoryAlbumRepository struct {
	memory map[uuid.UUID]domain.Album
}

func (repo *InMemoryAlbumRepository) Create(album domain.Album) error {
	albumId := album.Id
	if _, exists := repo.memory[albumId]; exists {
		return errors.New(fmt.Sprintf("Album with Id %s already exists.", albumId))
	}
	repo.memory[albumId] = album
	return nil
}

var albumInstance *AlbumRepository

func GetAlbumRepository() (AlbumRepository, error) {
	if albumInstance == nil {
		var repo AlbumRepository = &InMemoryAlbumRepository{memory: make(map[uuid.UUID]domain.Album)}
		albumInstance = &repo
	}
	return *albumInstance, nil
}
