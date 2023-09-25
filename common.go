package blockscout

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
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

var blockScoutApiUrl string

func SetBlockScoutApiUrl(url string) {
	blockScoutApiUrl = url
}

func sendApiRpcRequest(url string, response *Response) error {

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

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	err = checkMessage(response.Message)

	if err != nil {
		return err
	}

	return nil
}

func checkMessage(mes string) error {
	if mes != "OK" {
		return errors.New(mes)
	}
	return nil
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

	err = checkMessage(response.Message)

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

func weiToEther(wei *big.Int) (*big.Float, error) {

	if wei.Sign() == -1 {
		return nil, errors.New("A negative value of wei is not allowed")
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, new(big.Float).SetInt(divisor))

	return ether, nil
}

func hexStringToBigInt(str string) (*big.Int, error) {
	str = str[2:]
	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(str, 16)

	if !ok {
		return nil, errors.New("Failed to convert the string to a big.Int")
	}

	return bigInt, nil
}

func decStringToBigInt(str string) (*big.Int, error) {
	var bigintVal big.Int

	_, success := bigintVal.SetString(str, 10)

	if !success {
		return nil, errors.New("Failed to convert the string to a big.Int")
	}

	return &bigintVal, nil
}

func decStringToBigFloat(str string) (*big.Float, error) {
	var floatValue big.Float
	_, success := floatValue.SetString(str)

	if !success {
		return nil, errors.New("Failed to convert string to big.Float")
	}

	return &floatValue, nil
}

func decStringToUint(str string) (uint, error) {
	intValue, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(intValue), nil
}

func marshalToBytes(result any) ([]byte, error) {
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func SaveToFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func UnixTimestampToNormal(unixTimestamp int64) string {
	t := time.Unix(unixTimestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
