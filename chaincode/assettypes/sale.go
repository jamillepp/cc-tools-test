package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Sale = assets.AssetType{
	Tag:         "sale",
	Label:       "Sale",
	Description: "Sale of products",
	Readers:     []string{"org1MSP", "org2MSP"},

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "product",
			Label:    "Product",
			DataType: "->product",
			Writers:  []string{"org1MSP", "org2MSP"},
		},
		{
			Required: true,
			Tag:      "price",
			Label:    "Price",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "store",
			Label:    "store",
			DataType: "->store",
		},
	},
}
