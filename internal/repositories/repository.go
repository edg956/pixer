package repositories

import (
	"github.com/edg956/pixer/internal/models"
)


type SpaceRepository interface {
	CreateSpace() models.Space
}
