package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdListAttributeBySchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-attribute-by-schema [nft-schema-code]",
		Short: "Query listAttributeBySchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqNftSchemaCode := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListAttributeBySchemaRequest{

				NftSchemaCode: reqNftSchemaCode,
			}

			res, err := queryClient.ListAttributeBySchema(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
