package keygen

import (
	"crypto/rand"
	"errors"
)

type Keygen struct {
}

func NewKeygen() *Keygen {
	return &Keygen{}
}

func (kg *Keygen) Generate(length uint) (string, error) {
	if length < 3 {
		return "", errors.New("The parameter to generate a key must be greather than 3")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortKey := make([]byte, length)

	_, err := rand.Read(shortKey)
	if err != nil {
		return "", err
	}

	for i := range shortKey {
		shortKey[i] = charset[int(shortKey[i])%len(charset)]
	}

	return string(shortKey), nil
}
