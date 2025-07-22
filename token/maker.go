package token

import "time"

type Maker interface {
	CreateToken(userName string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
