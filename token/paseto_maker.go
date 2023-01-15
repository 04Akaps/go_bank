package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)


type PasetoMaker struct {
	paseto *paseto.V2
	key []byte
}

// make paseto instance
func NewPasetoMaker(key string) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size!")
	}

	maker := &PasetoMaker {
		paseto: paseto.NewV2(),
		key : []byte(key),
	}

	return maker, nil
}

func (maker *PasetoMaker)CreateToken(username string, duration time.Duration) (string ,error){
	payload, err:= NewPayload(username, duration) 

	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.key, payload, nil)
}

func (maker *PasetoMaker)	VerifyToken(token string ) (*Payload, error){
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.key, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

