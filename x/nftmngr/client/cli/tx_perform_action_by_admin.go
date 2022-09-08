package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdPerformActionByAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perform-action-by-admin [nft-schema-code] [token-id] [action]",
		Short: "To do action",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argTokenId := args[1]
			argAction := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPerformActionByAdmin(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argTokenId,
				argAction,
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
