package utils

import "math/big"

func Mod_Inverse(a, p *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, p)
}
