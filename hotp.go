package otp

import (
	"crypto/hmac"
	"crypto/sha1"
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

//func HTOP(secret string, movingFactor int64, codeDigits int, addChecksum bool, truncationOffset int) string {
//}