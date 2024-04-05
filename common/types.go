package common

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
