package keeper_test

import (
	"fmt"
	"testing"

	_keeper "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	sim "github.com/thesixnetwork/sixnft/x/nftmngr/simulation"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
)

func TestAction(t *testing.T) {
	// Read in schema and metadata from JSON files
	schemaJSON, err := os.ReadFile("../simulation/schema.json")
	require.NoError(t, err)
	schemaInput := types.NFTSchemaINPUT{}
	err = jsonpb.UnmarshalString(string(schemaJSON), &schemaInput)
	require.NoError(t, err)

	metaJSON, err := os.ReadFile("../simulation/meta.json")
	require.NoError(t, err)

	metaInput := types.NftData{}
	err = jsonpb.UnmarshalString(string(metaJSON), &metaInput)
	require.NoError(t, err)

	// Simulate create metadata
	schema_, data_, globalAttributeSchema_ := sim.SimulateCreateMetadata(schemaInput, metaInput)
	

	// ** START: NOTHING TO CARE ABOUT THIS PART **

	keeper_, ctx := _keeper.NftmngrKeeper(t)
	keeper_.SetNFTSchema(ctx, schema_)

	if _, found := keeper_.GetNFTSchema(ctx, schema_.Code); !found {
		fmt.Println("Schema not found")
	} else {
		require.True(t, found)
		require.NoError(t, err)
		// require.Equal(t, schema_, val)
	}

	// Create metadata object with initial attributes
	keeper_.SetNftData(ctx, data_)

	if val, found := keeper_.GetNftData(ctx, schema_.Code, "1") ; !found {
		fmt.Println("Metadata not found")
	}else {
		require.True(t, found)
		require.NoError(t, err)
		require.Equal(t, data_, val)
	}
	// ** END: NOTHING TO CARE ABOUT THIS PART **


	// ** START: ACTUAL TEST **
	// Define the action to be processed
	selectAction := "check_in"

	// Create metadata object with initial attributes
	meta := types.NewMetadata(&schema_, &data_, types.AttributeOverriding_CHAIN, globalAttributeSchema_)

	// Process the action on the metadata
	actionParams_ := []*types.ActionParameter{}
	for _, action := range schema_.OnchainData.Actions {
		if action.Name == selectAction {
			keeper.ProcessAction(meta, action, actionParams_)
			break
		}
	}

	// Test 1: Check that the points attribute was updated correctly
	newPoints, err := meta.MustGetFloat("points")
	require.NoError(t, err)
	require.Equal(t, float64(50), newPoints)

	// Test 2: Check that the is_checked_in attribute was updated correctly
	isCheckIn, err := meta.MustGetBool("is_checked_in")
	require.NoError(t, err)
	require.True(t, isCheckIn)

	// Test 3: Check that the all_check_in attribute was updated correctly
	allCheckInNumber, err := meta.MustGetNumber("all_check_in")
	require.NoError(t, err)
	require.Equal(t, int64(1), allCheckInNumber)


	fmt.Printf("CASE 1 SUCCESS WITH UPDATE VALUE \n newPoints: %v\n isCheckIn :%v\n allCheckInNumber: %v\n", newPoints, isCheckIn, allCheckInNumber)
	// ** END: ACTUAL TEST **
}
