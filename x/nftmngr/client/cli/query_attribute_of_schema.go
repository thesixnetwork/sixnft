package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func CmdShowAttributeOfSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-attribute-of-schema [nft-schema-code]",
		Short: "shows a attribute_of_schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]

			params := &types.QueryGetAttributeOfSchemaRequest{
				NftSchemaCode: argNftSchemaCode,
			}

			res, err := queryClient.AttributeOfSchema(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
