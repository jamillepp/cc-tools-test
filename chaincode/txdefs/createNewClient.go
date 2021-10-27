package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Product on channel
// POST Method
var CreateNewClient = tx.Transaction{
	Tag:         "createNewClient",
	Label:       "Create New Client",
	Description: "Create a New Client",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf",
			Required: true,
		},
		{
			Tag:         "clientName",
			Label:       "Client name",
			Description: "Name of the Client",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)

		clientMap := make(map[string]interface{})
		clientMap["@assetType"] = "person"
		clientMap["id"] = id

		clientAsset, err := assets.NewAsset(clientMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new library on channel
		_, err = clientAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		productJSON, nerr := json.Marshal(clientAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return productJSON, nil
	},
}
