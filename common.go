package blockscout

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
)

var moduleValues = map[string]string{
	"account":     "?module=account",
	"block":       "?module=block",
	"contract":    "?module=contract",
	"logs":        "?module=logs",
	"stats":       "?module=stats",
	"token":       "?module=token",
	"transaction": "?module=transaction",
	"epoch":       "?module=epoch", // https://explorer.celo.org/mainnet/api-docs#epoch
}

var accountActions = map[string]string{
	"eth_get_balance": "&action=eth_get_balance",
	"balance":         "&action=balance",
	"balancemulti":    "&action=balancemulti",
	"pendingtxlist":   "&action=pendingtxlist",
	"txlist":          "&action=txlist",
	"txlistinternal":  "&action=txlistinternal",
	"tokentx":         "&action=tokentx",
	"tokenbalance":    "&action=tokenbalance",
	"tokenlist":       "&action=tokenlist",
	"getminedblocks":  "&action=getminedblocks",
	"listaccounts":    "&action=listaccounts",
}

var blockActions = map[string]string{
	"getblockreward":   "&action=getblockreward",
	"getblocknobytime": "&action=getblocknobytime",
	"eth_block_number": "&action=eth_block_number",
}

var contractActions = map[string]string{
	"listcontracts":         "&action=listcontracts",
	"getabi":                "&action=getabi",
	"getsourcecode":         "&action=getsourcecode",
	"verify_via_sourcify":   "&action=verify_via_sourcify",
	"verify_vyper_contract": "&action=verify_vyper_contract",
	"verifysourcecode":      "&action=verifysourcecode",
	"checkverifystatus":     "&action=checkverifystatus",
}

var logsActions = map[string]string{"getLogs": "&action=getLogs"}

var statsActions = map[string]string{
	"tokensupply":       "&action=tokensupply",
	"ethsupplyexchange": "&action=ethsupplyexchange",
	"ethsupply":         "&action=ethsupply",
	"coinsupply":        "&action=coinsupply",
	"coinprice":         "&action=coinprice",
	"totalfees":         "&action=totalfees",
}

var tokenActions = map[string]string{
	"getToken":         "&action=getToken",
	"getTokenHolders":  "&action=getTokenHolders",
	"bridgedTokenList": "&action=bridgedTokenList",
}

var transactionActions = map[string]string{
	"gettxinfo":          "&action=gettxinfo",
	"gettxreceiptstatus": "&action=gettxreceiptstatus",
	"getstatus":          "&action=getstatus",
}

var epochActions = map[string]string{
	"getvoterrewards":     "&action=getvoterrewards",
	"getvalidatorrewards": "&action=getvalidatorrewards",
	"getgrouprewards":     "&action=getgrouprewards",
	"getepoch":            "&action=getepoch",
}

var jsonRpcTemplate = map[string]interface{}{
	"id":      0,
	"jsonrpc": "2.0",
}

type Response struct {
	Message string      `json:"message"`
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
	ID      int         `json:"id"`
}

type Log struct {
	Address string   `json:"address"`
	Data    string   `json:"data"`
	Topics  []string `json:"topics"`
}

type Transaction struct {
	RevertReason      string      `json:"revertReason"`
	BlockHash         string      `json:"blockHash"`
	BlockNumber       string      `json:"blockNumber"`
	Confirmations     string      `json:"confirmations"`
	ContractAddress   string      `json:"contractAddress"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasLimit          string      `json:"gasLimit"`
	Gas               string      `json:"gas"`
	GasPrice          string      `json:"gasPrice"`
	GasUsed           string      `json:"gasUsed"`
	Hash              string      `json:"hash"`
	Input             string      `json:"input"`
	Logs              []Log       `json:"logs"`
	Success           bool        `json:"success"`
	IsError           string      `json:"isError"`
	Nonce             string      `json:"nonce"`
	TimeStamp         interface{} `json:"timeStamp"`
	To                string      `json:"to"`
	TransactionIndex  string      `json:"transactionIndex"`
	TxReceiptStatus   string      `json:"txreceipt_status"`
	Value             string      `json:"value"`
}

type InternalTransaction struct {
	BlockNumber     string `json:"blockNumber"`
	CallType        string `json:"callType"`
	ContractAddress string `json:"contractAddress"`
	ErrorCode       string `json:"errCode"`
	From            string `json:"from"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	Index           string `json:"index"`
	Input           string `json:"input"`
	IsError         string `json:"isError"`
	TimeStamp       string `json:"timeStamp"`
	To              string `json:"to"`
	TransactionHash string `json:"transactionHash"`
	Type            string `json:"type"`
	Value           string `json:"value"`
}

type BlockRewardInfo struct {
	BlockMiner           string `json:"blockMiner"`
	BlockNumber          string `json:"blockNumber"`
	BlockReward          string `json:"blockReward"`
	TimeStamp            int64  `json:"timeStamp"`
	UncleInclusionReward string `json:"uncleInclusionReward"`
	Uncles               string `json:"uncles"`
}

