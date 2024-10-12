package models

import "math/big"

// Share represents a share in Shamir's Secret Sharing
type Share struct {
	X *big.Int
	Y *big.Int
}
