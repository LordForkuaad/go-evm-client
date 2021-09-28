package detailed_test_token

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cc "go-evm-client/internal/contracts_template_interface"
	utils "go-evm-client/internal/utils"
	"go-evm-client/pkg/eth_rpc_client"
	"math/big"
	"reflect"
)

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
	Mint(opts *bind.TransactOpts,
		to common.Address,
		amount *big.Int,
	) (*types.Transaction, error)
	Burn(opts *bind.TransactOpts,
		from common.Address,
		amount *big.Int,
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

// DetailedTestTokenContract contains all the data needed to
// deploy and interact with the DetailedTestToken contract
type DetailedTestTokenContract struct {
	cc.Contract
	ConstructorArgs contractConstructorArgs
	detailedQueriableContractData
	Address  common.Address
	LastTx   *types.Transaction
	Instance IInstance
	StrToPrint 		string
}

// contractConstructorArgs are the details needed to deploy
// the DetailedTestToken. These details are passed onto the
// contract's constructor
type contractConstructorArgs struct {
	name   string
	symbol string
	amount *big.Int
}

// detailedQueriableContractData is a struct that holds all the data
// that can be queried from the contract
type detailedQueriableContractData struct {
	Name        string
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
	BalanceOf   map[common.Address]*big.Int
	Allowance   map[common.Address]map[common.Address]*big.Int
}

// ParseConstructorArguments is used to parse the slice of strings
// into the specific constructor arguments needed for contract deployment
// these are then stored in contractConstructorArgs
func (d *DetailedTestTokenContract) ParseConstructorArguments(
	contractArgs []string) error {
	neededArgs := reflect.TypeOf(contractConstructorArgs{}).NumField()
	recArgs := len(contractArgs)
	if recArgs != neededArgs {
		return fmt.Errorf("error: incorrect amount of arguments, args " +
			"needed : %d != args received %d", neededArgs, recArgs)
	}
	amount := new(big.Int)
	amount.SetString(contractArgs[2], 10)
	ccArgs := contractConstructorArgs{
		name:   contractArgs[0],
		symbol: contractArgs[1],
		amount: amount,
	}
	d.ConstructorArgs = ccArgs
	return nil
}

// DeployContract deploys the detailed test token contract and saves
// its instance, tx of deployment and contract address
func (d *DetailedTestTokenContract) DeployContract(
	auth *bind.TransactOpts,
	client eth_rpc_client.IEthClient,
) error {
	address, tx, instance, err := DeployDetailedTestToken(
		auth,
		client,
		d.ConstructorArgs.name,
		d.ConstructorArgs.symbol,
		d.ConstructorArgs.amount)
	if err != nil {
		return err
	}
	d.Address = address
	d.LastTx = tx
	d.Instance = instance
	return nil
}

// LoadContract loads the detailed test token contract and saves
// its instance, contract address
func (d *DetailedTestTokenContract) LoadContract(
	address *common.Address,
	client eth_rpc_client.IEthClient,
) error {
	instance, err := NewDetailedTestToken(*address, client)
	if err != nil {
		return err
	}
	d.Instance = instance
	d.Address = *address
	// Instantiate empty maps
	d.BalanceOf = map[common.Address]*big.Int{}
	d.Allowance = map[common.Address]map[common.Address]*big.Int{}
	return nil
}

// PrintDeploymentData outputs to the terminal the address and
// transaction of the deployed contract.
func (d *DetailedTestTokenContract) PrintDeploymentData() {
	fmt.Printf("Detailed Test Token Contract successfully deployed at %s, " +
		"see transaction here %s \n", d.Address.Hex(), d.LastTx.Hash().Hex())
}

// PrintLoadedContractData outs the success message of loading the
// contract as well as the address it's loaded at.
func (d *DetailedTestTokenContract) PrintLoadedContractData() {
	fmt.Printf("Detailed Test Token Contract successfully loaded at %s \n",
		d.Address.Hex())
}

// PrintContractDataAfterExecution print out whatever was saved in StrToPrint
func (d *DetailedTestTokenContract) PrintContractDataAfterExecution() {
	fmt.Printf("%s", d.StrToPrint)
}

// WriteContract executes write transaction which invokes a state
// change in the DetailedTestToken contract, this execution is based on
// function name and arguments. Function arguments are verified for
// length and then converted to fit the appropriate types.
func (d *DetailedTestTokenContract) WriteContract(
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
		tx, err1 := d.Instance.Transfer(auth, recipient, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Transferred %d tokens at " +
			"DetailedTestToken (%s) to address %s\n", amount, d.Address, recipient)
	case "approve":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		recipient := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := d.Instance.Approve(auth, recipient, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Approved %d tokens at " +
			"DetailedTestToken (%s) to address %s\n", amount, d.Address, recipient)
	case "transferfrom":
		err := utils.ValidateLength(&funcArgs, 3)
		if err != nil {
			return err
		}
		sender := common.HexToAddress(funcArgs[0])
		recipient := common.HexToAddress(funcArgs[1])
		amount := new(big.Int)
		amount.SetString(funcArgs[2], 10)
		tx, err1 := d.Instance.TransferFrom(auth, sender, recipient, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Transferred From %s %d tokens at " +
			"DetailedTestToken (%s) to address %s\n", sender, amount, d.Address,
			recipient)
	case "increaseallowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		spender := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := d.Instance.IncreaseAllowance(auth, spender, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Increased Allowance by %d tokens at " +
			"DetailedTestToken (%s) to address %s\n", amount, d.Address, spender)
	case "decreaseallowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		spender := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := d.Instance.DecreaseAllowance(auth, spender, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Decreased Allowance by %d tokens at " +
			"DetailedTestToken (%s) from address %s\n", amount, d.Address, spender)
	case "mint":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		to := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := d.Instance.Mint(auth, to, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Minted %d tokens to %s at " +
			"DetailedTestToken (%s)\n", amount, to, d.Address)
	case "burn":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		from := common.HexToAddress(funcArgs[0])
		amount := new(big.Int)
		amount.SetString(funcArgs[1], 10)
		tx, err1 := d.Instance.Burn(auth, from, amount)
		if err1 != nil {
			return err1
		}
		d.LastTx = tx
		d.StrToPrint = fmt.Sprintf("info: Burned %d tokens from %s at " +
			"DetailedTestToken (%s)\n", amount, from, d.Address)
	default:
		return nil
	}
	return nil
}

