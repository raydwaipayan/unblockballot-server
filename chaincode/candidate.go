package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Election contains election details
type Election struct {
	ID         string `json:"id"`
	Candidates []string
	Active     bool
}

//Candidate contains the candidature information
type Candidate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ElectionID string
	Votes      uint64
}

//NewElection adds a new election to the ledger
func (s *SmartContract) NewElection(ctx contractapi.TransactionContextInterface,
	electionID string, candidateJSON string) error {

	e := &Election{
		ID:         electionID,
		Candidates: make([]string, 0),
		Active:     true,
	}

	var candidates []Candidate
	json.Unmarshal([]byte(candidateJSON), &candidates)

	for _, val := range candidates {
		e.Candidates = append(e.Candidates, val.ID)
		val.Votes = 0
		val.ElectionID = e.ID

		bytes, _ := json.Marshal(val)
		if err := ctx.GetStub().PutState(val.ID, bytes); err != nil {
			return err
		}
	}

	bytes, _ := json.Marshal(e)
	if err := ctx.GetStub().PutState(e.ID, bytes); err != nil {
		return err
	}

	return nil
}
