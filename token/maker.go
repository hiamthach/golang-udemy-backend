package token

import (
	"time"
)

// Maker is a interface for creating tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
