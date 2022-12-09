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

func CmdChageOrgOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chage-org-owner [org-name] [to-new-owner]",
		Short: "Change Organization Owner",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrgName := args[0]
			argToNewOwner := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChageOrgOwner(
				clientCtx.GetFromAddress().String(),
				argOrgName,
				argToNewOwner,
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
