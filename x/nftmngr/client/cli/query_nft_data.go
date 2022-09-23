package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func CmdListNftData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-data",
		Short: "list all NftData",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNftDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NftDataAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowNftData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-data [nft-schema-code] [token-id]",
		Short: "shows a NftData",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]
			argTokenId := args[1]

			params := &types.QueryGetNftDataRequest{
				NftSchemaCode: argNftSchemaCode,
				TokenId:       argTokenId,
			}

			res, err := queryClient.NftData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
