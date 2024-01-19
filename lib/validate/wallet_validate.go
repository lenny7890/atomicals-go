package validate

import (
	"encoding/json"
	"os"
)

type Wallet struct {
	Phrase   string         `json:"phrase"`
	Primary  Primary        `json:"primary"`
	Funding  Funding        `json:"funding"`
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

// JsonFileReader get wallets/wallet.json data
func (w *Wallet) JsonFileReader(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var wallet Wallet

	err = json.Unmarshal(content, &wallet)
	if err != nil {
		return err
	}
	return nil
}
