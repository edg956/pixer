package repositories

import (
	"fmt"
	"github.com/edg956/pixer/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestSaveUser(t *testing.T) {
	user, err := domain.CreateNewUser("Test users", "test@server.com")
	if err != nil {
		t.Fatal(fmt.Sprintf("Could not create user %s", err))
		return
	}

	memory := make(map[uuid.UUID]domain.User)

	repository := FakeUserRepository{memory: memory}

	err = repository.Save(*user)

	if err != nil {
		t.Fatal("Error creating album")
		return
	}

	if value, ok := memory[user.ID]; ok {
		if value != *user {
			t.Errorf("Expected %v, got %v instead", *user, value)
		}
	} else {
		t.Errorf("User not persisted")
	}
}

func TestGetUserById(t *testing.T) {
	user, err := domain.CreateNewUser("Test User", "test@email.com")
	if err != nil {
		t.Fatalf("Could not create user %s", err)
		return
	}

	memory := map[uuid.UUID]domain.User{
		user.ID: *user,
	}

	repository := FakeUserRepository{memory: memory}

	obtained, err := repository.GetById(user.ID)

	if err != nil {
		t.Fatalf("Error retrieving user with ID %s", user.ID)
	}

	if obtained != *user {
		t.Errorf("Expected %v, got %v instead.", user, obtained)
	}
}
