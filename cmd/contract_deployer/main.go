package main

import (
	"errors"
	"fmt"
	cst "go-evm-client/internal/constants"
	cif "go-evm-client/internal/contract_interactor_facade"
	"go-evm-client/internal/utils"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	// CLI Application
	app *cli.App

	// Variables needed to deploy contract
	privateKey, rpc, contractType string
	gasLimit, gasPrice int
	contractArguments cli.StringSlice

	// Flags needed by the contract deployer
	privateKeyFlag = cli.StringFlag{
		Name:        "private, p",
		Usage:       "Private key of the account which will be used to deploy " +
			"the contract.",
		Destination: &privateKey,
	}
	evmRpcUrl = cli.StringFlag{
		Name:        "rpc, r",
		Usage:       "RPC URL of the EVM-compatible blockchain to deploy the " +
			"contract on.",
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
		Usage:       "Name of the contract you want to deploy. Options: " +
			"(detailed_test_token | fast_test_token )",
		Destination: &contractType,
	}
	contractArgs = cli.StringSliceFlag{
		Name:  "args, a",
		Usage: "List of contract arguments that are used in the contract " +
			"constructor.",
		Value: &contractArguments,
	}
)

// Start the CLI application with the required data
func init() {
	app = cli.NewApp()
	app.Name = "deployer"
	app.Usage = "Deploy solidity contracts to any evm compatible blockchain!"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		privateKeyFlag,
		evmRpcUrl,
		gasLimitFlag,
		gasPriceFlag,
		contractFlag,
		contractArgs,
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exitProgramMsg() {
	fmt.Println("Failed deployment exiting program!")
}

func main() {
	// Verify that the required string arguments
	okFlag := utils.RequiredFlagVerification(&[]string{privateKey, rpc, contractType})
	if !okFlag {
		err := errors.New("error: Missing required arguments")
		fmt.Printf("%v \n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	// Verify if the contract type exists
	okType := cst.VerifyContractTypeExists(contractType)
	if !okType {
		err := fmt.Errorf("error: Unsupported contract type %s", contractType)
		fmt.Printf("%v \n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	// Create the contract interactor object to easily interact with the contract
	contractInteractor, err := cif.NewContractDeployerFacade(
		privateKey,
		rpc,
		contractArguments,
		contractType,
		gasLimit,
		gasPrice,
	)
	if err != nil {
		fmt.Printf("%v \n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	// Using the interactor attempt to deploy the contract
	err1 := contractInteractor.DeployContract()
	if err1 != nil {
		fmt.Printf("%v \n", err)
		exitProgramMsg()
		os.Exit(1)
	}
	fmt.Println("Contract deployer finished successfully!")
	os.Exit(0)
}
