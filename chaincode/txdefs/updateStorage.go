package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Updates the storage
// POST Method
var UpdateStorage = tx.Transaction{
	Tag:         "updateStorage",
	Label:       "Update storage",
	Description: "Update the storage of some store",
	Method:      "POST",
	Callers:     []string{`$org1MSP`},

	Args: []tx.Argument{
		{
			Tag:      "store",
			Label:    "Store",
			DataType: "->store",
			Required: true,
		},
		{
			Tag:      "productName",
			Label:    "Product name",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "productLot",
			Label:    "Product Lot",
			DataType: "string",
			Required: true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		storeKey, ok := req["store"].(assets.Key)
		productName, _ := req["productName"].(string)
		productLot, _ := req["productLot"].(string)

		if !ok {
			return nil, errors.WrapError(nil, "Parameter store must be an asset")
		}

		productMap := make(map[string]interface{})
		productMap["@assetType"] = "product"
		productMap["productName"] = productName
		productMap["productLot"] = productLot

		productAsset, err := assets.NewAsset(productMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = productAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		storeAsset, err := storeKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		storage, ok := storeAsset.GetProp("storage").([]interface{})
		if !ok {
			storage = make([]interface{}, 0)
		}

		storage = append(storage, productAsset)
		fmt.Println(stub.WriteSet)
		_, err = storeAsset.Update(stub, map[string]interface{}{
			"storage": storage,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to update storage from store asset")
		}

		// Marshal asset back to JSON format
		responseJSON, nerr := json.Marshal(storage)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return responseJSON, nil
	},
}
