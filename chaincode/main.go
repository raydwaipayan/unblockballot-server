package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//ServerConfig contains endpoint info
type ServerConfig struct {
	CCID    string
	Address string
}

//SmartContract contents the contract api methods
type SmartContract struct {
	contractapi.Contract
}
