package client

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func (client *BlockScoutAPIClient) GetEthBalance(address string, block uint64) (*big.Float, error) {

	url := client.setAccountApiUrl("eth_get_balance") + "&address=" + address +
		"&block=" + strconv.FormatUint(block, 10)

	var response common.Response
	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error == "Balance not found" {
		return nil, common.ErrBalanceNotFound
	}

	ethBalance, err := common.HexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	balance, err := common.WeiToEther(ethBalance)

	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (client *BlockScoutAPIClient) GetBalance(address string, block uint64) (*big.Float, error) {

	url := client.setAccountApiUrl("eth_get_balance") +
		"&address=" + address +
		"&block=" + strconv.FormatUint(uint64(block), 10)

	var response common.Response
	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error == "Balance not found" {
		return nil, err
	}
	// TO DO
	ethBalance, err := common.HexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}
	balance, err := common.WeiToEther(ethBalance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (client *BlockScoutAPIClient) GetBalanceMulti(addrs []string) error {

	url := client.setAccountApiUrl("balancemulti") + "&address="

	if len(addrs) < 20 {
		for i := 0; i < len(addrs); i++ {
			url += addrs[i] + ","
		}

		url = url[:len(url)-1]

		var result []common.BalanceMulti
		err := client.sendApiRpcRequestResult(url, &result)

		if err != nil {
			return err
		}

		// for _, res := range result {
		// 	fmt.Printf("Addr: %s, balance: %f\n",
		// 		res.Account, weiToEther(decStingToBigInt(res.Balance)))
		// }

	} else {
		fmt.Println("Max 20!")
	}
	return nil
}

func (client *BlockScoutAPIClient) GetTxList(address string) ([]common.Transaction, error) {

	url := client.setAccountApiUrl("txlist") + "&address=" + address

	var result []common.Transaction
	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BlockScoutAPIClient) GetTxListInternal(address string) ([]common.InternalTransaction, error) {

	url := client.setAccountApiUrl("txlistinternal") +
		"&address=" + address

	var result []common.InternalTransaction
	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BlockScoutAPIClient) GetTokenList(address string) ([]common.TokenInfo, error) {

	url := client.setAccountApiUrl("tokenlist") +
		"&address=" + address

	var result []common.TokenInfo
	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BlockScoutAPIClient) GetMinedBlocks(address string) ([]common.MinedBlock, error) {

	url := client.setAccountApiUrl("getminedblocks") +
		"&address=" + address

	var result []common.MinedBlock
	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BlockScoutAPIClient) GetTokenBalance(account string, tokenAddr string) (*big.Int, error) {

	url := client.setAccountApiUrl("tokenbalance") +
		"&contractaddress=" + tokenAddr +
		"&address=" + account

	var response common.Response
	err := client.sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	balance, err := common.HexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (client *BlockScoutAPIClient) setAccountApiUrl(actions string) string {
	url := client.URL + moduleValues["account"] +
		accountActions[actions]
	return url
}
