package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// legacy query should be removed once all clients have been updated to use gRPC queries
func CmdListActionRequestV063() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-request-v063",
		Short: "list all ActionOracleRequest-v063",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActionRequestRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActionRequestAllV063(context.Background(), params)
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

func CmdShowActionRequestV063() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-request-v063 [id]",
		Short: "shows a ActionOracleRequest-v063",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetActionRequestRequest{
				Id: id,
			}

			res, err := queryClient.ActionOracleRequestV063(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
