package blockscout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendEthereumRPCRequest(method string, params interface{}) (*bytes.Buffer, error) {
	jsonRpcTemplate["method"] = method
	jsonRpcTemplate["params"] = params

	requestBody, err := json.Marshal(jsonRpcTemplate)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(blockScoutApiUrl+"/eth-rpc", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody := new(bytes.Buffer)
	responseBody.ReadFrom(resp.Body)

	return responseBody, nil
}

func GetCurrentBlockEthRpc() {
	blockNumberRequest, err := sendEthereumRPCRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Block Number Request:", blockNumberRequest)
	}
}

func GetBalanceEthRpc(addr string) {
	balanceRequest, err := sendEthereumRPCRequest("eth_getBalance", []interface{}{addr, "latest"})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance Request:", balanceRequest)
	}

}
