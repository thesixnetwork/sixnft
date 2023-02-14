package cli

import (
	"context"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// will be deprecated next version (074)
func CmdListNFTSchemaV063() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-schemaV063",
		Short: "list all NFTSchema",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNFTSchemaRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NFTSchemaAllV063(context.Background(), params)
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

// will be deprecated next version (074)
func CmdShowNFTSchemaV063() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-schemaV063 [code]",
		Short: "shows a NFTSchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCode := args[0]

			params := &types.QueryGetNFTSchemaRequest{
				Code: argCode,
			}

			res, err := queryClient.NFTSchemaV063(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
