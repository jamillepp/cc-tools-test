package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Person = assets.AssetType{
	Tag:   "person",
	Label: "Person",

	Props: []assets.AssetProp{
		{
			// Mandatory property
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf",
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
	},
}
