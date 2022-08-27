package repositories

import (
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type AlbumRepository interface {
	Save(album *domain.Album) error
}

type FakeAlbumRepository struct {
	memory map[uuid.UUID]domain.Album
}

func (repo *FakeAlbumRepository) Save(album *domain.Album) error {
	repo.memory[album.ID] = *album
	return nil
}

var albumInstance *AlbumRepository

func GetAlbumRepository() (AlbumRepository, error) {
	if albumInstance == nil {
		var repo AlbumRepository = &FakeAlbumRepository{memory: make(map[uuid.UUID]domain.Album)}
		albumInstance = &repo
	}
	return *albumInstance, nil
}
