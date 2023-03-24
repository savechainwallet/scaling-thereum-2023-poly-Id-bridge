package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateDID(didId string) (res string, privateKey string, err error) {

	getPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
	privateKeyBytes := crypto.FromECDSA(getPrivateKey)
	privateKey = hexutil.Encode(privateKeyBytes)
	res = fmt.Sprintf("%s:%s", didId, thePublicAddress[2:])
	return

}
