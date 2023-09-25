package blockscout

func GetTxInfo(hash string) (*Transaction, error) {
	url := blockScoutApiUrl + moduleValues["transaction"] +
		transactionActions["gettxinfo"] +
		"&txhash=" + hash

	var result Transaction

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetTxReceiptStatus(hash string) (*Status, error) {
	url := blockScoutApiUrl + moduleValues["transaction"] +
		transactionActions["gettxreceiptstatus"] +
		"&txhash=" + hash

	var result Status

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetTxStatus(hash string) (*TxStatus, error) {
	url := blockScoutApiUrl + moduleValues["transaction"] +
		transactionActions["gettxinfo"] +
		"&txhash=" + hash

	var result TxStatus

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
