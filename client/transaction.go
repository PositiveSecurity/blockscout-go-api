package client

import "github.com/PositiveSecurity/blockscout-go-api/common"

func (client *BlockScoutAPIClient) GetTxInfo(hash string) (*common.Transaction, error) {
	url := client.setTxApiUrl("gettxinfo", hash)

	var result common.Transaction

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) GetTxReceiptStatus(hash string) (*common.Status, error) {

	url := client.setTxApiUrl("gettxreceiptstatus", hash)

	var result common.Status

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) GetTxStatus(hash string) (*common.TxStatus, error) {

	url := client.setTxApiUrl("gettxinfo", hash)

	var result common.TxStatus

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) setTxApiUrl(actions, hash string) string {
	url := client.URL + moduleValues["transaction"] +
		transactionActions[actions] +
		"&txhash=" + hash
	return url
}
