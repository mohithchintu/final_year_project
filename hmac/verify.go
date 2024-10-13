package hmac

import (
	"crypto/hmac"
)

func VerifyHMAC(data, receivedMac []byte) bool {
	expectedMac := GenerateHMAC(data)
	return hmac.Equal(expectedMac, receivedMac)
}
