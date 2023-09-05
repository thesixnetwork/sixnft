package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdListActionBySchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-by-schema [nft-schema-code]",
		Short: "Query listActionBySchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqNftSchemaCode := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListActionBySchemaRequest{

				NftSchemaCode: reqNftSchemaCode,
			}

			res, err := queryClient.ListActionBySchema(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
