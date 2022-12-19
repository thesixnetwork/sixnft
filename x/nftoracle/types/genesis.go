package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MintRequestList:            []MintRequest{},
		ActionRequestList:          []ActionOracleRequest{},
		CollectionOwnerRequestList: []CollectionOwnerRequest{},
		OracleConfig:               nil,
		ActionSignerList:           []ActionSigner{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in mintRequest
	mintRequestIdMap := make(map[uint64]bool)
	mintRequestCount := gs.GetMintRequestCount()
	for _, elem := range gs.MintRequestList {
		if _, ok := mintRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for mintRequest")
		}
		if elem.Id >= mintRequestCount {
			return fmt.Errorf("mintRequest id should be lower or equal than the last id")
		}
		mintRequestIdMap[elem.Id] = true
	}
	// Check for duplicated ID in actionRequest
	actionRequestIdMap := make(map[uint64]bool)
	actionRequestCount := gs.GetActionRequestCount()
	for _, elem := range gs.ActionRequestList {
		if _, ok := actionRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for actionRequest")
		}
		if elem.Id >= actionRequestCount {
			return fmt.Errorf("actionRequest id should be lower or equal than the last id")
		}
		actionRequestIdMap[elem.Id] = true
	}
	// Check for duplicated ID in collectionOwnerRequest
	collectionOwnerRequestIdMap := make(map[uint64]bool)
	collectionOwnerRequestCount := gs.GetCollectionOwnerRequestCount()
	for _, elem := range gs.CollectionOwnerRequestList {
		if _, ok := collectionOwnerRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for collectionOwnerRequest")
		}
		if elem.Id >= collectionOwnerRequestCount {
			return fmt.Errorf("collectionOwnerRequest id should be lower or equal than the last id")
		}
		collectionOwnerRequestIdMap[elem.Id] = true
	}
	// Check for duplicated index in actionSigner
	actionSignerIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionSignerList {
		index := string(ActionSignerKey(elem.ActorAddress))
		if _, ok := actionSignerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionSigner")
		}
		actionSignerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
