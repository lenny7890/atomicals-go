package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/atomicals-go/lib/api"
	"github.com/atomicals-go/lib/atomicals"
	"github.com/atomicals-go/lib/validate"
)

// MintDft is mint-dft cmd
var MintDft = &cobra.Command{
	Use:  "mint-dft",
	RunE: runMintDft,
}

func runMintDft(cmd *cobra.Command, args []string) error {
	// TODO do mint
	dir, _ := os.Getwd()
	wallet, err := validate.WalletJsonFileReader(fmt.Sprintf("%s/wallets/wallet.json", dir))
	if err != nil {
		return err
	}

	isRBF := true
	ticker := "atom"
	satsbyte := 150

	endpoints := atomicals.Config.GetStringSlice("ELECTRUMX_PROXY_BASE_URL")
	atomical := atomicals.Atomicals{ElectrumAPI: api.NewElectrumAPI(endpoints)}

	return atomical.MintDft(isRBF, satsbyte, wallet.Primary.Address, ticker, wallet.Funding.WIF)
}
