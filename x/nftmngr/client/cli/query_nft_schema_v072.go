package cli

import (
	"context"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListNFTSchemaV072() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-schemaV072",
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

			res, err := queryClient.NFTSchemaAllV072(context.Background(), params)
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

func CmdShowNFTSchemaV072() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-schemaV072 [code]",
		Short: "shows a NFTSchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCode := args[0]

			params := &types.QueryGetNFTSchemaRequest{
				Code: argCode,
			}

			res, err := queryClient.NFTSchemaV072(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
