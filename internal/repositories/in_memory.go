package repositories

import (
	"github.com/edg956/pixer/internal/models"
)

type InMemorySpaceRepository struct {
	memory map[string]interface{}
}

func (r InMemorySpaceRepository) CreateSpace() models.Space {
	return new(models.Space)
}
