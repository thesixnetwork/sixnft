package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateSyncActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-sync-action-signer [actor-address] [owner-address]",
		Short: "Requesting to sync actionSigner with smartcontract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argActorAddress := args[0]
			argOwnerAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSyncActionSigner(
				clientCtx.GetFromAddress().String(),
				argActorAddress,
				argOwnerAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
