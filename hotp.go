package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func movingFactorToCounter(movingFactor int64) (text []byte) {
	text = make([]byte, 8)
	for i := 7; i >= 0; i-- {
		text[i] = byte(movingFactor & 0xff)
		movingFactor = movingFactor >> 8
	}
	return
}

func hmacSHA1(key, text []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write([]byte(text))
	return h.Sum(nil)
}

var digitsPower = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

func HOTP(secret string, movingFactor int64, codeDigits int, truncationOffset int) string {
	text := movingFactorToCounter(movingFactor)
	hash := hmacSHA1([]byte(secret), text)
	offset := int(hash[len(hash)-1] & 0xf)

	if (0 <= truncationOffset) && (truncationOffset < len(hash)-4) {
		offset = truncationOffset
	}

	binary := ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)

	otp := int(binary) % digitsPower[codeDigits]
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", codeDigits), otp)
}