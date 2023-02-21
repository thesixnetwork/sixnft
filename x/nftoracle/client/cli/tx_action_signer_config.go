package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func CmdCreateActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action-signer-config [chain] [rpc-endpoint] [contract-address] [contract-name] [contract-owner]",
		Short: "Create a new actionSignerConfig",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChain := args[0]

			// Get value arguments
			argRpcEndpoint := args[1]
			argContractAddress := args[2]
			argContractName := args[3]
			argContractOwner := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
				argRpcEndpoint,
				argContractAddress,
				argContractName,
				argContractOwner,
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

func CmdUpdateActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action-signer-config [chain] [contract-address] [contract-name] [contract-owner]",
		Short: "Update a actionSignerConfig",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChain := args[0]

			// Get value arguments
			rpcEndpoint := args[1]
			argContractAddress := args[2]
			argContractName := args[3]
			argContractOwner := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
				rpcEndpoint,
				argContractAddress,
				argContractName,
				argContractOwner,
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

func CmdDeleteActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-action-signer-config [chain]",
		Short: "Delete a actionSignerConfig",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexChain := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
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
