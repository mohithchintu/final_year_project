package models

import (
	"crypto/elliptic"
	"math/big"

	"github.com/mohithchintu/final_year_project/hmac"
)

type Device struct {
	ID         string
	PrivateKey *big.Int
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	Share      *Share
	Peers      map[string]*Device
	Curve      elliptic.Curve
	Threshold  int
	GroupKey   *big.Int
}

// ReceiveCoefficients allows a device to receive polynomial coefficients from another device
func (device *Device) ReceiveCoefficients(senderID string, coefficients []*big.Int) {
	println("Device", device.ID, "received coefficients from", senderID)
}

type Message struct {
	SenderID string
	Data     struct {
		Share *Share
		HMAC  []byte
	}
}

// Receive share and verify its integrity
func (device *Device) ReceiveShare(msg Message) {
	data := msg.Data

	shareBytes := append(data.Share.X.Bytes(), data.Share.Y.Bytes()...)
	if hmac.VerifyHMAC(shareBytes, data.HMAC) {
		println("Device", device.ID, "received valid share from", msg.SenderID)
		device.Share = data.Share
	} else {
		println("Device", device.ID, "received invalid share from", msg.SenderID)
	}
}
