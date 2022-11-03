package repo

import (
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/ent"
)

type IRepo interface {
	IUserRepo
}

type IUserRepo interface {
	GetUserByID(id uuid.UUID) (*ent.User, error)
}
