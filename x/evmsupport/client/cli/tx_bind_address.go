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

func CmdBindAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bind-address [eth-address] [signature] [signed-message]",
		Short: "To bind EVM address with native address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEthAddress := args[0]
			argSignature := args[1]
			argSignedMessage := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBindAddress(
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
