package repo

import (
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/ent"
)

type IRepo interface {
	IUserRepo
	ISessionRepo
}

type IUserRepo interface {
	CreateUser(data *ent.User) error
	GetUserByID(id uuid.UUID) (*ent.User, error)
	GetUserNotDisableByEmail(email string) (*ent.User, error)
}

type ISessionRepo interface {
	CreateSession(data *ent.Session) (*ent.Session, error)
	GetSessionByID(id uuid.UUID) (*ent.Session, error)
}
