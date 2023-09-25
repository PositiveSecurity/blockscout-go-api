package blockscout

// https://explorer.celo.org/mainnet/

import (
	"fmt"
	"strconv"
)

func GetEpoch(epoch uint64) (Epoch, error) {
	url := blockScoutApiUrl +
		moduleValues["epoch"] +
		epochActions["getepoch"] +
		"&epochNumber=" + strconv.FormatUint(epoch, 10)

	var result Epoch
	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return Epoch{}, err
	}

	fmt.Println(result.BlockHash)

	return result, nil
}
