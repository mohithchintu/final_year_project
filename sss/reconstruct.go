package sss

import (
	"math/big"

	"github.com/mohithchintu/final_year_project/models"
	"github.com/mohithchintu/final_year_project/utils"
)

func ReconstructSecret(shares []*models.Share) *big.Int {
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
		lagrange := new(big.Int).Mul(numerator, utils.Mod_Inverse(denominator, prime))
		lagrange.Mod(lagrange, prime)
		temp := new(big.Int).Mul(share.Y, lagrange)
		temp.Mod(temp, prime)
		secret.Add(secret, temp)
		secret.Mod(secret, prime)
	}
	return secret
}
