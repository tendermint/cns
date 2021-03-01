package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"strings"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/cns/x/cns/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cns queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdQueryChainInfo())

	return cmd
}

// GetCmdQueryChainInfo implements the chaininfo query command.
func GetCmdQueryChainInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-info [chain-name]",
		Short: "Query chain info",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query information of a chain name.

Example:
$ %s query cns chain-info <chain-name>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryChainInfoRequest{ChainName: args[0]}
			res, err := queryClient.QueryChainInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Info)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
