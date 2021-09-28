package eth_rpc_client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ea "go-evm-client/pkg/eth_account"
	"math/big"
	"testing"
)

type MockedEthClient struct {
	mock.Mock
}

func (m *MockedEthClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (m *MockedEthClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (m *MockedEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	panic("implement me")
}

func (m *MockedEthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	panic("implement me")
}

func (m *MockedEthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (m *MockedEthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (m *MockedEthClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	panic("implement me")
}

func (m *MockedEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	panic("implement me")
}

func (m *MockedEthClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	panic("implement me")
}

func (m *MockedEthClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	panic("implement me")
}

// Actual functions which we will be using to test our methods
func (m *MockedEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return (args.Get(0)).(*big.Int), args.Error(1)
}

func (m *MockedEthClient)	BlockNumber(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return (args.Get(0)).(uint64), args.Error(1)
}

func (m *MockedEthClient) PendingNonceAt(
	ctx context.Context,
	account common.Address,
) (uint64, error){
	args := m.Called(ctx, account)
	return (args.Get(0)).(uint64), args.Error(1)
}


func (m *MockedEthClient) Close() {
}

func TestCreateClient(t *testing.T) {
	tests := []struct {
		testName			 string
		url 					 string
		expectedError  error
		dialClientFunc func(string) (*ethclient.Client, error)
		expectedUrl 	 string
		expectedClient *ethclient.Client
	}{
		{
			testName: "CreateClient successfull all data returned.",
			url: "http://127.0.0.1:8545/",
			expectedError: nil,
			dialClientFunc: func(_ string) (*ethclient.Client, error) {
				return &ethclient.Client{}, nil
			},
			expectedUrl: "http://127.0.0.1:8545/",
			expectedClient: &ethclient.Client{},
		},
		{
			testName: "CreateClient failed url not valid.",
			url: "127.0.0.1:8545/",
			expectedError: errors.New("rpc url invalid"),
			dialClientFunc: func(_ string) (*ethclient.Client, error) {
				return nil, errors.New("rpc url invalid")
			},
			expectedUrl: "",
			expectedClient: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			dialClient = tt.dialClientFunc
			ethObj, err := CreateClient(tt.url)
			if ethObj != nil {
				assert.NoError(t, err)
				assert.Equal(t, ethObj.RawUrl, tt.expectedUrl)
				assert.Equal(t, ethObj.EthClient, tt.expectedClient)
			}
			assert.Equal(t, err, tt.expectedError)
		})
	}
}

func TestEthRpcClientLoadBlockChainState(t *testing.T) {
	tests := []struct {
		testName			string
		chainId				*big.Int
		blockNumber			uint64
		expectedErrorChain error
		expectedErrorBlock error
		expectedChainId		*big.Int
		expectedBlockNumber uint64
	}{
		{
			testName:		"LoadBlockChainState successfully returned all data.",
			chainId:		big.NewInt(9001),
			blockNumber:	uint64(3000),
			expectedErrorChain:	nil,
			expectedErrorBlock:	nil,
			expectedChainId: big.NewInt(9001),
			expectedBlockNumber: uint64(3000),
		},
		{
			testName:		"LoadBlockChainState ChainID failure.",
			chainId:		big.NewInt(9001),
			blockNumber:	uint64(3000),
			expectedErrorChain:	errors.New("failed ChainID"),
			expectedErrorBlock:	nil,
			expectedChainId: nil,
			expectedBlockNumber: uint64(3000),
		},
		{
			testName:		"LoadBlockChainState BlockNumber failure.",
			chainId:		big.NewInt(9001),
			blockNumber:	uint64(3000),
			expectedErrorChain:	nil,
			expectedErrorBlock:	errors.New("failed BlockNumber"),
			expectedChainId: big.NewInt(9001),
			expectedBlockNumber: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			currContext := context.Background()
			ethClientConn := new(MockedEthClient)
			ethClientConn.On("ChainID", currContext).Return(tt.chainId,
					tt.expectedErrorChain)
			ethClientConn.On("BlockNumber", currContext).Return(tt.blockNumber,
					tt.expectedErrorBlock)

			ethRpcClient := EthRpcClient{ethClientConn, "http://127.0.0.1:8545/"}
			bchState, err := ethRpcClient.LoadBlockChainState(currContext)
			if bchState != nil {
				assert.NoError(t, err)
				assert.Equal(t, bchState.BlockNumber, tt.expectedBlockNumber)
				assert.Equal(t, bchState.ChainId, tt.expectedChainId)
			}
			if err == tt.expectedErrorChain {
				assert.Equal(t, err, tt.expectedErrorChain)
			} else if err == tt.expectedErrorBlock {
				assert.Equal(t, err, tt.expectedErrorBlock)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}

func TestEthRpcClientGetDataForTransaction(t *testing.T) {
	tests := []struct {
		testName	string
		userAccount	*ea.UserAccount
		chainId *big.Int
		gasLimit int
		gasPrice int
		nonce uint64
		newFunc func(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
		expectedErrorNonce error
		expectedErrorAuth error
		expectedNonce *big.Int
		expectedValue *big.Int
		expectedGasLimit uint64
		expectedGasPrice *big.Int
	}{
		{
			testName: "GetDataForTransaction successfully returned all data.",
			userAccount: &ea.UserAccount{},
			chainId: big.NewInt(9001),
			gasLimit: int(800000),
			gasPrice: int(1000),
			nonce: uint64(1),
			newFunc: func(_ *ecdsa.PrivateKey, _ *big.Int) (*bind.TransactOpts, error) {
				return &bind.TransactOpts{}, nil
			},
			expectedErrorNonce:	nil,
			expectedErrorAuth: nil,
			expectedNonce: big.NewInt(int64(1)),
			expectedValue: big.NewInt(0),
			expectedGasLimit: uint64(800000),
			expectedGasPrice:  big.NewInt(int64(1000)),
		},
		{
			testName: "GetDataForTransaction fails to return nonce.",
			userAccount: &ea.UserAccount{},
			chainId: big.NewInt(9001),
			gasLimit: int(800000),
			gasPrice: int(1000),
			nonce: uint64(1),
			newFunc: func(_ *ecdsa.PrivateKey, _ *big.Int) (*bind.TransactOpts, error) {
				return &bind.TransactOpts{}, nil
			},
			expectedErrorNonce:	errors.New("failed Nonce"),
			expectedErrorAuth: nil,
			expectedNonce: big.NewInt(int64(0)),
			expectedValue: big.NewInt(0),
			expectedGasLimit: uint64(800000),
			expectedGasPrice:  big.NewInt(int64(1000)),
		},
		{
			testName: "GetDataForTransaction fails to create auth object.",
			userAccount: &ea.UserAccount{},
			chainId: big.NewInt(9001),
			gasLimit: int(800000),
			gasPrice: int(1000),
			nonce: uint64(1),
			newFunc: func(_ *ecdsa.PrivateKey, _ *big.Int) (*bind.TransactOpts, error) {
				return nil, errors.New("failed creating auth obj")
			},
			expectedErrorNonce:	nil,
			expectedErrorAuth: errors.New("failed creating auth obj"),
			expectedNonce: big.NewInt(int64(0)),
			expectedValue: big.NewInt(0),
			expectedGasLimit: uint64(800000),
			expectedGasPrice:  big.NewInt(int64(1000)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			currContext := context.Background()
			ethClientConn := new(MockedEthClient)
			ethClientConn.On("PendingNonceAt", currContext, tt.userAccount.Account,
				).Return(tt.nonce, tt.expectedErrorNonce)

			newKeyedTransactionWithChainID = tt.newFunc

			ethRpcClient := EthRpcClient{ethClientConn, "http://127.0.0.1:8545/"}
			authState, err := ethRpcClient.GetDataForTransaction(currContext,
				tt.userAccount, tt.chainId, tt.gasLimit, tt.gasPrice)
			if authState != nil {
				assert.NoError(t, err)
				assert.Equal(t, authState.Nonce, tt.expectedNonce)
				assert.Equal(t, authState.Value, tt.expectedValue)
				assert.Equal(t, authState.GasLimit, tt.expectedGasLimit)
				assert.Equal(t, authState.GasPrice, tt.expectedGasPrice)
			}
			if err == tt.expectedErrorNonce {
				assert.Equal(t, err, tt.expectedErrorNonce)
			} else if err != nil {
				// A bit hacky
				assert.Equal(t, err, tt.expectedErrorAuth)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}