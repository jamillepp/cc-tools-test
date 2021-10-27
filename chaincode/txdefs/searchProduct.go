package txdefs

import (
	"encoding/json"

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
		productKey, ok := req["product"].(assets.Key)

		if !ok {
			return nil, errors.WrapError(nil, "Parameter store must be an asset")
		}

		storeAsset, err := storeKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		// Marshal asset back to JSON format
		storeJSON, nerr := json.Marshal(storeAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return storeJSON, nil
	},
}
