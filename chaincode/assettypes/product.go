package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Product = assets.AssetType{
	Tag:         "product",
	Label:       "Product",
	Description: "Product from the store.",

	Props: []assets.AssetProp{
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "productName",
			Label:    "Product name",
			DataType: "string",
			// Validate funcion
			Validate: func(productName interface{}) error {
				nameStr := productName.(string)
				if nameStr == "" {
					return fmt.Errorf("productName must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "productLot",
			Label:    "Product lot",
			DataType: "string",
			Validate: func(productLot interface{}) error {
				productLotStr := productLot.(string)
				if productLotStr == "" {
					return fmt.Errorf("productLot must be non-empty")
				}
				if len(productLotStr) < 8 {
					return fmt.Errorf("invalid productLot: size must be a least 8")
				}
				return nil
			},
		},
	},
}
