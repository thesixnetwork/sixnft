package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sixnft/x/evmsupport/types"
)

var _ = strconv.Itoa(0)

func CmdRemoveBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-binding [eth-address] [signature] [signed-message]",
		Short: "To remove binding address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEthAddress := args[0]
			argSignature := args[1]
			argSignedMessage := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveBinding(
				clientCtx.GetFromAddress().String(),
				argEthAddress,
				argSignature,
				argSignedMessage,
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
