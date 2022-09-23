package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

func CmdListAddressBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-address-binding",
		Short: "list all AddressBinding",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAddressBindingRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AddressBindingAll(context.Background(), params)
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

func CmdShowAddressBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-address-binding [eth-address] [native-address]",
		Short: "shows a AddressBinding",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argEthAddress := args[0]
			// argNativeAddress := args[1]

			params := &types.QueryGetAddressBindingRequest{
				EthAddress: argEthAddress,
				// NativeAddress: argNativeAddress,
			}

			res, err := queryClient.AddressBinding(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
