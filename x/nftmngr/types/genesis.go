package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NFTSchemaList:           []NFTSchema{},
		NftDataList:             []NftData{},
		ActionByRefIdList:       []ActionByRefId{},
		OrganizationList:        []Organization{},
		NFTSchemaByContractList: []NFTSchemaByContract{},
		NftFeeConfig:            nil,
		NFTFeeBalance:           nil,
		MetadataCreatorList:     []MetadataCreator{},
		NftCollectionList:       []NftCollection{},
		ActionExecutorList:      []ActionExecutor{},
		SchemaAttributeList:     []SchemaAttribute{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in nFTSchema
	nFTSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.NFTSchemaList {
		index := string(NFTSchemaKey(elem.Code))
		if _, ok := nFTSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nFTSchema")
		}
		nFTSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nftData
	nftDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.NftDataList {
		index := string(NftDataKey(elem.NftSchemaCode, elem.TokenId))
		if _, ok := nftDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nftData")
		}
		nftDataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionByRefId
	actionByRefIdIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionByRefIdList {
		index := string(ActionByRefIdKey(elem.RefId))
		if _, ok := actionByRefIdIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionByRefId")
		}
		actionByRefIdIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in organization
	organizationIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrganizationList {
		index := string(OrganizationKey(elem.Name))
		if _, ok := organizationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for organization")
		}
		organizationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nFTSchemaByContract
	nFTSchemaByContractIndexMap := make(map[string]struct{})

	for _, elem := range gs.NFTSchemaByContractList {
		index := string(NFTSchemaByContractKey(elem.OriginContractAddress))
		if _, ok := nFTSchemaByContractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nFTSchemaByContract")
		}
		nFTSchemaByContractIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in metadataCreator
	metadataCreatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.MetadataCreatorList {
		index := string(MetadataCreatorKey(elem.NftSchemaCode))
		if _, ok := metadataCreatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for metadataCreator")
		}
		metadataCreatorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nftCollection
	nftCollectionIndexMap := make(map[string]struct{})

	for _, elem := range gs.NftCollectionList {
		index := string(NftCollectionKey(elem.NftSchemaCode))
		if _, ok := nftCollectionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nftCollection")
		}
		nftCollectionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionExecutor
	actionExecutorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionExecutorList {
		index := string(ActionExecutorKey(elem.NftSchemaCode, elem.ExecutorAddress))
		if _, ok := actionExecutorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionExecutor")
		}
		actionExecutorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in schemaAttribute
	schemaAttributeIndexMap := make(map[string]struct{})

	for _, elem := range gs.SchemaAttributeList {
		index := string(SchemaAttributeKey(elem.NftSchemaCode, elem.Name))
		if _, ok := schemaAttributeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for schemaAttribute")
		}
		schemaAttributeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
