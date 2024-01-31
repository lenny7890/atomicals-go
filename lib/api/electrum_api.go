package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type ElectrumAPI struct {
	Endpoints []string `json:"endpoints"`
}

func NewElectrumAPI(endpoints []string) *ElectrumAPI {
	return &ElectrumAPI{
		Endpoints: endpoints,
	}
}

func (e *ElectrumAPI) call(requestVerb, method string, params any) (gjson.Result, error) {

	var value gjson.Result

	for _, baseUrl := range e.Endpoints {
		url := fmt.Sprintf("%s/%s", baseUrl, method)

		var err error
		var resp *grequests.Response
		options := grequests.RequestOptions{}

		if requestVerb == http.MethodPost {

			options.Headers = map[string]string{
				"Content-Type": "application/json",
			}

			options.JSON = params
			resp, err = grequests.Post(url, &options)
		} else {
			// string encoded array
			arrStr, _ := json.Marshal(params)
			options.Params = map[string]string{
				"params": string(arrStr),
			}
			resp, err = grequests.Get(url, &options)
		}

		if err != nil {
			logrus.Warnf("error creating request for endpoint %s: %v\n", baseUrl, err)
			continue
		}
		if !resp.Ok {
			logrus.Warnf("electrum proxy [%s] is %d", baseUrl, resp.StatusCode)
			continue
		}

		return gjson.ParseBytes(resp.Bytes()).Get("response"), nil
	}

	return value, errors.New("all endpoints failed")
}

func (e *ElectrumAPI) AtomicalsGetByTicker(ticker string) (result *ElectrumResponse[*GetByTickerResult], err error) {

	call, err := e.call(http.MethodGet, "blockchain.atomicals.get_by_ticker", []string{ticker})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(call.Raw), &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (e *ElectrumAPI) AtomicalsGetFtInfo(atomicalId string) (result *ElectrumResponse[*FtInfo], err error) {

	call, err := e.call(http.MethodGet, "blockchain.atomicals.get_ft_info", []string{atomicalId})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(call.Raw), &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
