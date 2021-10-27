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
var UpdateStorage = tx.Transaction{
	Tag:         "updateStorage",
	Label:       "Update storage",
	Description: "Update the storage of some store",
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
			Tag:      "productName",
			Label:    "Product name",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "productionBatch",
			Label:    "Production Batch",
			DataType: "string",
			Required: true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		storeKey, ok := req["store"].(assets.Key)
		productName, _ := req["productName"].(string)
		productionBatch, _ := req["productionBatch"].(string)

		if !ok {
			return nil, errors.WrapError(nil, "Parameter store must be an asset")
		}

		productMap := make(map[string]interface{})
		productMap["@assetType"] = "product"
		productMap["productName"] = productName
		productMap["productionBatch"] = productionBatch

		productAsset, err := assets.NewAsset(productMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		storeAsset, err := storeKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		storage := storeAsset.GetProp("storage").([]assets.Key)
		storage = append(storage, assets.Key(productAsset))

		storeAsset.Update(stub, map[string]interface{}{
			"storage": storage,
		})

		// Marshal asset back to JSON format
		storeJSON, nerr := json.Marshal(storeAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return storeJSON, nil
	},
}
