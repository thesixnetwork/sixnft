package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdToggleAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toggle-action [code] [action] [disable]",
		Short: "Broadcast message toggleAction [true = disable, false = enable]",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCode := args[0]
			argAction := args[1]
			argDisable := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// string to bool
			disableBool, err := strconv.ParseBool(argDisable)
			if err != nil {
				return err
			}

			msg := types.NewMsgToggleAction(
				clientCtx.GetFromAddress().String(),
				argCode,
				argAction,
				disableBool,
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
