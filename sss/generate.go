package sss

import (
	"crypto/rand"
	"math/big"

	"github.com/mohithchintu/final_year_project/models"
)

func GeneratePolynomial(secret *big.Int, degree int) ([]*big.Int, error) {
	coefficients := make([]*big.Int, degree+1)
	coefficients[0] = secret
	for i := 1; i <= degree; i++ {
		coeff, err := rand.Int(rand.Reader, prime)
		if err != nil {
			return nil, err
		}
		coefficients[i] = coeff
	}
	return coefficients, nil
}

func GenerateShares(coefficients []*big.Int, n int) ([]*models.Share, error) {
	shares := make([]*models.Share, n)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		y := evaluatePolynomial(coefficients, x)
		shares[i-1] = &models.Share{X: x, Y: y}
	}
	return shares, nil
}

func evaluatePolynomial(coefficients []*big.Int, x *big.Int) *big.Int {
	result := big.NewInt(0)
	for i := len(coefficients) - 1; i >= 0; i-- {
		result.Mul(result, x)
		result.Add(result, coefficients[i])
		result.Mod(result, prime)
	}
	return result
}
