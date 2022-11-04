package repo

import (
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/ent"
	"github.com/realHoangHai/go-assmin/internal/ent/user"
	"strings"
)

func (r *Repo) CreateUser(data *ent.User) error {
	_, err := r.User.
		Create().
		SetName(strings.TrimSpace(data.Name)).
		SetEmail(strings.TrimSpace(data.Email)).
		SetPhone(strings.TrimSpace(data.Phone)).
		SetPassword(data.Password).
		SetStatus(data.Status).
		Save(r.ctx)
	return errors.ErrDB(err)
}

func (r *Repo) GetUserByID(id uuid.UUID) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) GetUserNotDisableByEmail(email string) (*ent.User, error) {
	result, err := r.User.
		Query().
		Where(
			user.EmailEqualFold(email),
		).
		Only(r.ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.ErrEntityNotFound("User", err)
		}
		return nil, errors.ErrDB(err)
	}
	return result, err
}
