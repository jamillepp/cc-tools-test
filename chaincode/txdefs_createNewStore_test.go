package main_test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	cc "github.com/goledgerdev/cc-tools-demo/chaincode"
	"github.com/goledgerdev/cc-tools/mock"
)

func TestCreateNewStore(t *testing.T) {
	stub := mock.NewMockStub("org1MSP", new(cc.CCDemo))

	expectedResponse := map[string]interface{}{
		"@assetType":   "store",
		"@key":         "store:7b2334c5-a19f-59e4-8cb4-a9fb9e85329f",
		"@lastTouchBy": "org1MSP",
		"@lastTx":      "createNewStore",
		"owner": map[string]interface{}{
			"@assetType": "person",
			"@key":       "person:820ae33f-37d1-5771-9ded-4aa4b5380752",
		},
		"storeName": "Michael's",
	}
	req := map[string]interface{}{
		"storeName": "Michael's",
		"owner": map[string]interface{}{
			"@assetType": "person",
			"@key":       "person:820ae33f-37d1-5771-9ded-4aa4b5380752",
		},
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		t.FailNow()
	}

	res := stub.MockInvoke("createNewStore", [][]byte{
		[]byte("createNewStore"),
		reqBytes,
	})

	if res.GetStatus() != 200 {
		log.Println(res)
		t.FailNow()
	}

	var resPayload map[string]interface{}
	err = json.Unmarshal(res.GetPayload(), &resPayload)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(resPayload, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", resPayload)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}

	var state map[string]interface{}
	stateBytes := stub.State["store:7b2334c5-a19f-59e4-8cb4-a9fb9e85329f"]
	err = json.Unmarshal(stateBytes, &state)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(state, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", state)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}
}
