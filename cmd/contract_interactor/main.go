package main

import (
	"errors"
	"fmt"
	cst "go-evm-client/internal/constants"
	cif "go-evm-client/internal/contract_interactor_facade"
	"go-evm-client/internal/utils"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strings"
)

var (
	// CLI Application
	app *cli.App

	// Variables needed to load contract and interact with contract
	privateKey, rpc, contractType, contractAddress, funcName string
	funcArguments cli.StringSlice
	gasLimit, gasPrice int

	// Flags needed by the contract deployer
	privateKeyFlag = cli.StringFlag{
		Name:        "private, p",
		Usage:       "Private key of the account which will be used to interact with the contract.",
		Destination: &privateKey,
	}
	evmRpcUrl = cli.StringFlag{
		Name:        "rpc, r",
		Usage:       "RPC URL of the EVM-compatible blockchain where the contract is deployed.",
		Destination: &rpc,
	}
	gasLimitFlag = cli.IntFlag{
		Name:        "gaslimit, gl",
		Usage:       "Gas limit is the maximum amount of gas you are willing to " +
			"pay for the deployment of the contract.",
		Value: 8000000,
		Destination: &gasLimit,
	}
	gasPriceFlag = cli.IntFlag{
		Name:        "gasprice, gp",
		Usage:       "Gas Price is the amount you want to pay for the deployment " +
			"of the contract.",
		Value: 1000,
		Destination: &gasPrice,
	}
	contractFlag = cli.StringFlag{
		Name:        "contract, c",
		Usage:       "Name of the contract you want to deploy. Options: (detailed_test_token | fast_test_token ).",
		Destination: &contractType,
	}
	addressFlag = cli.StringFlag{
		Name:        "address, a",
		Usage:       "Address of the contract.",
		Destination: &contractAddress,
	}
	funNameFlag = cli.StringFlag{
		Name:        "function, f",
		Usage:       "The name of the function which you wish to execute in the contract.",
		Destination: &funcName,
	}
	funcArgs = cli.StringSliceFlag{
		Name:  "args, fa",
		Usage: "List of function arguments that are used in the contract function.",
		Value: &funcArguments,
	}
)

// Start the CLI application with the required data
func init() {
	app = cli.NewApp()
	app.Name = "interactor"
	app.Usage = "Interact with solidity contracts with any chain!"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		privateKeyFlag,
		evmRpcUrl,
		gasLimitFlag,
		gasPriceFlag,
		contractFlag,
		addressFlag,
		funNameFlag,
		funcArgs,
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exitProgramMsg() {
	fmt.Println("Failed contract interaction exiting program!")
}

func main() {
	// Convert the function name to lowercase for ease of user use
	funcName = strings.ToLower(funcName)
	// Verify that the required string arguments
	okFlag := utils.RequiredFlagVerification(&[]string{
		privateKey, rpc, contractType, contractAddress, funcName})
	if !okFlag {
		err := errors.New("error: Missing required arguments")
		fmt.Printf("%v\n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	// Verify if the contract type exists
	okType := cst.VerifyContractTypeExists(contractType)
	if !okType {
		err := fmt.Errorf("error: Unsupported contract type %s", contractType)
		fmt.Printf("%v\n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	// Verify if the function name exists
	okFuncName := cst.VerifyFunctionNameExists(contractType, funcName)
	if !okFuncName {
		err := fmt.Errorf("error: Unsupported function name %s for contract " +
			"type %s", funcName, contractType)
		fmt.Printf("%v\n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	contractExecutor, err := cif.NewContractExecutionFacade(
		privateKey,
		rpc,
		contractType,
		contractAddress,
		funcName,
		funcArguments,
		gasLimit,
		gasPrice,
	)
	if err != nil {
		fmt.Printf("%v\n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	err1 := contractExecutor.LoadContract()
	if err1 != nil {
		fmt.Printf("%v\n", err1)
		exitProgramMsg()
		os.Exit(1)
	}
	err2 := contractExecutor.ExecuteContract()
	if err2 != nil {
		fmt.Printf("%v\n", err2)
		exitProgramMsg()
		os.Exit(1)
	}
}
