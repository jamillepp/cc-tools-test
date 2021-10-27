package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Store on channel
// POST Method
var CreateNewStore = tx.Transaction{
	Tag:         "createNewStore",
	Label:       "Create New Store",
	Description: "Create a New Store",
	Method:      "POST",
	Callers:     []string{`$org\dMSP`}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "storeName",
			Label:       "Store name",
			Description: "Name of the Store",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "owner",
			Label:       "Owner",
			Description: "Store's Owner",
			DataType:    "->person",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["storeName"].(string)
		ownerKey, ok := req["owner"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter owner must be an asset")
		}

		storeMap := make(map[string]interface{})

		ownerAsset, err := ownerKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		ownerMap := (map[string]interface{})(*ownerAsset)

		updatedOwnerKey := make(map[string]interface{})
		updatedOwnerKey["@assetType"] = "person"
		updatedOwnerKey["@key"] = ownerMap["@key"]

		storeMap["@assetType"] = "store"
		storeMap["storeName"] = name
		storeMap["owner"] = updatedOwnerKey

		storeAsset, err := assets.NewAsset(storeMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = storeAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		storeJSON, nerr := json.Marshal(storeAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return storeJSON, nil
	},
}
