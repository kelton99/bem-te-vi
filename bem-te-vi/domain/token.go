package domain

import (
	"fmt"
	"time"
)

const (
	TokenLen = 32
)

type Token struct {
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}

func (t *Token) Validate() error {
	if t.Token == "" {
		return fmt.Errorf("token is required")
	}
	if len(t.Token) != TokenLen {
		return fmt.Errorf("token must be %d characters", TokenLen)
	}
	if t.Expiry.IsZero() {
		return fmt.Errorf("expiry is required")
	}
	return nil
}
