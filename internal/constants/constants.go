package constants

import (
	cc "go-evm-client/internal/contracts_template_interface"
	dtt "go-evm-client/pkg/contracts/detailed_test_token"
	ftt "go-evm-client/pkg/contracts/fast_test_token"
)

// ContractNamesDict contains the contract name to the contract struct
// this is used to easily retrieve the required struct without a switch
// statement
var ContractNamesDict = map[string]cc.IContract{
	"detailed_test_token": &dtt.DetailedTestTokenContract{},
	"fast_test_token":     &ftt.FastTestTokenContract{},
}

// baseERC20Queries contains the list of accepted base queries
// of an erc20 token
var baseERC20Queries = []string{
	"name",
	"symbol",
	"decimals",
	"totalsupply",
	"balanceof",
	"allowance",
}

// baseERC20Queries contains the list of accepted base writes
// of an erc20 token
var baseERC20Writes = []string{
	"transfer",
	"approve",
	"transferfrom",
	"increaseallowance",
	"decreaseallowance",
}

// ContractNamesToFuncNames contains the mapping of the possible
// query/write functions that a contract can have
var ContractNamesToFuncNames = map[string]map[string][]string{
	"detailed_test_token": {
		"query": baseERC20Queries,
		"write": append(baseERC20Writes, []string{"mint", "burn"}...),
		"all": append(baseERC20Queries, append(baseERC20Writes, []string{"mint", "burn"}...)...),
	},
	"fast_test_token": {
		"query": baseERC20Queries,
		"write": baseERC20Writes,
		"all": append(baseERC20Queries, baseERC20Writes...),
	},
}

// VerifyContractTypeExists check if the contract type requested exists
func VerifyContractTypeExists(key string) bool {
	_, ok := ContractNamesDict[key]
	return ok
}

// VerifyFunctionNameExists check if the function name exists for the
// requested contract
func VerifyFunctionNameExists(contractType string, funcName string) bool {
	allFunctions := ContractNamesToFuncNames[contractType]["all"]
	for _, v := range allFunctions {
		if v == funcName {
			return true
		}
	}
	return false
}