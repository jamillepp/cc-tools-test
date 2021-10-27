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
var SearchProduct = tx.Transaction{
	Tag:         "searchProduct",
	Label:       "Search Product",
	Description: "Search for some product on a store",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:      "store",
			Label:    "Store",
			DataType: "->store",
			Required: true,
		},
		{
			Tag:      "product",
			Label:    "product",
			DataType: "->product",
			Required: true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		storeKey, ok := req["store"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter store must be an asset")
		}
		productK, ok := req["product"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter product must be an asset")
		}

		storeMap, err := storeKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		productList, ok := storeMap["storage"].([]interface{})
		fmt.Println(storeMap)
		if !ok {
			return nil, errors.WrapError(err, "failed to get productlist")
		}
		response := make(map[string]interface{})
		for _, productInterface := range productList {
			productI, _ := productInterface.(map[string]interface{})

			productKey, err := assets.NewKey(productI)
			if err != nil {
				return nil, errors.WrapError(err, "Unable to create new Key from map to interface")
			}

			productMap, err := productKey.GetMap(stub)
			if err != nil {
				return nil, errors.WrapError(err, "failed to get asset from the ledger")
			}

			if productMap["@key"] == productK.Key() {
				fmt.Println(productMap)
				response["response"] = productMap
			}

		}

		// Marshal asset back to JSON format
		responseJSON, nerr := json.Marshal(response)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return responseJSON, nil
	},
}
