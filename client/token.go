package client

import "github.com/PositiveSecurity/blockscout-go-api/common"

func (client *BlockScoutAPIClient) GetToken(addr string) (*common.TokenInfo, error) {

	url := client.setTokenApiUrl("getToken") + "&contractaddress=" + addr
	var result common.TokenInfo

	err := client.sendApiRpcRequestResult(url, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *BlockScoutAPIClient) setTokenApiUrl(actions string) string {
	url := client.URL + moduleValues["token"] +
		tokenActions[actions]
	return url
}
