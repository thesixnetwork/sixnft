package cli

import (
	"strconv"

	// "encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdAddAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-attribute [code] [new-attibute]",
		Short: "Broadcast message addAttribute",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCode := args[0]
			argNewAttibute := args[1]
			// err = json.Unmarshal([]byte(args[1]), argNewAttibute)
			// if err != nil {
			// 	return err
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAttribute(
				clientCtx.GetFromAddress().String(),
				argCode,
				argNewAttibute,
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
