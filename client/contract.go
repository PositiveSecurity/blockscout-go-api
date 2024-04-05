package client

import "github.com/PositiveSecurity/blockscout-go-api/common"

func (client *BlockScoutAPIClient) GetABI(addr string) string {

	url := client.setContractApiUrl("getabi", addr)

	var response common.Response
	client.sendApiRpcRequest(url, &response)

	return response.Result.(string)
}

func (client *BlockScoutAPIClient) GetContractInfo(addr string) (*[]common.ContractInfo, error) {

	url := client.setContractApiUrl("getsourcecode", addr)

	var result []common.ContractInfo

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	//fmt.Println("Solc", result[0].CompilerVersion)

	return &result, nil
}

func (client *BlockScoutAPIClient) setContractApiUrl(actions string, addr string) string {
	url := client.URL + moduleValues["contract"] +
		accountActions[actions] + "&address=" + addr
	return url
}
