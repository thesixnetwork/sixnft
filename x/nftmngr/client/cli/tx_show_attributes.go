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

func CmdShowAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-attributes [nft-schema-code] [attribute-name]",
		Short: "To show attribute from market place aka delete from hidden list",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argAttributeName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgShowAttributes(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argAttributeName,
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
