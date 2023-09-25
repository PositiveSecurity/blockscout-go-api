package blockscout

import (
	"fmt"
	"math/big"
	"strconv"
)

func GetCurrentBlockRpcApi() (*big.Int, error) {
	url := blockScoutApiUrl +
		moduleValues["block"] +
		blockActions["eth_block_number"]

	var response Response

	err := sendApiRpcRequest(url, &response)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	currentBlock, err := hexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return currentBlock, nil
}

func GetBlockRewardRpcApi(block uint64) (*BlockRewardInfo, error) {
	url := blockScoutApiUrl +
		moduleValues["block"] +
		blockActions["getblockreward"] +
		"&blockno=" + strconv.FormatUint(block, 10)

	var result BlockRewardInfo

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetBlockByTime(timestamp uint64, closest string) (*BlockNumber, error) {

	url := blockScoutApiUrl +
		moduleValues["block"] +
		blockActions["getblocknobytime"] +
		"&timestamp=" + strconv.FormatUint(timestamp, 10) +
		"&closest=" + closest

	var result BlockNumber

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
