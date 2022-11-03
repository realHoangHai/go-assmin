package paseto

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
	"time"
)

// pasetoProvider is a PASETO token maker
type pasetoProvider struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoProvider creates a new paseto provider
func NewPasetoProvider(symmetricKey string) tokenprovider.TokenMaker {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		panic(fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize))
	}

	return &pasetoProvider{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
}

// CreateToken creates a new token for a specific username and duration
func (p *pasetoProvider) CreateToken(userId string, duration time.Duration) (string, *tokenprovider.Payload, error) {
	payload, err := tokenprovider.NewPayload(userId, duration)
	if err != nil {
		return "", payload, err
	}
	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (p *pasetoProvider) VerifyToken(token string) (*tokenprovider.Payload, error) {
	payload := &tokenprovider.Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
