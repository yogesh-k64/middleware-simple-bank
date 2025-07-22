package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < 32 {
		return nil, fmt.Errorf("invalid secret key length: must be at least %d characters", minSecretKeyLength)
	}

	pasetoMaker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return pasetoMaker, nil
}

func (p PasetoMaker) CreateToken(userName string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userName, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}
func (p PasetoMaker) VerifyToken(token string) (*Payload, error) {
	pasetoPayload := &Payload{}
	err := p.paseto.Decrypt(token, p.symmetricKey, pasetoPayload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if err := pasetoPayload.Valid(); err != nil {
		return nil, err
	}

	return pasetoPayload, nil
}
