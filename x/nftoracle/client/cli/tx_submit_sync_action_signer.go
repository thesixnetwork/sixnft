package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdSubmitSyncActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-sync-action-signer [sync-id] [actor-address] [owner-address] [expire-at]",
		Short: "Submit Verified actionSigner from smartcontract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSyncId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argActorAddress := args[1]
			argOwnerAddress := args[2]
			argExpireAt := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitSyncActionSigner(
				clientCtx.GetFromAddress().String(),
				argSyncId,
				argActorAddress,
				argOwnerAddress,
				argExpireAt,
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
