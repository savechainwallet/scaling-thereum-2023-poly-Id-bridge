package polybase

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestList(t *testing.T) {

	client, err := NewPolybaseClient("2699d4da7d06f21b83499b4ddf752ca1819785eda9aeee09d21bf6de43a9395f", "poly-id-bridge", "https://testnet.polybase.xyz")
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.NewString()
	createRequest := make([]interface{}, 0)
	createRequest = append(createRequest, id)
	createRequest = append(createRequest, uuid.NewString())
	createRequest = append(createRequest, []byte(""))
	createRequest = append(createRequest, true)
	createRequest = append(createRequest, []byte(""))

	resp, err := client.CreateRecord("Session", createRequest)

	if err != nil {
		t.Error(err)
	}
	fmt.Print(resp)

	records, err := client.LisRecords("Session")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(records)

}
