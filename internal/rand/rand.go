package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	SessionTokenBytes = 32
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	bRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("Bytes: %w", err)
	}
	if bRead < n {
		return nil, fmt.Errorf("Bytes:  Not enough random bytes read")
	}
	return b, nil
}

// String returns the specified number of random bytes encoded as a string using base64
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("String: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func SessionToken() (string, error) {
	return String(SessionTokenBytes)
}
