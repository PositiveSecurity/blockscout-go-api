package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var jsonRpcTemplate = map[string]interface{}{
	"id":      0,
	"jsonrpc": "2.0",
}

func (client *BlockScoutAPIClient) GetCurrentBlockEthRpc() {
	blockNumberRequest, err := client.sendEthereumRPCRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Block Number Request:", blockNumberRequest)
	}
}

func (client *BlockScoutAPIClient) GetBalanceEthRpc(addr string) {
	balanceRequest, err := client.sendEthereumRPCRequest("eth_getBalance", []interface{}{addr, "latest"})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance Request:", balanceRequest)
	}

}

func (client *BlockScoutAPIClient) sendEthereumRPCRequest(method string, params interface{}) (*bytes.Buffer, error) {
	jsonRpcTemplate["method"] = method
	jsonRpcTemplate["params"] = params

	requestBody, err := json.Marshal(jsonRpcTemplate)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(client.URL+"/eth-rpc", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody := new(bytes.Buffer)
	responseBody.ReadFrom(resp.Body)

	return responseBody, nil
}
