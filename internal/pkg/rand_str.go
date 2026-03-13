package pkg

import (
	"crypto/rand"
	"fmt"
)

func NewRandStr(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("rand.Read failed: %w", err)
	}
	return fmt.Sprintf("%x", b), nil
}
