package blockscout

import (
	"math/big"
)

func GetTokenTotalSupply(addr string) (*big.Int, error) {
	url := blockScoutApiUrl +
		moduleValues["stats"] +
		statsActions["tokensupply"] +
		"&contractaddress=" + addr

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := decStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetTotalSupplyNativeCoin() (*big.Float, error) {
	url := blockScoutApiUrl +
		moduleValues["stats"] +
		statsActions["ethsupplyexchange"]

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	res, err := decStringToBigFloat(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetCoinPriceUSD() (*CoinPrice, error) {
	url := blockScoutApiUrl +
		moduleValues["stats"] +
		statsActions["coinprice"]

	var result CoinPrice

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
