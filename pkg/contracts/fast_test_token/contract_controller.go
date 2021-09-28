package fast_test_token

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cc "go-evm-client/internal/contracts_template_interface"
	utils "go-evm-client/internal/utils"
	"go-evm-client/pkg/eth_rpc_client"
	"math/big"
)

// IInstance is the interface needed for these contract functions
type IInstance interface {
	Transfer(
		opts *bind.TransactOpts,
		recipient common.Address,
		amount *big.Int,
	) (*types.Transaction, error)
	Approve(
		opts *bind.TransactOpts,
		spender common.Address,
		amount *big.Int,
	) (*types.Transaction, error)
	TransferFrom(
		opts *bind.TransactOpts,
		sender common.Address,
		recipient common.Address,
		amount *big.Int,
	) (*types.Transaction, error)
	IncreaseAllowance(
		opts *bind.TransactOpts,
		spender common.Address,
		addedValue *big.Int,
	) (*types.Transaction, error)
	DecreaseAllowance(
		opts *bind.TransactOpts,
		spender common.Address,
		subtractedValue *big.Int,
	) (*types.Transaction, error)
	Name(opts *bind.CallOpts) (string, error)
	Symbol(opts *bind.CallOpts) (string, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
	Allowance(
		opts *bind.CallOpts,
		owner common.Address,
		spender common.Address,
	) (*big.Int, error)
}

// FastTestTokenContract contains all the data needed to
// deploy and interact with the FastTestToken contract
type FastTestTokenContract struct {
	cc.Contract
	fastQueriableContractData
	ConstructorArgs contractConstructorArgs
	Address  		common.Address
	LastTx   		*types.Transaction
	Instance 		IInstance
	StrToPrint 		string
}

// contractConstructorArgs is empty and FastTestToken doesn't
// have any arguments on constructor, but it is kept here for
// formality and standardization across contracts
type contractConstructorArgs struct {
}

// fastQueriableContractData is a struct that holds all the data
// that can be queried from the contract
type fastQueriableContractData struct {
	Name		string
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
	BalanceOf   map[common.Address]*big.Int
	Allowance   map[common.Address]map[common.Address]*big.Int
}

// ParseConstructorArguments is kept here so that this contract can
// be used by the templated Contract method
func (f *FastTestTokenContract) ParseConstructorArguments(_ []string) error {
	return nil
}

// DeployContract deploys the fast test token contract and saves
// its instance, tx of deployment and contract address
func (f *FastTestTokenContract) DeployContract(
	auth *bind.TransactOpts,
	client eth_rpc_client.IEthClient,
) error {
	address, tx, instance, err := DeployFastTestToken(auth, client)
	if err != nil {
		return err
	}
	f.Address = address
	f.LastTx = tx
	f.Instance = instance
	return nil
}

// LoadContract loads the fast test token contract and saves
// its instance, contract address
func (f *FastTestTokenContract) LoadContract(
	address *common.Address,
	client eth_rpc_client.IEthClient,
) error {
	instance, err := NewFastTestToken(*address, client)
	if err != nil {
		return err
	}
	f.Instance = instance
	f.Address = *address
	// Insansiate empty maps
	// Instansiate empty maps
	f.BalanceOf = map[common.Address]*big.Int{}
	f.Allowance = map[common.Address]map[common.Address]*big.Int{}
	return nil
}

// PrintDeploymentData outputs to the terminal the address and
// transaction of the deployed contract.
func (f *FastTestTokenContract) PrintDeploymentData() {
	fmt.Printf("Fast Test Token Contract successfully deployed at %s, " +
		"see transaction here %s\n", f.Address.Hex(), f.LastTx.Hash().Hex())
}

// PrintLoadedContractData outs the success message of loading the
// contract as well as the address it's loaded at.
func (f *FastTestTokenContract) PrintLoadedContractData() {
	fmt.Printf("Fast Test Token Contract successfully loaded at %s\n",
		f.Address.Hex())
}

// PrintContractDataAfterExecution print out whatever was saved in StrToPrint
func (f *FastTestTokenContract) PrintContractDataAfterExecution() {
	fmt.Printf("%s", f.StrToPrint)
}

// WriteContract executes write transaction which invokes a state
// change in the FastTestToken contract, this execution is based on
// function name and arguments. Function arguments are verified for
// length and then converted to fit the appropriate types.
func (f *FastTestTokenContract) WriteContract(
	auth *bind.TransactOpts,
	funcName string,
	funcArgs []string,
) error {
	switch funcName {
	case "transfer":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		recipient := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := f.Instance.Transfer(auth, recipient, amount)
		if err1 != nil {
			return err1
		}
		f.LastTx = tx
		f.StrToPrint = fmt.Sprintf("info: Transferred %d tokens at " +
			"FastTestToken (%s) to address %s\n", amount, f.Address, recipient)
	case "approve":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		recipient := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := f.Instance.Approve(auth, recipient, amount)
		if err1 != nil {
			return err1
		}
		f.LastTx = tx
		f.StrToPrint = fmt.Sprintf("info: Approved %d tokens at " +
			"FastTestToken (%s) to address %s\n", amount, f.Address, recipient)
	case "transferfrom":
		err := utils.ValidateLength(&funcArgs, 3)
		if err != nil {
			return err
		}
		sender := common.HexToAddress(funcArgs[0])
		recipient := common.HexToAddress(funcArgs[1])
		amount := new(big.Int)
		amount.SetString(funcArgs[2], 10)
		tx, err1 := f.Instance.TransferFrom(auth, sender, recipient, amount)
		if err1 != nil {
			return err1
		}
		f.LastTx = tx
		f.StrToPrint = fmt.Sprintf("info: Transferred From %s %d tokens at " +
			"FastTestToken (%s) to address %s\n", sender, amount, f.Address, recipient)
	case "increaseallowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		spender := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := f.Instance.IncreaseAllowance(auth, spender, amount)
		if err1 != nil {
			return err1
		}
		f.LastTx = tx
		f.StrToPrint = fmt.Sprintf("info: Increased Allowance by %d tokens at " +
			"FastTestToken (%s) to address %s\n", amount, f.Address, spender)
	case "decreaseallowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		spender := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := f.Instance.DecreaseAllowance(auth, spender, amount)
		if err1 != nil {
			return err1
		}
		f.LastTx = tx
		f.StrToPrint = fmt.Sprintf("info: Decreased Allowance by %d tokens at " +
			"FastTestToken (%s) from address %s\n", amount, f.Address, spender)
	default:
		return nil
	}
	return nil
}

// QueryContract executes query functions which do not invoke a state
// change in the FastTestToken contract, this execution is based on
// function name and arguments. Function arguments are verified for
// length and then converted to fit the appropriate types.
// Data retrieved is then stored in fastQueriableContractData
func (f *FastTestTokenContract) QueryContract(
	funcName string,
	funcArgs []string,
) error {
	switch funcName {
	case "name":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		name, err1 := f.Instance.Name(nil)
		if err1 != nil {
			return err1
		}
		f.Name = name
		f.StrToPrint = fmt.Sprintf("info: Token Name: %s for FastTestToken (%s)" +
			"\n", name, f.Address)
	case "symbol":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		symbol, err1 := f.Instance.Symbol(nil)
		if err1 != nil {
			return err1
		}
		f.Symbol = symbol
		f.StrToPrint = fmt.Sprintf("info: Token Symbol: %s for FastTestToken (%s)" +
			"\n", symbol, f.Address)
	case "decimals":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		decimals, err1 := f.Instance.Decimals(nil)
		if err1 != nil {
			return err1
		}
		f.Decimals = decimals
		f.StrToPrint = fmt.Sprintf("info: Token Decimals: %d for FastTestToken (%s)" +
			"\n", decimals, f.Address)
	case "totalsupply":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		totalSupply, err1 := f.Instance.TotalSupply(nil)
		if err1 != nil {
			return err1
		}
		f.TotalSupply = totalSupply
		f.StrToPrint = fmt.Sprintf("info: Token TotalSupply: %d for FastTestToken (%s)" +
			"\n", totalSupply, f.Address)
	case "balanceof":
		err := utils.ValidateLength(&funcArgs, 1)
		if err != nil {
			return err
		}
		account := common.HexToAddress(funcArgs[0])
		balOfAccount, err1 := f.Instance.BalanceOf(nil, account)
		if err1 != nil {
			return err1
		}
		f.BalanceOf[account] = balOfAccount
		f.StrToPrint = fmt.Sprintf("info: Token Balance of %s : %d for FastTestToken (%s)" +
			"\n", account, balOfAccount, f.Address)
	case "allowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		owner := common.HexToAddress(funcArgs[0])
		spender := common.HexToAddress(funcArgs[1])
		alwOfAccounts, err1 := f.Instance.Allowance(nil, owner, spender)
		if err1 != nil {
			return err1
		}
		// Since this is a nested map of data we must check if `owner`
		// has a map instantiated towards them, if not create one
		if _, ok := f.Allowance[owner]; !ok {
			f.Allowance[owner] = map[common.Address]*big.Int{}
		}
		f.Allowance[owner][spender] = alwOfAccounts
		f.StrToPrint = fmt.Sprintf("info: Token Allowance of spender %s " +
			"from owner %s is %d for FastTestToken (%s)" +
			"\n", spender, owner, alwOfAccounts, f.Address)
	default:
		return nil
	}
	return nil
}
