# blockscout-api

[![Go Reference](https://pkg.go.dev/badge/github.com/PositiveSecurity/blockscout-go-api.svg)](https://pkg.go.dev/github.com/PositiveSecurity/blockscout-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/PositiveSecurity/blockscout-go-api)](https://goreportcard.com/report/github.com/PositiveSecurity/blockscout-go-api)

Golang-client for [blockscout.com](https://www.blockscout.com/) (and its family), with almost complete implementation (accounts, transactions, tokens, contracts, blocks, epoch)

- RPC API Endpoints
- ETH RPC API
- REST API

While there are several explores available to blockchain projects, most are closed systems (ie Etherscan, Etherchain). Blockscout provides a much needed open-source alternative.

# Usage

```
go get github.com/PositiveSecurity/blockscout-go-api
```

Create an API instance and off you go.

```go
package main

import (
	"fmt"

	"github.com/PositiveSecurity/blockscout-go-api"
	"github.com/PositiveSecurity/blockscout-go-api/client"
	"github.com/PositiveSecurity/blockscout-go-api/common"
)

func main() {

	var client client.BlockScoutAPIClient

	// set your url api
	client.SetBlockScoutApiUrl(blockscout.EthMainnet)
	client.SetBlockScoutApiUrlV2(blockscout.EthMainnetV2)

	// get the current block number
	block, err := client.GetCurrentBlockRpcApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)

	// convert *big.Int to uint64
	blocknum, err := common.BigIntToUint64(block)
	if err != nil {
		fmt.Println(err)
	}

	// get the ETH balance  of vitalik.eth on the current block number
	addr := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	balance, err := client.GetBalance(addr, blocknum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)

	// api v2
	body, err := client.GetBlockInfo("0")
	if err != nil {
		fmt.Println(err)
	}

	// print json
	common.PrettyPrintJSON(body)

}

```
