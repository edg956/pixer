package domain

import (
	"github.com/google/uuid"
)

type User struct {
	id uuid.UUID
	name string
	email string
}


type Album struct {
	id uuid.UUID
	name string
	owner User
}
