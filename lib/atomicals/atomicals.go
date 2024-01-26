package atomicals

import (
	"github.com/spf13/viper"

	"github.com/atomicals-go/lib/api"
)

type Atomicals struct {
	ElectrumAPI *api.ElectrumAPI `json:"electrum_api"`
}

var Config *viper.Viper
