package atomicals

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/sirupsen/logrus"
)

func (a *Atomicals) MintDft(rbf bool, satsbyte int, primaryAddress string, ticker string, fundingWIf string) error {

	// TODO

	_, err := btcutil.DecodeWIF(fundingWIf)
	if err != nil {
		return err
	}

	// get ticker info
	atomicalIdResult, err := a.ElectrumAPI.AtomicalsGetByTicker(ticker)
	if err != nil {
		return err
	}
	logrus.Infof("get ticker: %+v", atomicalIdResult)
	atomicalID := atomicalIdResult.Get("result").Get("atomical_id").String()
	ftInfoResult, err := a.ElectrumAPI.AtomicalsGetFtInfo(atomicalID)
	if err != nil {
		return err
	}

	globalInfo := ftInfoResult.Get("global").Map()
	atomicalInfo := ftInfoResult.Get("result")

	atomicalDecorated := make(map[string]string)

	logrus.Info(globalInfo, atomicalInfo, atomicalDecorated)

	return nil
}
