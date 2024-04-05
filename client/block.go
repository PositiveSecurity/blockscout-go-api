package client

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func (client *BlockScoutAPIClient) GetCurrentBlockRpcApi() (*big.Int, error) {
	url := client.setBlockApiUrl("eth_block_number")

	var response common.Response

	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	currentBlock, err := common.HexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return currentBlock, nil
}

func (client *BlockScoutAPIClient) GetBlockRewardRpcApi(block uint64) (*common.BlockRewardInfo, error) {

	url := client.setBlockApiUrl("getblockreward") + "&blockno=" + strconv.FormatUint(block, 10)

	var result common.BlockRewardInfo

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) GetBlockByTime(timestamp uint64, closest string) (*common.BlockNumber, error) {

	url := client.setBlockApiUrl("getblocknobytime") + "&timestamp=" + strconv.FormatUint(timestamp, 10) +
		"&closest=" + closest

	var result common.BlockNumber

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) setBlockApiUrl(actions string) string {
	url := client.URL + moduleValues["block"] +
		blockActions[actions]
	return url
}

func (client *BlockScoutAPIClient) GetBlockInfo(block_number_or_hash string) ([]byte, error) {
	url := client.URLv2 + "blocks/" + block_number_or_hash
	body, err := client.sendV2ApiRequest(url)

	if err != nil {
		return nil, err
	}

	return body, nil

}
