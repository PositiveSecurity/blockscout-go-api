package client

// https://explorer.celo.org/mainnet/

import (
	"fmt"
	"strconv"

	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func (client *BlockScoutAPIClient) GetEpoch(epoch uint64) (common.Epoch, error) {
	url := client.URL +
		moduleValues["epoch"] +
		epochActions["getepoch"] +
		"&epochNumber=" + strconv.FormatUint(epoch, 10)

	var result common.Epoch
	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return common.Epoch{}, err
	}

	fmt.Println(result.BlockHash)

	return result, nil
}
