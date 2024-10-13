package helpers

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/mohithchintu/final_year_project/models"
)

func InitializeDevice(id string, threshold int) *models.Device {
	device := &models.Device{
		ID:        id,
		Peers:     make(map[string]*models.Device),
		Threshold: threshold,
		Curve:     elliptic.P256(),
	}
	GenerateECCKeys(device)
	return device
}

// Generate ECC key pair
func GenerateECCKeys(device *models.Device) error {
	curve := device.Curve
	privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return err
	}
	device.PrivateKey = new(big.Int).SetBytes(privateKey)
	device.PublicKeyX = x
	device.PublicKeyY = y
	return nil
}
