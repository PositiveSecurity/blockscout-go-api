package client

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
