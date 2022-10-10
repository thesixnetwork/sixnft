package cli

import (
	"strconv"

	// "encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdAddAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-attribute [code] [location] [new-attibute]",
		Short: "Broadcast message addAttribute",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCode := args[0]
			attributeLocation := args[1]
			argNewAttibute := args[2]
			// err = json.Unmarshal([]byte(args[1]), argNewAttibute)
			// if err != nil {
			// 	return err
			// }

			// convert attributeLocation to number
			attributeLocationInt, err := strconv.Atoi(attributeLocation)
			if err != nil {
				return err
			}

			attributeLo := types.AttributeLocation_NFT_ATTRIBUTE // 0

			if attributeLocationInt == 1 {
				attributeLo = types.AttributeLocation_TOKEN_ATTRIBUTE
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAttribute(
				clientCtx.GetFromAddress().String(),
				argCode,
				attributeLo,
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
