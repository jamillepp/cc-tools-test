package main

import (
	txdefs "github.com/goledgerdev/cc-tools-demo/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.CreateNewPerson,
	txdefs.CreateNewStore,
	txdefs.SearchProduct,
	txdefs.UpdateStorage,
	txdefs.GetStoreByOwner,
	txdefs.MakeSale,
}
