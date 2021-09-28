package contract_interactor_facade

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	consts "go-evm-client/internal/constants"
	cc "go-evm-client/internal/contracts_template_interface"
	utils "go-evm-client/internal/utils"
	ethacc "go-evm-client/pkg/eth_account"
	ethrpc "go-evm-client/pkg/eth_rpc_client"
	"math/big"
)

// baseContractInteractorFacade holds data common to both the deployer 
// and executor facades
type baseContractInteractorFacade struct {
	userAccount         *ethacc.UserAccount
	ethClient           *ethrpc.EthRpcClient
	currBlockchainState *ethrpc.BlockChainState
	auth                *bind.TransactOpts
	contractType        string
}

// contractDeployerFacade will keep all the necessary data needed to handle 
// contract deployment
type contractDeployerFacade struct {
	baseContractInteractorFacade
	contractArgs []string
}

// NewContractDeployerFacade goes through the processes of creating an
// interactive contract deployer object which is then used to deploy
// contracts
func NewContractDeployerFacade(
	privateKey string,
	rpc string,
	contractArgs []string,
	contractType string,
	gasLimit int,
	gasPrice int,
) (*contractDeployerFacade, error) {
	fmt.Println("Starting account and blockchain connection process.")
	// Process the private key from the flag
	userAccount, err := ethacc.CreateAccount(privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successfully accessed account for Public Key: %s\n",
		userAccount.Account)

	// Connect to the RPC client with the give URL
	ethClient, err1 := ethrpc.CreateClient(rpc)
	if err1 != nil {
		return nil, fmt.Errorf("error: failed to connect to given " +
			"rpc url : %v \n", err1)
	}
	defer ethClient.CloseClient()

	// Attempt to load data from the blockchain given the connected RPC Client
	currBlockchainState, err2 := ethClient.LoadBlockChainState(context.Background())
	if err2 != nil {
		return nil, fmt.Errorf("error: failed to retrieve account using " +
			"provided private key : %v\n", err2)
	}

	fmt.Printf("Succesfully connected to RPC client %s. Current Block Height %d "+
		", for chain id: %d\n", ethClient.RawUrl, currBlockchainState.BlockNumber,
		currBlockchainState.ChainId)

	// Using the client and the account get data needed for contract deployment
	auth, err3 := ethClient.GetDataForTransaction(context.Background(),
		userAccount, currBlockchainState.ChainId, gasLimit, gasPrice)
	if err3 != nil {
		return nil, fmt.Errorf("error: failed to get data for transaction " +
			"processing: %v\n", err3)
	}

	contractDeployerFacade := &contractDeployerFacade{
		baseContractInteractorFacade{
			userAccount,
			ethClient,
			currBlockchainState,
			auth,
			contractType,
		},
		contractArgs,
	}
	fmt.Println("Successfully completed account and blockchain connection " +
		"process.")
	return contractDeployerFacade, nil
}

// DeployContract deploys the contract according to the
// contract types deployment procedure
func (c *contractDeployerFacade) DeployContract() error {
	fmt.Println("Starting contract deployer process.")
	contract := cc.Contract{
		IContract: consts.ContractNamesDict[c.contractType],
	}
	err := contract.DeployContract(
		c.contractArgs,
		c.auth,
		c.ethClient.EthClient)
	if err != nil {
		return err
	}
	fmt.Println("Successfully completed contract deployer process.")
	return nil
}

// contractExecutorFacade will keep all the necessary data needed to handle
// contract execution
type contractExecutorFacade struct {
	baseContractInteractorFacade
	contractAddress common.Address
	funcName        string
	funcArguments   []string
}

// NewContractExecutionFacade goes through the processes of creating an
// interactive contract executor object which is then used to interact with
// contracts
func NewContractExecutionFacade(
	privateKey string,
	rpc string,
	contractType string,
	contractAddress string,
	funcName string,
	funcArguments []string,
	gasLimit int,
	gasPrice int,
) (*contractExecutorFacade, error) {
	fmt.Println("Starting account and blockchain connection process.")
	// Process the private key from the flag
	userAccount, err := ethacc.CreateAccount(privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Succesfully accessed account returned Public Key: %s\n",
		userAccount.Account)

	// Connect to the RPC client with the give URL
	ethClient, err1 := ethrpc.CreateClient(rpc)
	if err1 != nil {
		return nil, fmt.Errorf("error: failed to connect to given " +
			"rpc url : %v \n", err1)
	}
	defer ethClient.CloseClient()

	// Attempt to load data from the blockchain given the connected RPC Client
	currBlockchainState, err2 := ethClient.LoadBlockChainState(context.Background())
	if err2 != nil {
		return nil, fmt.Errorf("error: failed to retrieve account using " +
			"provided private key : %v\n", err2)
	}

	fmt.Printf("Succesfully connected to RPC client %s. Current Block " +
		"Height %d, for chain id: %d\n", ethClient.RawUrl,
		currBlockchainState.BlockNumber, currBlockchainState.ChainId)

	contAddress := common.HexToAddress(contractAddress)
	// Verify the contract exists at the specified address
	ok := ethClient.VerifyContractExistsAtAddress(context.Background(),
		big.NewInt(int64(currBlockchainState.BlockNumber)), contAddress)
	if !ok {
		return nil, fmt.Errorf("error: contract doesn't exist at given address " +
			": %s\n", contractAddress)
	}
	// Using the client and the account get data needed for contract deployment
	auth, err3 := ethClient.GetDataForTransaction(context.Background(),
		userAccount, currBlockchainState.ChainId, gasLimit, gasPrice)
	if err3 != nil {
		return nil, fmt.Errorf("error: failed to get data for transaction " +
			"processing: %v\n", err3)
	}

	contractExecutorFacade := &contractExecutorFacade{
		baseContractInteractorFacade{
			userAccount,
			ethClient,
			currBlockchainState,
			auth,
			contractType,
		},
		contAddress,
		funcName,
		funcArguments,
	}
	fmt.Println("Successfully completed account and blockchain connection " +
		"process.")
	return contractExecutorFacade, nil
}

// LoadContract loads the contract according to the
// contract types loading procedure
func (c *contractExecutorFacade) LoadContract() error {
	fmt.Println("Starting contract loader process.")
	contract := cc.Contract{
		IContract: consts.ContractNamesDict[c.contractType],
	}
	err := contract.LoadContract(
		&c.contractAddress,
		c.ethClient.EthClient)
	if err != nil {
		return err
	}
	fmt.Println("Successfully completed contract loading process.")
	return nil
}

// ExecuteContract executes the given functions on a loaded
// contract. It checks if the function is a query or write operation
// and calls the appropriate functions for each.
func (c *contractExecutorFacade) ExecuteContract() error {
	fmt.Println("Starting contract executor process.")
	queryFuncs := consts.ContractNamesToFuncNames[c.contractType]["query"]
	writeFuncs := consts.ContractNamesToFuncNames[c.contractType]["write"]
	contract := cc.Contract{
		IContract: consts.ContractNamesDict[c.contractType],
	}
	if utils.Contains(&queryFuncs, c.funcName) {
		// Use the query function
		err := contract.QueryContract(c.funcName, c.funcArguments)
		if err != nil {
			return err
		}
	} else if utils.Contains(&writeFuncs, c.funcName) {
		// Use the write function
		err := contract.WriteContract(c.auth, c.funcName, c.funcArguments)
		if err != nil {
			return err
		}
	}
	// There shouldn't be an else{} statement for if the funcName doesn't exist
	// in either slice. Function is to be run assuming all data is provided for.
	fmt.Println("Successfully completed contract execution process.")
	return nil
}
