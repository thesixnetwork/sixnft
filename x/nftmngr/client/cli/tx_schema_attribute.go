package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func CmdCreateSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-schema-attribute [nft-schema-code] [name] [base64NewAttriuteDefenition]",
		Short: "Create a new schema_attribute",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNftSchemaCode := args[0]
			indexName := args[1]
			// Get value arguments
			argBase64NewAttriuteDefenition := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSchemaAttribute(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexName,
				argBase64NewAttriuteDefenition,
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

func CmdUpdateSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-schema-attribute [nft-schema-code] [name] [base66UpdateAttriuteDefenition]",
		Short: "Update a schema_attribute",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNftSchemaCode := args[0]
			indexName := args[1]
			// Get value arguments
			argBase64NewAttriuteDefenition := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSchemaAttribute(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexName,
				argBase64NewAttriuteDefenition,
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

func CmdDeleteSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-schema-attribute [nft-schema-code] [name]",
		Short: "Delete a schema_attribute",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNftSchemaCode := args[0]
			indexName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSchemaAttribute(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexName,
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
