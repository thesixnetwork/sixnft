package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateNFTSchema())
	cmd.AddCommand(CmdCreateMetadata())
	cmd.AddCommand(CmdPerformActionByAdmin())
	cmd.AddCommand(CmdAddAttribute())
	cmd.AddCommand(CmdAddAction())
	cmd.AddCommand(CmdSetNFTAttribute())
	cmd.AddCommand(CmdSetBaseUri())
	cmd.AddCommand(CmdToggleAction())
	cmd.AddCommand(CmdChangeSchemaOwner())
	cmd.AddCommand(CmdAddSystemActioner())
	cmd.AddCommand(CmdRemoveSystemActioner())
	cmd.AddCommand(CmdResyncAttributes())
	cmd.AddCommand(CmdShowAttributes())
	cmd.AddCommand(CmdSetFeeConfig())
	cmd.AddCommand(CmdSetMintauth())
	// this line is used by starport scaffolding # 1

	return cmd
}
