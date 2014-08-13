package otp

import (
	"fmt"
	"math"
)

type HOTP struct {
	Secret  string
	Length  uint8
	Counter uint64
}

func (h *HOTP) setDefaults() {
	if len(h.Secret) == 0 {
		h.Secret = randomString(100)
	}
	if h.Length == 0 {
		h.Length = 6
	}
}

func (h *HOTP) normalize() {
	if h.Length > 10 {
		h.Length = 10
	}
}

func (h *HOTP) Get() string {
	h.setDefaults()
	h.normalize()
	text := counterToBytes(h.Counter)
	hash := hmacSHA1([]byte(h.Secret), text)
	binary := truncate(hash)
	otp := int64(binary) % int64(math.Pow10(int(h.Length)))
	hotp := fmt.Sprintf(fmt.Sprintf("%%0%dd", h.Length), otp)
	return hotp
}