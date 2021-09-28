# GO-EVM-CLIENT

This is a CLI tool built to deploy and interact with smart contracts on any evm-compatible chain.

## Supported contracts

This is a list of supported contracts, that this tool can handle.

1) `FastTestToken`: Basic ERC20 Token with everything pre-determined and no constructor arguments.
2) `DetailedTestToken`: Basic ERC20 Token with 3 constructor arguments and two extra functions to mint and burn tokens.

## Prerequisites

1) `ethermint node`: To deploy your contracts on a local ethermint node please download and install
the ethermint binaries following this [tutorial](https://ethermint.dev/quickstart/installation.html)
2) `Solc`: A Solidity compiler is needed to generate bytecode and abis from smart contracts which
will then be used to generate GoLange code for them. Follow this tutorial to learn more [tutorial](https://goethereumbook.org/en/smart-contract-compile/)
3) `abigen`: This tool's installation can be found in the [tutorial](https://goethereumbook.org/en/smart-contract-compile/) above and is needed to generate Golang based smart contract interfaces.
4) `golang`: Download and install [Golang](https://golang.org/dl/).
5) `ganache-cli` (OPTIONAL): Download and install [Ganache-cli](https://www.trufflesuite.com/ganache)

At the time of writing here are my versions for the above programs

1) `ethermint node`: `0.5.0`
2) `solc/solcjs`: `0.6.12+commit.27d51765.Emscripten.clang`
3) `abigen`: `1.10.8-stable`
4) `go`: `go1.16.8 linux/amd64`
5) `ganache`: `Ganache CLI v6.12.0 (ganache-core: 2.13.0)`

## Scripts

In the `scripts/` folder one can find numerous scripts used to help you with your tasks. Before using the scripts do not forget to set the appropriate permissions using `chmod u+x`. Note these scripts assume you are connecting to a node on `http://127.0.0.1:8545/` if your node has a different `IP:PORT` please adjust accordingly.

1) `build_contracts.sh`: Compiles the contracts using `solcjs` and creates the go packages for the contracts in the appropriate folders
2) `deployed_detailed_contract_ganache.sh`: Using the `cmd/contract_deployer` it deploys the detailed token contract with args to ganache
3) `deployed_fast_contract_ganache.sh`: Using the `cmd/contract_deployer` it deploys the fast token contract without args to ganache
4) `deploy_detailed_contract_local_node.sh`: Using the `cmd/contract_deployer` it deploys the detailed token contract with args to a local ethermint node as well as outputs the private and public keys of your node so that you will be able to use it for contract interactions.
5) `deploy_fast_contract_local_node.sh`: Using the `cmd/contract_deployer` it deploys the fast token contract without args to a local ethermint node as well as outputs the private and public keys of your node so that you will be able to use it for contract interactions.
6) `execute_all_txs_detailed_token.sh`: Using the `cmd/contract_interactor` it loads the detailed token contract from a local ethermint node and executes all the transactions that are possible on that contract. **NOTE** Flags are required for this script.
7) `execute_all_txs_fast_token.sh`: Using the `cmd/contract_interactor` it loads the fast token contract from a local ethermint node and executes all the transactions that are possible on that contract. **NOTE** Flags are required for this script.
8) `init_ethermint_local_node.sh` script used to start a local ethermint node provided you have the binary installed.
9) `run_tests.sh` This will run the GO tests of the program.
10) `start_ganache.sh` This will start the ganache-cli server with a deterministic account.
11) `install_modules.sh` This will download the required go modules and extra so you won't have errors.

## Contract Deployer

The entry code can be found in `cmd/contract_deployer/main.go`. This will deploy a contract based on the arguments you have provided.

###  Usage

To run the deployer enter the command below, additional flags maybe required depending on
the contract type.

`go run cmd/contract_deployer/main.go -p YOUR_PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONSTRUCTOR_ARGUMENTS`

#### DetailedTokenContract Deployment Example

`go run cmd/contract_deployer/main.go -p 266B1CD15B7670B9124B7B67FA92CEEEBEDA56F7B1D5B2E8AA70DD80AB9B7861 -r "http://127.0.0.1:8545" -c detailed_test_token" -a "MintSwapToken" -a "MST" -a "100000000000000000000000000"`

#### FastTokenContract Deployment Example

`go run cmd/contract_deployer/main.go -p 266B1CD15B7670B9124B7B67FA92CEEEBEDA56F7B1D5B2E8AA70DD80AB9B7861 -r "http://127.0.0.1:8545" -c fast_test_token"`

#### Flags

1) `-p`: This is the private key of the account.
2) `-r`: This is the RPC URL of the blockchain you will be connecting to.
3) `-c`: This is the contract type, current supported types are `detailed_test_token` and `fast_test_token`
4) `-a`: These are additional flags for constructor arguments.

