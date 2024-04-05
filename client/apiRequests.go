package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func (*BlockScoutAPIClient) sendApiRpcRequest(url string, response *common.Response) error {

	//fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) == "Forbidden" {
		fmt.Println("Forbidden")
		return nil
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	err = common.CheckError(response)

	if err != nil {
		return err
	}

	return nil
}

func (*BlockScoutAPIClient) sendApiRpcRequestResult(url string, result any) error {

	// fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) == "Forbidden" {
		fmt.Println("Forbidden")
		return nil
	}

	var response common.Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	err = common.CheckError(&response)

	if err != nil {
		return err
	}

	resultBytes, err := common.MarshalToBytes(response.Result)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resultBytes, &result)

	if err != nil {
		return err
	}

	return nil
}
