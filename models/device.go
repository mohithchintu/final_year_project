package models

import "math/big"

type Device struct {
	ID         string
	PrivateKey *big.Int
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	Share      *Share
	Peers      map[string]*Device
	// Peers map[string]
	GroupKey  []byte
	Threshold int
}
