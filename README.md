# blockscout-api

Golang-client for [blockscout.com](https://www.blockscout.com/) (and its family), with almost complete implementation (accounts, transactions, tokens, contracts, blocks, epoch)

- RPC API Endpoints
- ETH RPC API

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
)

func main() {

	// set your url api
	blockscout.SetBlockScoutApiUrl("https://eth.blockscout.com/api")

	// get the current block number
	block, err := blockscout.GetCurrentBlockRpcApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)

	// convert *big.Int to uint64
	blocknum, err := blockscout.BigIntToUint64(block)
	if err != nil {
		fmt.Println(err)
	}

	// get the ETH balance  of vitalik.eth on the current block number
	addr := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	balance, err := blockscout.GetBalance(addr, blocknum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}
```
