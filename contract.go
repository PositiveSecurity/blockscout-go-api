package blockscout

func GetABI(addr string) string {
	url := blockScoutApiUrl + moduleValues["contract"] +
		contractActions["getabi"] +
		"&address=" + addr

	var response Response
	sendApiRpcRequest(url, &response)

	return response.Result.(string)
}

func GetContractInfo(addr string) (*[]ContractInfo, error) {
	url := blockScoutApiUrl + moduleValues["contract"] +
		contractActions["getsourcecode"] +
		"&address=" + addr

	var result []ContractInfo

	err := sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	//fmt.Println("Solc", result[0].CompilerVersion)

	return &result, nil
}
