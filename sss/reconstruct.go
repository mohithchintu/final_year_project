package sss

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project/helpers"
	"github.com/mohithchintu/final_year_project/models"
)

func ReconstructGroupKey(shares []*models.Share) *big.Int {
	secret := big.NewInt(0)
	for i, share := range shares {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)
		for j, otherShare := range shares {
			if i != j {
				numerator.Mul(numerator, new(big.Int).Neg(otherShare.X))
				numerator.Mod(numerator, prime)
				diff := new(big.Int).Sub(share.X, otherShare.X)
				denominator.Mul(denominator, diff)
				denominator.Mod(denominator, prime)
			}
		}
		lagrange := new(big.Int).Mul(numerator, helpers.Mod_Inverse(denominator, prime))
		lagrange.Mod(lagrange, prime)
		temp := new(big.Int).Mul(share.Y, lagrange)
		temp.Mod(temp, prime)
		secret.Add(secret, temp)
		secret.Mod(secret, prime)
	}
	fmt.Println("Reconstructed Group Key:", secret)
	return secret
}

// Simulate device failure and still reconstruct the group key
func HandleDeviceFailure(devices []*models.Device, threshold int) *big.Int {
	validShares := []*models.Share{}
	for _, device := range devices {
		if device.Share != nil {
			validShares = append(validShares, device.Share)
		}
		if len(validShares) >= threshold {
			break
		}
	}

	if len(validShares) < threshold {
		fmt.Println("Not enough valid shares to reconstruct the key!")
		return big.NewInt(0)
	}

	return ReconstructGroupKey(validShares)
}
