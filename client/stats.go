package client

import (
	"math/big"

	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func (client *BlockScoutAPIClient) GetTokenTotalSupply(addr string) (*big.Int, error) {

	url := client.setStatsApiUrl("tokensupply") + "&contractaddress=" + addr

	var response common.Response
	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := common.DecStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *BlockScoutAPIClient) GetTotalSupplyNativeCoin() (*big.Float, error) {

	url := client.setStatsApiUrl("ethsupplyexchange")

	var response common.Response
	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := common.DecStringToBigFloat(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *BlockScoutAPIClient) GetCoinPriceUSD() (*common.CoinPrice, error) {
	url := client.setStatsApiUrl("coinprice")

	var result common.CoinPrice

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) setStatsApiUrl(actions string) string {
	url := client.URL + moduleValues["stats"] +
		statsActions[actions]
	return url
}
