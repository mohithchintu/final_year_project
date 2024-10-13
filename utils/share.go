package utils

import (
	"github.com/mohithchintu/final_year_project/hmac"
	"github.com/mohithchintu/final_year_project/models"
)

// Distribute shares to peers (without using GroupKey during this phase)
func DistributeShares(device *models.Device, shares []*models.Share) {
	i := 0
	for _, peer := range device.Peers {
		shareBytes := append(shares[i].X.Bytes(), shares[i].Y.Bytes()...)
		hmac := hmac.GenerateHMAC(shareBytes)
		peer.ReceiveShare(models.Message{
			SenderID: device.ID,
			Data: struct {
				Share *models.Share
				HMAC  []byte
			}{
				Share: shares[i],
				HMAC:  hmac,
			},
		})
		i++
	}
}
