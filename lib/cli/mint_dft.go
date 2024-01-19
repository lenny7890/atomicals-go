package cli

import (
	"github.com/spf13/cobra"
)

// MintDft is mint-dft cmd
var MintDft = &cobra.Command{
	Use:  "mint-dft",
	RunE: runMintDft,
}

func runMintDft(cmd *cobra.Command, args []string) error {
	// TODO do mint
	return nil
}
