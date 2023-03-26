package utils

import (
	"encoding/pem"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateDID(didId string) (res string, privateKey []byte, err error) {

	getPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
	privateKeyPEM := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: crypto.FromECDSA(getPrivateKey),
	}
	privateKey = pem.EncodeToMemory(privateKeyPEM)

	res = fmt.Sprintf("%s:%s", didId, thePublicAddress[2:])
	return

}
