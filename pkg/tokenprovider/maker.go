package tokenprovider

import (
	"time"
)

// TokenMaker is an interface for managing tokens
type TokenMaker interface {
	// CreateToken creates a new token for a specific userId and duration
	CreateToken(userId string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
