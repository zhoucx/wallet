package pkg

import (
	"fmt"
	"time"
)

func NewUid() (string, error) {
	randSuffix, err := NewRandStr(8)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), randSuffix), nil
}
