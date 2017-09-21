package otp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
)

func generateRandomSecret(size int, encodeToBase32 bool) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	if encodeToBase32 {
		return base32.StdEncoding.EncodeToString(bytes)
	}
	return string(bytes)
}

func counterToBytes(counter uint64) (text []byte) {
	text = make([]byte, 8)
	for i := 7; i >= 0; i-- {
		text[i] = byte(counter & 0xff)
		counter = counter >> 8
	}
	return
}

func hmacSHA(hash Hash, key, text []byte) []byte {
	h := sha1.New
	if hash != DefaultHashAlgo {
		switch(hash) {
		case SHA256:
			h = sha256.New
		case SHA512:
			h = sha512.New
		}
	}
	H := hmac.New(h, key)
	H.Write([]byte(text))
	return H.Sum(nil)
}

func truncate(hash []byte) int {
	offset := int(hash[len(hash)-1] & 0xf)
	return ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)
}
