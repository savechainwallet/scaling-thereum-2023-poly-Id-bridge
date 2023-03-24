package polybase

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type PolybaseClient struct {
	privateKey *ecdsa.PrivateKey
	namespace  string
	url        string
}

func (c *PolybaseClient) createAuthHeader(body []byte) (header string, err error) {
	//Format message to sign
	now := time.Now()
	timestamp := now.UnixMilli()
	messageBodyString := string(body)
	message := fmt.Sprintf("%d.%s", timestamp, messageBodyString)
	hash := crypto.Keccak256Hash([]byte(message))
	sig, err := crypto.Sign(hash.Bytes(), c.privateKey)
	if err != nil {
		return
	}
	sigBase64 := hex.EncodeToString(sig)
	header = fmt.Sprintf("v=0,t=%d,h=eth-personal-sign,sig=0x%s", timestamp, sigBase64)
	return

}

func (c *PolybaseClient) LisRecords(collection string) (map[string]interface{}, error) {
	client := &http.Client{}
	respDecoded := make(map[string]interface{})
	path := url.QueryEscape(fmt.Sprintf("%s/%s", c.namespace, collection))
	url := fmt.Sprintf("%s/v0/collections/%s/records", c.url, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respDecoded, err
	}
	body := make(map[string]interface{})
	bodyBytes, _ := json.Marshal(body)
	header, err := c.createAuthHeader(bodyBytes)
	if err != nil {
		return respDecoded, err
	}
	req.Header.Add("X-Polybase-Signature", header)
	resp, err := client.Do(req)
	if err != nil {
		return respDecoded, err
	}
	err = json.NewDecoder(resp.Body).Decode(&respDecoded)
	return respDecoded, err
}

func (c *PolybaseClient) CreateRecord(collection string, args []interface{}) (map[string]interface{}, error) {
	client := &http.Client{}

	request, _ := json.Marshal(map[string]interface{}{
		"args": args,
	})
	respDecoded := make(map[string]interface{})

	path := url.QueryEscape(fmt.Sprintf("%s/%s", c.namespace, collection))
	url := fmt.Sprintf("%s/v0/collections/%s/records", c.url, path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	if err != nil {
		return respDecoded, err
	}
	header, err := c.createAuthHeader(request)
	if err != nil {
		return respDecoded, err
	}
	req.Header.Add("X-Polybase-Signature", header)
	resp, err := client.Do(req)
	if err != nil {
		return respDecoded, err
	}
	err = json.NewDecoder(resp.Body).Decode(&respDecoded)
	return respDecoded, err

}

func NewPolybaseClient(privateKey, namespace, url string) (*PolybaseClient, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return &PolybaseClient{}, err
	}
	return &PolybaseClient{
		privateKey: key,
		namespace:  namespace,
		url:        url,
	}, nil
}
