package utils

import (
	"crypto/rand"
	"encoding/base32"
	"strings"
	"time"
)

// NewSecret returns a underlying base32 encoded secret. This should only be displayed
// the first time a user enables 2FA, and should be transmitted over a secure connection.
// Useful for supporting TOTP clients that don't support QR scanning.
func NewSecret() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return strings.ToUpper(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b))
}

func Calculate2FACode(secret string) int {
	return ComputeCode(secret, time.Now().UTC().Unix()/30)
}

func Verify2FACode(secret string, password string) (bool, error) {
	o := OTPConfig{
		Secret:     secret,
		WindowSize: 2,
		UTC:        true,
	}
	return o.Authenticate(password)
}
