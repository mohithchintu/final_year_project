package helpers

import (
	"github.com/mohithchintu/final_year_project/models"
)

func InitializeDevice(id string, threshold int) *models.Device {
	device := &models.Device{
		ID:        id,
		Peers:     make(map[string]*models.Device),
		Threshold: threshold,
	}
	return device
}

func ShareIds(devices []*models.Device) {
	for i, device := range devices {
		for j, peer := range devices {
			if i != j {
				device.Peers[peer.ID] = peer
			}
		}
	}
}
