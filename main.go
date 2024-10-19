package main

import (
	"fmt"

	"github.com/mohithchintu/final_year_project/helpers"
	"github.com/mohithchintu/final_year_project/models"
	"github.com/mohithchintu/final_year_project/sss"
	"github.com/mohithchintu/final_year_project/utils"
)

const (
	numDevices = 6 // Number of devices
	threshold  = 4 // Threshold for reconstruction
)

var devices []*models.Device

func main() {
	// Initialize devices
	for i := 1; i <= numDevices; i++ {
		deviceID := fmt.Sprintf("Device%d", i)
		device := helpers.InitializeDevice(deviceID, threshold)
		devices = append(devices, device)
	}

	// Establish peer relationships (each device knows all other devices)
	// for i, device := range devices {
	// 	for j, peer := range devices {
	// 		if i != j {
	// 			device.Peers[peer.ID] = peer
	// 		}
	// 	}
	// }
	helpers.ShareIds(devices)

	// Each device generates its own polynomial and shares the coefficients with peers
	for _, device := range devices {
		coefficients, _ := sss.GenerateAndSharePolynomial(device, threshold-1)
		shares := sss.GenerateShares(coefficients, numDevices)
		utils.DistributeShares(device, shares) // No need to pass GroupKey here
	}

	// Simulate group key reconstruction with a threshold of shares
	reconstructedKey := sss.HandleDeviceFailure(devices, threshold)
	fmt.Println("Final Reconstructed Group Key:", reconstructedKey)
}
