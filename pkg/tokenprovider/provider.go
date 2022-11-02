package tokenprovider

import (
	"time"
)

var (
	ZToken Provider
)

// Provider is an interface for managing tokens
type Provider interface {
	// CreateToken creates a new token for a specific userId and duration
	CreateToken(userId string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
