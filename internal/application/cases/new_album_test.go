package cases

import (
	"github.com/edg956/pixer/internal/domain"
	unitOfWork "github.com/edg956/pixer/internal/infrastructure/uow"
	"testing"
)

func TestCreateNewAlbum(t *testing.T) {
	user, err := domain.CreateNewUser("Test User", "test@email.com")
	if err != nil {
		t.Fatal("Could not create user")
		return
	}

	uow, err := unitOfWork.GetFakeUoW()
	if err != nil {
		t.Fatal("Could not instantiate UoW")
		return
	}

	if err = uow.GetUsers().Save(*user); err != nil {
		t.Fatal("Could not persist user")
		return
	}

	cmd := NewAlbumCommand{UserID: user.ID, AlbumName: "My Album"}

	r, err := CreateNewAlbum(cmd, uow)

	if err != nil {
		t.Errorf("Could not create album from command %v", cmd)
		return
	}

	value, err := uow.GetAlbums().GetById(r)

	if err != nil {
		t.Errorf("Album %s not persisted", r)
		return
	}

	if value.Name != "My Album" {
		t.Errorf("Expected album to be named My Album. Got %s instead", value.Name)
	}
}
