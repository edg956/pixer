package repositories

import (
	"errors"
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(id uuid.UUID) (domain.User, error)
}

type InMemoryUserRepository struct {
	memory map[string]domain.User
}

func (repo *InMemoryUserRepository) GetById(id uuid.UUID) (domain.User, error) {
	if val, ok := repo.memory[id.String()]; ok {
		return val, nil
	}

	return domain.User{}, errors.New(fmt.Sprintf("User with ID %s not found", id.String()))
}

var userInstance *UserRepository

func GetUserRepository() (*UserRepository, error) {
	if userInstance == nil {
		var repo UserRepository
		repo = &InMemoryUserRepository{memory: make(map[string]domain.User)}
		userInstance = &repo
	}
	return userInstance, nil
}
