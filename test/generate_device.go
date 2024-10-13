package test

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/mohithchintu/final_year_project/models"
)

func generateID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func NewDevice() *models.Device {
	return &models.Device{
		ID:        generateID(),
		Share:     &models.Share{X: big.NewInt(0), Y: big.NewInt(0)},
		Threshold: 3,
	}
}
