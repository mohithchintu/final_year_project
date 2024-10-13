package main

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project/models"
	"github.com/mohithchintu/final_year_project/sss"
	"github.com/mohithchintu/final_year_project/test"
)

func main() {
	secret := big.NewInt(84092309382497840)

	n := 5
	threshold := 3

	// testing
	devices := make([]*models.Device, 5)

	for i := 0; i < len(devices); i++ {
		devices[i] = test.NewDevice()
	}

	// Step 1: Generate the polynomial for the secret
	coefficients, err := sss.GeneratePolynomial(secret, threshold-1)
	if err != nil {
		fmt.Println("Error generating polynomial:", err)
		return
	}

	// Step 2: Generate shares
	shares, err := sss.GenerateShares(coefficients, n)
	if err != nil {
		fmt.Println("Error generating shares:", err)
		return
	}

	fmt.Println("Generated shares:")

	for i, share := range shares {
		devices[i].Share.X = share.X
		devices[i].Share.Y = share.Y
		fmt.Printf("Device %d: %s\n", i+1, share)
	}

	// Step 3: Threshold
	subsetShares := shares[:threshold]

	// Step 4: Reconstruct the secret
	reconstructedSecret := sss.ReconstructSecret(subsetShares)
	fmt.Println("Reconstructed Secret:", reconstructedSecret)
}
