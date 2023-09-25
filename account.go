package blockscout

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

func GetEthBalance(address string, block uint64) (*big.Float, error) {

	url := blockScoutApiUrl + moduleValues["account"] +
		accountActions["eth_get_balance"] +
		"&address=" + address +
		"&block=" + strconv.FormatUint(block, 10)

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error == "Balance not found" {
		return nil, errors.New("Balance not found")
	}

	ethBalance, err := hexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	balance, err := weiToEther(ethBalance)

	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetBalance(address string, block uint64) (*big.Float, error) {

	url := blockScoutApiUrl + moduleValues["account"] +
		accountActions["eth_get_balance"] +
		"&address=" + address +
		"&block=" + strconv.FormatUint(uint64(block), 10)

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error == "Balance not found" {
		return nil, err
	}
	// TO DO
	ethBalance, err := hexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}
	balance, err := weiToEther(ethBalance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func GetBalanceMulti(addrs []string) error {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["balancemulti"] +
		"&address="

	if len(addrs) < 20 {
		for i := 0; i < len(addrs); i++ {
			url += addrs[i] + ","
		}

		url = url[:len(url)-1]

		var result []BalanceMulti
		err := sendApiRpcRequestResult(url, &result)

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

func GetTxList(address string) ([]Transaction, error) {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["txlist"] +
		"&address=" + address

	var result []Transaction
	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTxListInternal(address string) ([]InternalTransaction, error) {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["txlistinternal"] +
		"&address=" + address

	var result []InternalTransaction
	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTokenList(address string) ([]TokenInfo, error) {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["tokenlist"] +
		"&address=" + address

	var result []TokenInfo
	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetMinedBlocks(address string) ([]MinedBlock, error) {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["getminedblocks"] +
		"&address=" + address

	var result []MinedBlock
	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTokenBalance(account string, token string) (*big.Int, error) {
	url := blockScoutApiUrl +
		moduleValues["account"] +
		accountActions["tokenbalance"] +
		"&contractaddress=" + token +
		"&address=" + account

	var response Response
	err := sendApiRpcRequest(url, &response)

	if err != nil {
		return nil, err
	}

	balance, err := hexStringToBigInt(response.Result.(string))

	if err != nil {
		return nil, err
	}

	return balance, nil
}