type BalanceMulti struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
	Stale   bool   `json:"stale"`
}

type CoinPrice struct {
	CoinBTCTimestamp string `json:"coin_btc_timestamp"`
	CoinUSD          string `json:"coin_usd"`
	CoinUSDTimestamp string `json:"coin_usd_timestamp"`
}

type TokenInfo struct {
	Balance         string `json:"balance"`
	Cataloged       bool   `json:"cataloged"`
	ContractAddress string `json:"contractAddress"`
	Decimals        string `json:"decimals"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	TotalSupply     string `json:"totalSupply"`
	Type            string `json:"type"`
}

type ContractInfo struct {
	ABI                   string `json:"ABI"`
	CompilerVersion       string `json:"CompilerVersion"`
	ContractName          string `json:"ContractName"`
	FileName              string `json:"FileName"`
	ImplementationAddress string `json:"ImplementationAddress"`
	IsProxy               string `json:"IsProxy"`
	OptimizationUsed      string `json:"OptimizationUsed"`
	SourceCode            string `json:"SourceCode"`
}

type TxStatus struct {
	ErrDescription string `json:"errDescription"`
	IsError        string `json:"isError"`
}

type Status struct {
	Status string `json:"status"`
}

type BlockNumber struct {
	BlockNumber string `json:"blockNumber"`
}

type MinedBlock struct {
	BlockNumber string `json:"blockNumber"`
	BlockReward string `json:"blockReward"`
	TimeStamp   string `json:"timeStamp"`
}

type Epoch struct {
	BlockHash                          string `json:"blockHash"`
	BlockNumber                        string `json:"blockNumber"`
	CarbonOffsettingTargetEpochRewards string `json:"carbonOffsettingTargetEpochRewards"`
	CommunityTargetEpochRewards        string `json:"communityTargetEpochRewards"`
	ElectableValidatorsMax             string `json:"electableValidatorsMax"`
	GoldTotalSupply                    string `json:"goldTotalSupply"`
	ReserveBolster                     string `json:"reserveBolster"`
	ReserveGoldBalance                 string `json:"reserveGoldBalance"`
	RewardsMultiplier                  string `json:"rewardsMultiplier"`
	RewardsMultiplierMax               string `json:"rewardsMultiplierMax"`
	RewardsMultiplierOver              string `json:"rewardsMultiplierOver"`
	RewardsMultiplierUnder             string `json:"rewardsMultiplierUnder"`
	StableUsdTotalSupply               string `json:"stableUsdTotalSupply"`
	TargetTotalSupply                  string `json:"targetTotalSupply"`
	TargetVotingFraction               string `json:"targetVotingFraction"`
	TargetVotingYield                  string `json:"targetVotingYield"`
	TargetVotingYieldAdjustmentFactor  string `json:"targetVotingYieldAdjustmentFactor"`
	TargetVotingYieldMax               string `json:"targetVotingYieldMax"`
	TotalLockedGold                    string `json:"totalLockedGold"`
	TotalNonVoting                     string `json:"totalNonVoting"`
	TotalVotes                         string `json:"totalVotes"`
	ValidatorTargetEpochRewards        string `json:"validatorTargetEpochRewards"`
	VoterTargetEpochRewards            string `json:"voterTargetEpochRewards"`
	VotingFraction                     string `json:"votingFraction"`
}

// URL API in the format "https://instance_base_url/api"
var blockScoutApiUrl string

// Custom errors
var (
	ErrBigInt          = errors.New("failed to convert string to big.Int")
	ErrBigFloat        = errors.New("failed to convert string to big.Float")
	ErrNegativeWei     = errors.New("a negative value of wei is not allowed")
	ErrBigIntTooLarge  = errors.New("*big.Int is too large to fit in a uint64")
	ErrNegativeBigInt  = errors.New("cannot convert negative *big.Int to uint64")
	ErrBalanceNotFound = errors.New("balance not found")
)

func SetBlockScoutApiUrl(url string) {
	blockScoutApiUrl = url
}

func sendApiRpcRequest(url string, response *Response) error {

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

	err = checkError(response)

	if err != nil {
		return err
	}

	return nil
}

func checkError(response *Response) error {
	if response.Message == "OK" || response.Message == "" {
		return nil
	} else if response.Error == "" {
		return nil
	} else {
		return errors.New("error response")
	}
}

func sendApiRpcRequestResult(url string, result any) error {

	fmt.Println(url)

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

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	err = checkError(&response)

	if err != nil {
		return err
	}

	resultBytes, err := marshalToBytes(response.Result)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resultBytes, &result)

	if err != nil {
		return err
	}

	return nil
}

func SaveToFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Convert WEI to ETH
func WeiToEther(wei *big.Int) (*big.Float, error) {

	if wei.Sign() == -1 {
		return nil, ErrNegativeWei
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, new(big.Float).SetInt(divisor))

	return ether, nil
}
