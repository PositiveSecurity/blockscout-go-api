package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (*BlockScoutAPIClient) sendV2ApiRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating request: %v", err)
		return nil, errors.New(errMsg)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error sending request: %v", err)
		return nil, errors.New(errMsg)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error reading response body: %v", err)
		return nil, errors.New(errMsg)
	}

	return body, nil
}
