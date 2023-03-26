package utils

import (
	"encoding/pem"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(ttl time.Duration, content interface{}, k []byte, iis string) (token string, err error) {
	block, _ := pem.Decode(k)

	key, err := crypto.ToECDSA(block.Bytes)

	//key, err := jwt.ParseECPrivateKeyFromPEM(k)
	if err != nil {
		return
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = content             // Our custom data.
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.
	claims["iis"] = iis

	token, err = jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(key)
	return
}
