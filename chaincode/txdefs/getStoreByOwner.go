package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

func getStores(sw *stubwrapper.StubWrapper, person assets.Key) ([]interface{}, error) {
	query := fmt.Sprintf(`{
		"selector": {
			"@assetType": "store",
			"owner.@key": "%s"
		}
	}`, person.Key())

	iterator, err := sw.GetQueryResult(query)
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error getting query result", 500)
	}

	var response []interface{}

	for iterator.HasNext() {
		res, err := iterator.Next()
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error iterating response", 500)
		}

		owners := make(map[string]interface{})

		err = json.Unmarshal(res.Value, &owners)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error getting query result", 500)
		}
		response = append(response, owners)

	}
	return response, nil
}

var GetStoreByOwner = tx.Transaction{
	Tag:         "getStoreByOwner",
	Label:       "Get stores by owner",
	Description: "Get stores br owner",
	ReadOnly:    true,

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "person",
			Label:    "person",
			DataType: "->person",
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		person, _ := req["person"].(assets.Key)

		storesRes, err := getStores(sw, person)
		if err != nil {
			return nil, errors.WrapError(err, "Couldn't get store by owner")
		}

		responseJSON, nerr := json.Marshal(storesRes)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return responseJSON, nil
	},
}
