package otp

import (
	"fmt"
	"math"
	"encoding/base32"
)

// HOTP is used to generate tokens based on RFC-4226.
//
// Example:
//
//  hotp := &HOTP{Secret: "your-secret", Counter: 1000, Length: 8}
//  token := hotp.Get()
//
// HOTP assumes a set of default values for Secret, Length and Counter.
// If no Secret is informed, HOTP will generate a random one that you
// need to store with the Counter, for future token verifications. Check
// this package constants to see the current default values.
type HOTP struct {
	Secret  string // The secret used to generate the token
	Length  uint8  // The token size, with a maximum determined by MaxLength
	Counter uint64 // The counter used as moving factor
}

func (h *HOTP) setDefaults() {
	if len(h.Secret) == 0 {
		h.Secret = randomString(DefaultRandomSecretLength)
	}
	if h.Length == 0 {
		h.Length = 6
	}
}

func (h *HOTP) normalize() {
	if h.Length > MaxLength {
		h.Length = MaxLength
	}
}

// Get a token generated with the current HOTP settings
func (h *HOTP) Get() string {
	h.setDefaults()
	h.normalize()
	text := counterToBytes(h.Counter)
	secretBytes, _ := base32.StdEncoding.DecodeString(h.Secret)
	hash := hmacSHA1(secretBytes, text)
	binary := truncate(hash)
	otp := int64(binary) % int64(math.Pow10(int(h.Length)))
	hotp := fmt.Sprintf(fmt.Sprintf("%%0%dd", h.Length), otp)
	return hotp
}