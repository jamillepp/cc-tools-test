package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Product = assets.AssetType{
	Tag:         "product",
	Label:       "Product",
	Description: "Product from the shop.",

	Props: []assets.AssetProp{
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "productName",
			Label:    "Product name",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "productionBatch",
			Label:    "Production Batch",
			DataType: "integer",
		},
	},
}
