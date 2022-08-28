package repositories

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type AlbumRepository interface {
	Save(album *domain.Album) error
	GetById(id uuid.UUID) (*domain.Album, error)
}

type FakeAlbumRepository struct {
	memory map[uuid.UUID]domain.Album
}

func NewFakeAlbumRepository(memory *map[uuid.UUID]domain.Album) AlbumRepository {
	return AlbumRepository(&FakeAlbumRepository{memory: *memory})
}

func (repo *FakeAlbumRepository) Save(album *domain.Album) error {
	repo.memory[album.ID] = *album
	return nil
}

func (repo *FakeAlbumRepository) GetById(id uuid.UUID) (*domain.Album, error) {
	if val, exists := repo.memory[id]; exists {
		return &val, nil
	}

	return nil, fmt.Errorf("album %s does not exist", id)
}
