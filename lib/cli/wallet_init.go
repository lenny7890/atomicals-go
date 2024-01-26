package cli

import (
	"github.com/spf13/cobra"
)

// MintDft is mint-dft cmd
var WalletInit = &cobra.Command{
	Use:  "wallet-init",
	RunE: runWalletInit,
}

func runWalletInit(cmd *cobra.Command, args []string) error {

	return nil
}
