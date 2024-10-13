package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateHMAC(data []byte) []byte {
	h := hmac.New(sha256.New, []byte("static-secret-hmac-key"))
	h.Write(data)
	return h.Sum(nil)
}
