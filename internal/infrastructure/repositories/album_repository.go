package repositories

import (
	"errors"
	"fmt"
	"github.com/edg956/pixer/internal/domain"
)

type AlbumRepository interface {
	Create(album domain.Album) error
}

type InMemoryAlbumRepository struct {
	memory map[string]domain.Album
}

func (repo *InMemoryAlbumRepository) Create(album domain.Album) error {
	albumId := album.Id.String()
	if _, exists := repo.memory[albumId]; exists {
		return errors.New(fmt.Sprintf("Album with Id %s already exists.", albumId))
	}
	repo.memory[albumId] = album
	return nil
}

var albumInstance *AlbumRepository

func GetAlbumRepository() (*AlbumRepository, error) {
	if albumInstance == nil {
		var repo AlbumRepository
		repo = &InMemoryAlbumRepository{memory: make(map[string]domain.Album)}
		albumInstance = &repo
	}
	return albumInstance, nil
}
