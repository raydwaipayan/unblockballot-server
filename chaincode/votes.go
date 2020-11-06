package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Voter structure for storing votes
type Voter struct {
	ID          string `json:"id"`
	CandidateID string `json:"candidateID"`
	Time        string `json:"time"`
}

//AddVote Add a new vote to ledger
func (s *SmartContract) AddVote(ctx contractapi.TransactionContextInterface,
	voterJSON string) error {

	v := Voter{}
	err := json.Unmarshal([]byte(voterJSON), &v)

	if err != nil {
		return err
	}

	bytes, err := ctx.GetStub().GetState(v.CandidateID)
	if err != nil {
		return errors.New("Candidate not found")
	}

	candidate := Candidate{}
	err = json.Unmarshal(bytes, &candidate)
	if err != nil {
		return err
	}

	if _, err := ctx.GetStub().GetState(v.ID); err == nil {
		return errors.New("User has already voted")
	}

	bytes, _ = json.Marshal(v)
	err = ctx.GetStub().PutState(v.ID, bytes)

	if err != nil {
		return err
	}

	candidate.Votes++
	bytes, _ = json.Marshal(candidate)
	return ctx.GetStub().PutState(v.CandidateID, bytes)
}
