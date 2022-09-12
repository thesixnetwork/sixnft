package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MintRequestList: []MintRequest{},
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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
