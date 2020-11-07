package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	types "github.com/raydwaipayan/unblockballot-server/types"
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

// InitLedger adds a base state
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	e := types.Election{
		ID: "1",
	}
	bytes, _ := json.Marshal(e)
	return ctx.GetStub().PutState(e.ID, bytes)
}

//NewElection adds a new election to the ledger
func (s *SmartContract) NewElection(ctx contractapi.TransactionContextInterface,
	electionID string, candidateJSON string) error {

	e := &types.Election{
		ID:         electionID,
		Candidates: make([]string, 0),
		Active:     true,
	}

	var candidates []types.Candidate
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

//GetElection retrieves an election from the ledger
func (s *SmartContract) GetElection(ctx contractapi.TransactionContextInterface,
	electionID string) (types.Election, error) {

	e := types.Election{}

	result, err := ctx.GetStub().GetState(electionID)

	if err != nil {
		return e, errors.New("Election not found")
	}

	err = json.Unmarshal(result, &e)
	return e, err
}

//GetCandidate retrieve data for a candidate
func (s *SmartContract) GetCandidate(ctx contractapi.TransactionContextInterface,
	candidateID string) (types.Candidate, error) {

	bytes, err := ctx.GetStub().GetState(candidateID)
	candidate := types.Candidate{}

	if err != nil {
		return candidate, errors.New("Candidate not found")
	}

	err = json.Unmarshal(bytes, &candidate)

	return candidate, errors.New("Error parsing stored data: " + string(bytes))
}

//GetAllCandidates retrieve fata for all candidates for an election
func (s *SmartContract) GetAllCandidates(ctx contractapi.TransactionContextInterface,
	electionID string) ([]types.Candidate, error) {

	candidates := make([]types.Candidate, 0)
	bytes, err := ctx.GetStub().GetState(electionID)
	if err != nil {
		return candidates, err
	}

	e := types.Election{}
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		return candidates, err
	}

	for _, val := range e.Candidates {
		c, _ := s.GetCandidate(ctx, val)
		candidates = append(candidates, c)
	}

	return candidates, nil
}

//AddVote Add a new vote to ledger
func (s *SmartContract) AddVote(ctx contractapi.TransactionContextInterface,
	voterID string, candidateID string, time string) error {

	v := types.Vote{
		VoterID:     voterID,
		CandidateID: candidateID,
		Time:        time,
	}

	bytes, err := ctx.GetStub().GetState(v.CandidateID)
	if err != nil {
		return errors.New("Candidate not found")
	}

	candidate := types.Candidate{}
	err = json.Unmarshal(bytes, &candidate)
	if err != nil {
		return err
	}

	if _, err := ctx.GetStub().GetState(v.VoterID); err == nil {
		return errors.New("User has already voted")
	}

	bytes, _ = json.Marshal(v)
	err = ctx.GetStub().PutState(v.VoterID, bytes)

	if err != nil {
		return err
	}

	candidate.Votes++
	bytes, _ = json.Marshal(candidate)
	return ctx.GetStub().PutState(v.CandidateID, bytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create ballotchain chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting ballotchain chaincode: %s", err.Error())
	}
}
