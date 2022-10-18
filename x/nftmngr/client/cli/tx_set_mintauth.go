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

func CmdSetMintauth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-mintauth [nft-schema-code] [authorize-to]",
		Short: "To resync action to minted metadata",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argAuthorizeTo:= args[1]
			if err != nil {
				return err
			}

			argAuthorizeToInt, err := strconv.Atoi(argAuthorizeTo)
			if err != nil {
				return err
			}
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			authrizeTo := types.AuthorizeTo_SYSTEM // 0
			if argAuthorizeToInt == 1 {
				authrizeTo = types.AuthorizeTo_ALL
			}
			msg := types.NewMsgSetMintauth(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				authrizeTo,
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
