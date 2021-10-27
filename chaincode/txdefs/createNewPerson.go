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
var CreateNewPerson = tx.Transaction{
	Tag:         "createNewPerson",
	Label:       "Create New Person",
	Description: "Create a New Person",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf",
			Required: true,
		},
		{
			Tag:         "name",
			Label:       "Person name",
			Description: "Name of the Person",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)
		name, _ := req["name"].(string)

		personMap := make(map[string]interface{})
		personMap["@assetType"] = "person"
		personMap["id"] = id
		personMap["name"] = name

		personAsset, err := assets.NewAsset(personMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new library on channel
		_, err = personAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		personJSON, nerr := json.Marshal(personAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return personJSON, nil
	},
}
