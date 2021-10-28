package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Store = assets.AssetType{
	Tag:   "store",
	Label: "Store",

	Props: []assets.AssetProp{
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "storeName",
			Label:    "Store name",
			DataType: "string",
			// Validate funcion
			Validate: func(storeName interface{}) error {
				storeNameStr := storeName.(string)
				if storeNameStr == "" {
					return fmt.Errorf("storeName must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "owner",
			Label:    "Owner",
			DataType: "->person",
		},
		{
			// Asset reference list
			Tag:      "storage",
			Label:    "Storage",
			DataType: "[]->product",
		},
	},
}
