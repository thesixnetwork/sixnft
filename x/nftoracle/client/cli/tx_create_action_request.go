package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sixnft/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateActionRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action-request [vm] [base-64-action-signature]",
		Short: "To create Action Request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVm := args[0]
			argBase64ActionSignature := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionRequest(
				clientCtx.GetFromAddress().String(),
				argVm,
				argBase64ActionSignature,
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
