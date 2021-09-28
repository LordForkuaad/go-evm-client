package eth_rpc_client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ea "go-evm-client/pkg/eth_account"
	"math/big"
)

// IEthClient is the interface specification to what is needed for
// the RPC Client
type IEthClient interface {
	bind.ContractBackend
	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	Close()
}

// EthRpcClient contains the connected client as well
// as the RawUrl for future use.
type EthRpcClient struct {
	EthClient IEthClient
	RawUrl    string
}

// dialClient makes it easier to test by keeping it outside CreateClient
var dialClient = ethclient.Dial

// CreateClient given the url of the rpc it attempts to establish
// a connection with the node.
func CreateClient(RawUrl string) (*EthRpcClient, error) {
	ethConnection, err := dialClient(RawUrl)
	if err != nil {
		return nil, err
	}
	return &EthRpcClient{ethConnection, RawUrl}, err
}

// CloseClient closes the connection with the client
func (e *EthRpcClient) CloseClient() {
	e.EthClient.Close()
}

// BlockChainState stores the current ChainId and the BlockHeight
// of the connected node. This is needed for transactions and debugging
type BlockChainState struct {
	BlockNumber uint64
	ChainId     *big.Int
}

// LoadBlockChainState using the rpc client query the node for its chain id and
// block height.
func (e *EthRpcClient) LoadBlockChainState(ctx context.Context) (
		*BlockChainState, error) {
	chainId, err := e.EthClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	blockNumber, err1 := e.EthClient.BlockNumber(ctx)
	if err1 != nil {
		return nil, err1
	}
	return &BlockChainState{blockNumber, chainId}, nil
}

// newKeyedTransactionWithChainID makes it easier to test by keeping it outside GetDataForTransaction
var newKeyedTransactionWithChainID = bind.NewKeyedTransactorWithChainID

// GetDataForTransaction gets the authorization data to process transactions
// at a given gas limit and gas price.
func (e *EthRpcClient) GetDataForTransaction(
	ctx context.Context,
	userAccount *ea.UserAccount,
	chainId *big.Int,
	gasLimit int,
	gasPrice int,
) (*bind.TransactOpts, error) {

	nonce, err := e.EthClient.PendingNonceAt(ctx, userAccount.Account)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Retrieved Account nonce %d \n", nonce)

	auth, err1 := newKeyedTransactionWithChainID(userAccount.PrivateKey, chainId)
	if err1 != nil {
		return nil, err1
	}

	// Set Auth Data
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice =  big.NewInt(int64(gasPrice))

	return auth, nil
}

// VerifyContractExistsAtAddress attempt to retrieve the contract at a specific block
func (e *EthRpcClient) VerifyContractExistsAtAddress(
	ctx context.Context,
	blockNumber *big.Int,
	contractAddress common.Address,
) bool {
	code, err := e.EthClient.CodeAt(ctx,contractAddress, blockNumber)
	if err != nil || len(code) == 0{
		return false
	}
	return true
}