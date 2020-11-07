package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
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

//NewElection adds a new election to the ledger
func (s *SmartContract) NewElection(ctx contractapi.TransactionContextInterface,
	electionID string, candidateID []string, candidateName []string) error {

	e := &types.Election{
		ID:         electionID,
		Candidates: make([]string, 0),
		Active:     true,
	}

	for idx := range candidateID {
		e.Candidates = append(e.Candidates, candidateID[idx])

		c := &types.Candidate{
			ID:         candidateID[idx],
			Name:       candidateName[idx],
			ElectionID: electionID,
			Votes:      0,
		}

		bytes, _ := json.Marshal(c)
		if err := ctx.GetStub().PutState(c.ID, bytes); err != nil {
			return err
		}
	}

	bytes, _ := json.Marshal(e)
	if err := ctx.GetStub().PutState(e.ID, bytes); err != nil {
		return err
	}

	return nil
}

//GetCandidate retrieve data for a candidate
func (s *SmartContract) GetCandidate(ctx contractapi.TransactionContextInterface,
	candidateID string) (types.Candidate, error) {

	bytes, err := ctx.GetStub().GetState(candidateID)
	candidate := types.Candidate{}

	if err != nil {
		return candidate, err
	}

	err = json.Unmarshal(bytes, &candidate)

	return candidate, err
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
		c, _ := s.GetCandidate(ctx, val.ID)
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
	err = ctx.GetStub().PutState(v.ID, bytes)

	if err != nil {
		return err
	}

	candidate.Votes++
	bytes, _ = json.Marshal(candidate)
	return ctx.GetStub().PutState(v.CandidateID, bytes)
}

func main() {
	config := ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting unblockballot chaincode: %s", err.Error())
	}
}
