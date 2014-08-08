package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"math"
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

func truncate(hash []byte) int {
	offset := int(hash[len(hash)-1] & 0xf)
	return ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)
}

func GenerateHOTP(secret string, movingFactor int64, codeDigits int8) string {
	if codeDigits > 10 {
		codeDigits = 10
	}
	text := movingFactorToCounter(movingFactor)
	hash := hmacSHA1([]byte(secret), text)
	binary := truncate(hash)
	otp := int64(binary) % int64(math.Pow10(int(codeDigits)))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", codeDigits), otp)
}