package repositories

import (
	"errors"
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user domain.User) error
	GetById(id uuid.UUID) (domain.User, error)
}

type FakeUserRepository struct {
	memory map[uuid.UUID]domain.User
}

func NewFakeUserRepository(memory *map[uuid.UUID]domain.User) UserRepository {
	return UserRepository(&FakeUserRepository{memory: *memory})
}

func (repo *FakeUserRepository) Save(user domain.User) error {
	repo.memory[user.ID] = user
	return nil
}

func (repo *FakeUserRepository) GetById(id uuid.UUID) (domain.User, error) {
	if val, ok := repo.memory[id]; ok {
		return val, nil
	}

	return domain.User{}, errors.New(fmt.Sprintf("User with ID %s not found", id.String()))
}
