package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateHMAC(data []byte) []byte {
	timestamp := time.Now().Unix()
	timestampBytes := []byte(fmt.Sprintf("%d", timestamp))
	h := hmac.New(sha256.New, timestampBytes)
	h.Write(data)
	return h.Sum(nil)
}
