package tokenprovider

import (
	"github.com/google/uuid"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
	"time"
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expried_at"`
}

// NewPayload creates a new token payload with a specific userId and duration
func NewPayload(id string, duaration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		UserID:    id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duaration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.ErrExpiredToken
	}
	return nil
}
