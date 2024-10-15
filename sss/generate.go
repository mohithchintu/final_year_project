package sss

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/mohithchintu/final_year_project/models"
)

func GenerateAndSharePolynomial(device *models.Device, degree int) ([]*big.Int, error) {
	rand.Seed(time.Now().UnixNano())
	coefficients := make([]*big.Int, degree+1)
	secret := big.NewInt(int64(rand.Intn(10000000999)))
	coefficients[0] = secret

	fmt.Println(device.ID, "generated secret:", secret)

	for i := 1; i <= degree; i++ {
		coeff := big.NewInt(int64(rand.Intn(10000000999)))
		coefficients[i] = coeff
	}

	ShareCoefficientsWithPeers(device, coefficients)
	return coefficients, nil
}

// Share coefficients with peers (simulated network communication)
func ShareCoefficientsWithPeers(device *models.Device, coefficients []*big.Int) {
	for _, peer := range device.Peers {
		fmt.Printf("Device %s sends coefficients to Device %s\t", device.ID, peer.ID)
		peer.ReceiveCoefficients(device.ID, coefficients)
	}
}

// Generate shares by evaluating the polynomial at distinct x-values
func GenerateShares(coefficients []*big.Int, n int) []*models.Share {
	shares := make([]*models.Share, n)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		y := EvaluatePolynomial(coefficients, x)
		shares[i-1] = &models.Share{X: x, Y: y}
	}
	return shares
}

// Evaluate the polynomial at a specific x-value
func EvaluatePolynomial(coefficients []*big.Int, x *big.Int) *big.Int {
	result := big.NewInt(0)
	for i := len(coefficients) - 1; i >= 0; i-- {
		result.Mul(result, x)
		result.Add(result, coefficients[i])
		result.Mod(result, prime)
	}
	return result
}
