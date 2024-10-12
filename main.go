package main

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project/sss"
)

func main() {
	secret := big.NewInt(84092309382497840)

	n := 5
	threshold := 3

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
		fmt.Printf("Share %d -> X: %s, Y: %s\n", i+1, share.X.String(), share.Y.String())
	}

	// Step 3: Threshold
	subsetShares := shares[:threshold]

	// Step 4: Reconstruct the secret
	reconstructedSecret := sss.ReconstructSecret(subsetShares)
	fmt.Println("Reconstructed Secret:", reconstructedSecret)
}
