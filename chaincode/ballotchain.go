package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
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

//Voter structure for storing votes
type Voter struct {
	ID          string `json:"id"`
	CandidateID string `json:"candidateID"`
	Time        string `json:"time"`
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

//GetCandidate retrieve data for a candidate
func (s *SmartContract) GetCandidate(ctx contractapi.TransactionContextInterface,
	candidateID string) (Candidate, error) {

	bytes, err := ctx.GetStub().GetState(candidateID)
	candidate := Candidate{}

	if err != nil {
		return candidate, err
	}

	err = json.Unmarshal(bytes, &candidate)

	return candidate, err
}

//GetAllCandidates retrieve fata for all candidates for an election
func (s *SmartContract) GetAllCandidates(ctx contractapi.TransactionContextInterface,
	electionID string) ([]Candidate, error) {

	candidates := make([]Candidate, 0)
	bytes, err := ctx.GetStub().GetState(electionID)
	if err != nil {
		return candidates, err
	}

	e := Election{}
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
