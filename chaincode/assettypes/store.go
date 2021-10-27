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
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
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