// QueryContract executes query functions which do not invoke a state
// change in the DetailedTestToken contract, this execution is based on
// function name and arguments. Function arguments are verified for
// length and then converted to fit the appropriate types.
// Data retrieved is then stored in detailedQueriableContractData
func (d *DetailedTestTokenContract) QueryContract(
	funcName string,
	funcArgs []string,
) error {
	switch funcName {
	case "name":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		name, err1 := d.Instance.Name(nil)
		if err1 != nil {
			return err1
		}
		d.Name = name
		d.StrToPrint = fmt.Sprintf("info: Token Name: %s for DetailedTestToken (%s)" +
			"\n", name, d.Address)
	case "symbol":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		symbol, err1 := d.Instance.Symbol(nil)
		if err1 != nil {
			return err1
		}
		d.Symbol = symbol
		d.StrToPrint = fmt.Sprintf("info: Token Symbol: %s for DetailedTestToken (%s)" +
			"\n", symbol, d.Address)
	case "decimals":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		decimals, err1 := d.Instance.Decimals(nil)
		if err1 != nil {
			return err1
		}
		d.Decimals = decimals
		d.StrToPrint = fmt.Sprintf("info: Token Decimals: %d for DetailedTestToken (%s)" +
			"\n", decimals, d.Address)
	case "totalsupply":
		err := utils.ValidateLength(&funcArgs, 0)
		if err != nil {
			return err
		}
		totalSupply, err1 := d.Instance.TotalSupply(nil)
		if err1 != nil {
			return err1
		}
		d.TotalSupply = totalSupply
		d.StrToPrint = fmt.Sprintf("info: Token TotalSupply: %d for DetailedTestToken (%s)" +
			"\n", totalSupply, d.Address)
	case "balanceof":
		err := utils.ValidateLength(&funcArgs, 1)
		if err != nil {
			return err
		}
		account := common.HexToAddress(funcArgs[0])
		balOfAccount, err1 := d.Instance.BalanceOf(nil, account)
		if err1 != nil {
			return err1
		}
		d.BalanceOf[account] = balOfAccount
		d.StrToPrint = fmt.Sprintf("info: Token Balance of %s : %d for DetailedTestToken (%s)" +
			"\n", account, balOfAccount, d.Address)
	case "allowance":
		err := utils.ValidateLength(&funcArgs, 2)
		if err != nil {
			return err
		}
		owner := common.HexToAddress(funcArgs[0])
		spender := common.HexToAddress(funcArgs[1])
		alwOfAccounts, err1 := d.Instance.Allowance(nil, owner, spender)
		if err1 != nil {
			return err1
		}
		// Since this is a nested map of data we must check if `owner`
		// has a map instantiated towards them, if not create one
		if _, ok := d.Allowance[owner]; !ok {
			d.Allowance[owner] = map[common.Address]*big.Int{}
		}
		d.Allowance[owner][spender] = alwOfAccounts
		d.StrToPrint = fmt.Sprintf("info: Token Allowance of spender %s " +
			"from owner %s is %d for DetailedTestToken (%s)" +
			"\n", spender, owner, alwOfAccounts, d.Address)
	default:
		return nil
	}
	return nil
}
