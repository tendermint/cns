package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/cns/x/cns/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(RegisterChainNameCmd())
	cmd.AddCommand(UpdateChainInfoCmd())
	return cmd
}

func RegisterChainNameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Args:  cobra.ExactArgs(1),
		Short: "register chain name",
		Long: `register chain name :

$ <appd> tx cns register <chain-name> flags --from owner
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ibcClient, _ := cmd.Flags().GetString(FlagCanonicalIbcClient)
			sourceCodeUrl, _ := cmd.Flags().GetString(FlagSourceCodeUrl)
			version, _ := cmd.Flags().GetInt64(FlagVersion)
			commitHash, _ := cmd.Flags().GetString(FlagCommitHash)
			genesisHash, _ := cmd.Flags().GetString(FlagGenesisHash)
			seeds, _ := cmd.Flags().GetString(FlagSeeds)
			owner := clientCtx.GetFromAddress()

			msg := types.NewMsgRegisterChainName(args[0], owner.String(), seeds, sourceCodeUrl,
				ibcClient, commitHash, genesisHash, version)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	AddCnsTxFlags(cmd)
	return cmd
}

func UpdateChainInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Args:  cobra.ExactArgs(1),
		Short: "update chain info",
		Long: `update chain info :

$ <appd> tx cns update <chain-name> flags --from owner
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ibcClient, _ := cmd.Flags().GetString(FlagCanonicalIbcClient)
			sourceCodeUrl, _ := cmd.Flags().GetString(FlagSourceCodeUrl)
			version, _ := cmd.Flags().GetInt64(FlagVersion)
			commitHash, _ := cmd.Flags().GetString(FlagCommitHash)
			genesisHash, _ := cmd.Flags().GetString(FlagGenesisHash)
			seeds, _ := cmd.Flags().GetString(FlagSeeds)
			owner := clientCtx.GetFromAddress()

			msg := types.NewMsgUpdateChainInfo(args[0], owner.String(), seeds, sourceCodeUrl,
				ibcClient, commitHash, genesisHash, version)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	AddCnsTxFlags(cmd)

	return cmd
}
