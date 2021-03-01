package cli

import "github.com/spf13/cobra"

const (
	FlagCanonicalIbcClient = "canonical-ibc-client"
	FlagSourceCodeUrl      = "code-url"
	FlagVersion            = "version"
	FlagSeeds              = "seeds"
	FlagGenesisHash        = "genesis-hash"
	FlagCommitHash         = "commit-hash"
)

func AddCnsTxFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagCanonicalIbcClient, "", "Canonical ibc client")
	cmd.Flags().String(FlagSourceCodeUrl, "", "Link to source code url")
	cmd.Flags().Int64(FlagVersion, 0, "Current version of the chain")
	cmd.Flags().String(FlagCommitHash, "", "Commit hash of source code")
	cmd.Flags().String(FlagGenesisHash, "", "Hash of genesis file")
	cmd.Flags().String(FlagSeeds, "", "Comma seperated seeds of the chain")
}
