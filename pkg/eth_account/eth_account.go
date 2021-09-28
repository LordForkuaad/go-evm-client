package eth_account

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// UserAccount keeps all the necessary user data needed
// to sign transactions and provide visibility
type UserAccount struct {
	Account    common.Address
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

// CreateAccount given a private key in hex format we extract
// the Public and Private keys in ECDSA format as well as the
// Address.
func CreateAccount(privateHexKey string) (*UserAccount, error) {
	privateKey, err := crypto.HexToECDSA(privateHexKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err1 := errors.New("error: cannot assert type: publicKey is " +
			"not of type *ecdsa.PublicKey")
		return nil, err1
	}
	userAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &UserAccount{userAddress, publicKeyECDSA, privateKey}, nil
}
