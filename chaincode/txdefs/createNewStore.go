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
		owner, _ := req["owner"].(assets.Key)

		storeMap := make(map[string]interface{})
		storeMap["@assetType"] = "store"
		storeMap["storeName"] = name
		storeMap["owner"] = owner

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
