package repo

import (
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/ent"
	"github.com/realHoangHai/go-assmin/internal/ent/session"
)

func (r *Repo) CreateSession(data *ent.Session) (*ent.Session, error) {
	result, err := r.Session.
		Create().
		SetID(data.ID).
		SetUserID(data.UserID).
		SetRefreshToken(data.RefreshToken).
		SetUserAgent(data.UserAgent).
		SetClientIP(data.ClientIP).
		SetIsBlocked(data.IsBlocked).
		SetExpireTime(data.ExpireTime).
		Save(r.ctx)
	return result, err
}

func (r *Repo) GetSessionByID(id uuid.UUID) (*ent.Session, error) {
	result, err := r.Session.
		Query().
		Where(
			session.ID(id),
		).Only(r.ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.ErrEntityNotFound("User", err)
		}
		return nil, errors.ErrDB(err)
	}
	return result, nil
}
