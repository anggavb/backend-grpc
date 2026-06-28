package token

import (
	"fmt"
	"log"
	"time"

	"aidanwoods.dev/go-paseto"
	"golang.org/x/crypto/chacha20"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < chacha20.KeySize {
		return nil, fmt.Errorf("symmetric key is too short, must be at least %d characters", chacha20.KeySize)
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		log.Println("Error occurred while creating symmetric key:", err)
		return nil, err
	}

	return &PasetoMaker{symmetricKey: key}, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token := paseto.NewToken()

	token.SetIssuedAt(payload.IssuedAt.Time)
	token.SetNotBefore(time.Now())
	token.SetExpiration(payload.ExpiresAt.Time)

	// Store the payload in the token claims
	err = token.Set("payload", payload)
	if err != nil {
		log.Println("Error occurred while setting payload:", err)
		return "", err
	}

	encrypted := token.V4Encrypt(maker.symmetricKey, nil)
	return encrypted, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.ValidAt(time.Now()))

	parsedToken, err := parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		log.Println("Error occurred while parsing token:", err)
		return nil, err
	}

	// Extract the payload from the token claims
	var payload *Payload
	err = parsedToken.Get("payload", &payload)
	if err != nil {
		log.Println("Error occurred while parsing payload:", err)
		return nil, err
	}

	return payload, nil
}