## Contract Interactor

The entry code can be found in `cmd/contract_interactor/main.go`. This will load a contract based on the arguments you have provided and execute a transaction or read data from it.

###  Usage

To run the interactor enter the command below, additional flags maybe required depending on
the contract type.

Structure of command: `go run cmd/contract_interactor/main.go -p YOUR_PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f FUNCTION_NAME -fa FUNCTION_ARGUMENTS`

All Commands common to both tokens:

* `Call(): Name`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "name"`
* `Call(): Symbol`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "symbol"`
* `Call(): Decimals`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "decimals"`
* `Call(): TotalSupply`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "totalsupply"`
* `Call(): BalanceOf`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "balanceof" -fa PUB_KEY_1`
* `Call(): Allowance`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "allowance" -fa PUB_KEY_1 -fa PUB_KEY_2`
* `Transact(): Transfer`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "transfer" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`
* `Transact(): Approve`: `go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "approve" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`
* `Transact(): TransferFrom`:`go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "transferfrom" -fa PUB_KEY_1 -fa PUB_KEY_2 -fa TOKEN_AMOUNT`
* `Transact(): IncreaseAllowance`:`go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "increaseallowance" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`
* `Transact(): DecreaseAllowance`:`go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "decreaseallowance" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`

**NOTE** Only DetailedTestToken has these extra functions

* `Transact(): Mint`:`go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "mint" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`
* `Transact(): Burn`:`go run cmd/contract_interactor/main.go -p PRIVATE_KEY -r RPC_URL -c CONTRACT_TYPE -a CONTRACT_ADDRESS -f "burn" -fa PUB_KEY_2 -fa TOKEN_AMOUNT`

## Design

The applications start with a CLI APP process which takes in arguments and verifies that the required args exist. These args are then further verified such as if the Contract type exists of the function under that contract type exists. Using the [Facade Pattern](https://golangbyexample.com/facade-design-pattern-in-golang/) the rpc connection/account login/contract address verification are all handled and a contract interactor interface is returned. This contract Interactor interface can be used to Deploy/Load/Query/Write Smart contracts. This interface is based on the [template pattern](https://golangbyexample.com/template-method-design-pattern-golang/) as nearly all contracts will follow this same flow of execution.

### Deployment Process

* Read CLI Args
* Verify Required Args exist (Private Key, RPC URL, etc)
* Verify Contract Type exists in map of possible types
* Return a Facade Object after processing the necessary data needed for contract deployment
* Use this Facade Object to Deploy the contract, following a template routine
* First the arguments are verified for length and then type converted
* The contract is then deployed
* A message with the transaction hash is shown to the user

### Interactor/Executor Process

* Read CLI Args
* Verify Required Args exist (Private Key, RPC URL, etc)
* Verify Contract Type exists in map of possible types
* Verify Function Type exists under the contract type in map of possible combinations
* Return a Facade Object after processing the necessary data needed for contract execution
* **NOTE** The Facade Obj creation process first checks if there is code deployed at contract location
* Use this Facade Object to Load the contract, following a routine method
* Once the contract is loaded the appropriate functions are executed on it based on Query/Writes
* A message with the completed execution process will be presented to the user.

## Improvements Needed

* `solc/solcjs` Do not use this tool as it doesn't support imports from GitHub, or even locally.
The only solution is to have all the contracts in one file through the use of `truffle-flattener`,
even then the versions of the imported contracts must match that of `solc`. The
alternative to this tool is to have a regular `package.json` which would allow for regular
imports from GitHub and use `waffle` for contract compilation. A shell script would
need to be created to separate the ABI and the EVM Bytecode into their own files for `abigen`
to process them.

* In Solidity many of the contracts are building blocks of each other as seen in Detailed/Fast Tokens, the base contracts from where they inherit their functionality are the same (openzeppelin). What needs to be done in the code for these contracts is to have base contracts such as ERC20/Owner/ERC721 and have these instances inherit from them and extend their functionality. This will reduce the duplicate testing that I have as well as the duplicate contract executors.

* A logger needs to be added together with flags on log levels to control it. As I was running all the transactions the terminal became filled with spam with log data and it became hard to retrieve useful information from it.

* Testing was not completed, the focus was made on the individual Query/Write functions for the tokens as well as the RPC client. More testing needs to be added to achieve maximum coverage and reliability.

* Transaction Hash when Writing to a contract needs to be outputted.
* Documentation should be added to inform other developers on how to add more smart contracts to this tool.

## Noticed Issues with the Go-EVM library

* GasPriceEstimation doesn't work with Ethermint node, returns 0
* Transaction is generated even if contract doesn't exist and no errors are raised. (On Contract Write)