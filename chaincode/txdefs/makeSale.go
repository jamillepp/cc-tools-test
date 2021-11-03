package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

func getWhereIsAvailable(sw *stubwrapper.StubWrapper, product assets.Key) (map[string]interface{}, error) {
	query := fmt.Sprintf(`{
		"selector": {
			"storage": {
				"$elemMatch": {
					"@key": "%s"
				}
			}
		}
		}`, product.Key())

	iterator, err_ := sw.GetQueryResult(query)
	if err_ != nil {
		return nil, errors.WrapErrorWithStatus(err_, "error getting query result", 500)
	}

	res, err := iterator.Next()
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error iterating query response", 500)
	}
	response := make(map[string]interface{})
	err = json.Unmarshal(res.Value, &response)
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error during unmarshal of response", 500)
	}

	return response, nil
}

var MakeSale = tx.Transaction{
	Tag:         "makeSale",
	Label:       "Make a product sale",
	Description: "Make a product sale",
	Method:      "POST",
	Callers:     []string{"org1MSP", "org2MSP"},

	Args: tx.ArgList{
		{
			Tag:         "product",
			Label:       "Product",
			Description: "Product to be purchased",
			DataType:    "->product",
			Required:    true,
		},
		{
			Tag:         "price",
			Label:       "Price",
			Description: "Product's price",
			DataType:    "number",
			Required:    true,
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		product, _ := req["product"].(assets.Key)
		price, _ := req["price"].(float64)

		store, err := getWhereIsAvailable(sw, product)
		if err != nil {
			return nil, errors.WrapError(err, "couldn't get store where is available")
		}

		productMap := map[string]interface{}{
			"@assetType": "product",
			"@key":       product.Key(),
		}

		saleMap := make(map[string]interface{})
		saleMap["@assetType"] = "sale"
		saleMap["product"] = productMap
		saleMap["price"] = price
		saleMap["store"] = store

		saleAsset, err := assets.NewAsset(saleMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new saleAsset")
		}

		_, err = saleAsset.PutNew(sw)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving saleAsset on blockchain")
		}

		responseJSON, nerr := json.Marshal(saleAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return responseJSON, nil
	},
}
