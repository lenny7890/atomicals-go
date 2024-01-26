package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Phrase   string         `json:"phrase"`
	Primary  *Primary       `json:"primary"`
	Funding  *Funding       `json:"funding"`
	Imported map[string]any `json:"imported"`
}

type Primary struct {
	Address string `json:"address"`
	Path    string `json:"path"`
	WIF     string `json:"WIF"`
}

type Funding struct {
	Address string `json:"address"`
	Path    string `json:"path"`
	WIF     string `json:"WIF"`
}

// WalletJsonFileReader get wallets/wallet.json data
func WalletJsonFileReader(path string) (*Wallet, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wallet Wallet

	err = json.Unmarshal(content, &wallet)
	if err != nil {
		return nil, err
	}

	// check wallet param

	if wallet.Phrase == "" {
		return nil, errors.New(fmt.Sprintf(`phrase field not found in %s`, path))
	}

	// Validate is a valid mnemonic
	if !bip39.IsMnemonicValid(wallet.Phrase) {
		return nil, errors.New("phrase is not a valid mnemonic phrase!")
	}

	if wallet.Primary == nil {
		return nil, errors.New(`Wallet needs a primary address`)
	}

	if wallet.Funding == nil {
		return nil, errors.New(`Wallet needs a funding address`)
	}

	// Validate paths
	/*if (wallet.primary.path !== `m/44'/0'/0'/0/0`) {
	    console.log(`Primary path must be m/44'/0'/0'/0/0`);
	    throw new Error(`Primary path must be m/44'/0'/0'/0/0`);
	  }

	  if (wallet.funding.path !== `m/44'/0'/0'/1/0`) {
	    console.log(`Funding path must be m/44'/0'/0'/1/0`);
	    throw new Error(`Funding path must be m/44'/0'/0'/1/0`);
	  }*/

	// Validate WIF
	if wallet.Primary.WIF == "" {
		return nil, errors.New(`Primary WIF not set`)
	}
	if wallet.Funding.WIF == "" {
		return nil, errors.New(`Funding WIF not set`)
	}

	// Validate Addresses
	if wallet.Primary.Address == "" {
		return nil, errors.New(`primary address not set`)
	}
	if wallet.Funding.Address == "" {
		return nil, errors.New(`funding address not set`)
	}

	// seed := bip39.NewSeed(wallet.Phrase, "")
	// masterKey, _ := bip32.NewMasterKey(seed)
	// publicKey := masterKey.PublicKey()
	// masterKey, _ := bip32.NewMasterKey(seed)
	// derivePathPrimary := wallet.Primary.Path // `m/44'/0'/0'/0/0`;

	// Derive the primary and funding keys
	// primaryChild, _ := masterKey.NewChildKey(0)
	// fundingChild, _ := primaryChild.NewChildKey(0)

	// Get the public keys
	// primaryPubKey, _ := btcec.ParsePubKey(primaryChild.PublicKey().Key)
	// fundingPubKey, _ := btcec.ParsePubKey(fundingChild.PublicKey().Key)

	// Generate the p2tr addresses
	// primaryAddress, _ := btcutil.NewAddressPubKey(primaryPubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	// fundingAddress, _ := btcutil.NewAddressPubKey(fundingPubKey.SerializeCompressed(), &chaincfg.MainNetParams)

	return &wallet, nil
}
