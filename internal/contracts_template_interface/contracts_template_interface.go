package contracts_template_interface

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go-evm-client/pkg/eth_rpc_client"
)

// iDeployContract interface contains functions that are needed to deploy
// a contract
type iDeployContract interface {

	// DeployContract is used to deploy a contract to the chain
	DeployContract(auth *bind.TransactOpts, client eth_rpc_client.IEthClient) error

	// ParseConstructorArguments helps parse the constructor
	// arguments for each contract as need be
	ParseConstructorArguments(contractArgs []string) error

	// PrintDeploymentData is a generic way to output the result
	// of the contract deployment process
	PrintDeploymentData()
}

// iExecutorContract interface contains functions that are needed
// to load and interact with a contract
type iExecutorContract interface {

	// LoadContract is used to load the contract at a given address
	LoadContract(address *common.Address, client eth_rpc_client.IEthClient) error

	// WriteContract is used to send transactions to the contract to invoke
	// a state change
	WriteContract(
		auth *bind.TransactOpts,
		funcName string,
		funcArgs []string,
	) error

	// QueryContract is used to retrieve data from a contract without
	// invoking a state change
	QueryContract(funcName string, funcArgs []string) error

	// PrintLoadedContractData is a generic way to output the result
	// of the contract interaction
	PrintLoadedContractData()

	// PrintContractDataAfterExecution is a way to print any data
	// we need to know about the contract execution
	PrintContractDataAfterExecution()
}

// IContract interface contains both the deployer and executor
// interfaces
type IContract interface {
	iDeployContract
	iExecutorContract
}

// Contract is a way to hold and use various contracts that
// adhere to the IContract interface
type Contract struct {
	IContract IContract
}

// DeployContract first parses the constructor arguments and
// converts their types for deployment, it then deploys the contract
// to the network and prints it's address and transaction hash.
func (i *Contract) DeployContract(
	contractArgs []string,
	auth *bind.TransactOpts,
	client eth_rpc_client.IEthClient,
) error {
	err := i.IContract.ParseConstructorArguments(contractArgs)
	if err != nil {
		return err
	}
	err1 := i.IContract.DeployContract(auth, client)
	if err1 != nil {
		return err1
	}
	i.IContract.PrintDeploymentData()
	return nil
}

// LoadContract loads the contract at the address and outputs
// it's loaded data
func (i *Contract) LoadContract(
	address *common.Address,
	client eth_rpc_client.IEthClient,
) error {
	err := i.IContract.LoadContract(address, client)
	if err != nil {
		return err
	}
	i.IContract.PrintLoadedContractData()
	return nil
}

// QueryContract accesses the view only functions of a contract
// based on the provided function name and function arguments
func (i *Contract) QueryContract(
	funcName string,
	funcArgs []string,
) error {
	err := i.IContract.QueryContract(funcName, funcArgs)
	if err != nil {
		return err
	}
	i.IContract.PrintLoadedContractData()
	i.IContract.PrintContractDataAfterExecution()
	return nil
}

// WriteContract accesses the write functions of a contract
// based on the provided function name and function arguments
// as well as the auth object used for transaction sending
func (i *Contract) WriteContract(
	auth *bind.TransactOpts,
	funcName string,
	funcArgs []string,
) error {
	err := i.IContract.WriteContract(auth, funcName, funcArgs)
	if err != nil {
		return err
	}
	i.IContract.PrintLoadedContractData()
	i.IContract.PrintContractDataAfterExecution()
	return nil
}
