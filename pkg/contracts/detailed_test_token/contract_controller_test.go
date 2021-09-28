package detailed_test_token

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
)

type MockContractInstance struct {
	mock.Mock
}

func (m *MockContractInstance) Transfer(
	opts *bind.TransactOpts,
	recipient common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	args := m.Called(opts, recipient, amount)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) Approve(
	opts *bind.TransactOpts,
	spender common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	args := m.Called(opts, spender, amount)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) 	TransferFrom(
	opts *bind.TransactOpts,
	sender common.Address,
	recipient common.Address,
	amount *big.Int,
) (*types.Transaction, error){
	args := m.Called(opts, sender, recipient, amount)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) IncreaseAllowance(
	opts *bind.TransactOpts,
	spender common.Address,
	addedValue *big.Int,
) (*types.Transaction, error) {
	args := m.Called(opts, spender, addedValue)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) DecreaseAllowance(
	opts *bind.TransactOpts,
	spender common.Address,
	subtractedValue *big.Int,
) (*types.Transaction, error) {
	args := m.Called(opts, spender, subtractedValue)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) Mint(
	opts *bind.TransactOpts,
	to common.Address,
	amount *big.Int,
) (*types.Transaction, error){
	args := m.Called(opts, to, amount)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}

func (m *MockContractInstance) Burn(opts *bind.TransactOpts,
	from common.Address,
	amount *big.Int,
) (*types.Transaction, error){
	args := m.Called(opts, from, amount)
	return (args.Get(0)).(*types.Transaction), args.Error(1)
}


func (m *MockContractInstance) Name(_ *bind.CallOpts) (string, error) {
	args := m.Called(nil)
	return (args.Get(0)).(string), args.Error(1)
}

func (m *MockContractInstance) Symbol(_ *bind.CallOpts) (string, error) {
	args := m.Called(nil)
	return (args.Get(0)).(string), args.Error(1)
}

func (m *MockContractInstance) Decimals(_ *bind.CallOpts) (uint8, error){
	args := m.Called(nil)
	return (args.Get(0)).(uint8), args.Error(1)
}

func (m *MockContractInstance) TotalSupply(_ *bind.CallOpts) (*big.Int, error) {
	args := m.Called(nil)
	return (args.Get(0)).(*big.Int), args.Error(1)
}

func (m *MockContractInstance) BalanceOf(_ *bind.CallOpts, account common.Address) (*big.Int, error) {
	args := m.Called(nil, account)
	return (args.Get(0)).(*big.Int), args.Error(1)
}

func (m *MockContractInstance) 	Allowance(
	_ *bind.CallOpts,
	owner common.Address,
	spender common.Address,
) (*big.Int, error) {
	args := m.Called(nil, owner, spender)
	return (args.Get(0)).(*big.Int), args.Error(1)
}

func TestQueryContractNameMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedStrReturn	string
	}{
		{
			testName: "QueryContract func Name successful all data returned.",
			funcName: "name",
			funcArgs: []string{},
			strToPrint: "info: Token Name: DetailedTestToken for DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedStrReturn: "DetailedTestToken",
		},
		{
			testName: "QueryContract func Name fail arg len validation.",
			funcName: "name",
			funcArgs: []string{"fail"},
			strToPrint: "",
			expectedError: errors.New("error: 1 arguments does not match required 0"),
			expectedStrReturn: "",
		},
		{
			testName: "QueryContract func Name instance failure.",
			funcName: "name",
			funcArgs: []string{},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedStrReturn: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Name", nil).Return(tt.expectedStrReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.Name, tt.expectedStrReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestQueryContractSymbolMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedStrReturn	string
	}{
		{
			testName: "QueryContract func Symbol successful all data returned.",
			funcName: "symbol",
			funcArgs: []string{},
			strToPrint: "info: Token Symbol: DTT for DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedStrReturn: "DTT",
		},
		{
			testName: "QueryContract func Symbol fail arg len validation.",
			funcName: "symbol",
			funcArgs: []string{"fail"},
			strToPrint: "",
			expectedError: errors.New("error: 1 arguments does not match required 0"),
			expectedStrReturn: "",
		},
		{
			testName: "QueryContract func Symbol instance failure.",
			funcName: "symbol",
			funcArgs: []string{},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedStrReturn: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Symbol", nil).Return(tt.expectedStrReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.Symbol, tt.expectedStrReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestQueryContractDecimalsMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedIntReturn uint8
	}{
		{
			testName: "QueryContract func Decimals successful all data returned.",
			funcName: "decimals",
			funcArgs: []string{},
			strToPrint: "info: Token Decimals: 18 for DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedIntReturn: 18,
		},
		{
			testName: "QueryContract func Decimals fail arg len validation.",
			funcName: "decimals",
			funcArgs: []string{"fail"},
			strToPrint: "",
			expectedError: errors.New("error: 1 arguments does not match required 0"),
			expectedIntReturn: 0,
		},
		{
			testName: "QueryContract func Decimals instance failure.",
			funcName: "decimals",
			funcArgs: []string{},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedIntReturn: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Decimals", nil).Return(tt.expectedIntReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.Decimals, tt.expectedIntReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestQueryContractTotalSupplyMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedBigIntReturn *big.Int
	}{
		{
			testName: "QueryContract func TotalSupply successful all data returned.",
			funcName: "totalsupply",
			funcArgs: []string{},
			strToPrint: "info: Token TotalSupply: 100000000000 for DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedBigIntReturn: big.NewInt(100000000000),
		},
		{
			testName: "QueryContract func TotalSupply fail arg len validation.",
			funcName: "totalsupply",
			funcArgs: []string{"fail"},
			strToPrint: "",
			expectedError: errors.New("error: 1 arguments does not match required 0"),
			expectedBigIntReturn: nil,
		},
		{
			testName: "QueryContract func TotalSupply instance failure.",
			funcName: "totalsupply",
			funcArgs: []string{},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedBigIntReturn: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("TotalSupply", nil).Return(tt.expectedBigIntReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.TotalSupply, tt.expectedBigIntReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestQueryContractBalanceOfMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedBigIntReturn *big.Int
	}{
		{
			testName: "QueryContract func BalanceOf successful all data returned.",
			funcName: "balanceof",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df"},
			strToPrint: "info: Token Balance of 0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df : " +
				"100000000000 for DetailedTestToken (0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedBigIntReturn: big.NewInt(100000000000),
		},
		{
			testName: "QueryContract func BalanceOf fail arg len validation.",
			funcName: "balanceof",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 2 arguments does not match required 1"),
			expectedBigIntReturn: nil,
		},
		{
			testName: "QueryContract func BalanceOf instance failure.",
			funcName: "balanceof",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedBigIntReturn: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			account := common.HexToAddress(tt.funcArgs[0])
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("BalanceOf", nil, account).Return(tt.expectedBigIntReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			// Need to set the empty map
			dttc.BalanceOf = map[common.Address]*big.Int{}
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.BalanceOf[account], tt.expectedBigIntReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestQueryContractAllowanceMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedBigIntReturn *big.Int
	}{
		{
			testName: "QueryContract func Allowance successful all data returned.",
			funcName: "allowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0xfd6f5A60D2D8b12039F906D112f10Fb66F881087"},
			strToPrint: "info: Token Allowance of spender 0xfd6f5A60D2D8b12039F906D112f10Fb66F881087 " +
				"from owner 0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df is 100000000000 for DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedBigIntReturn: big.NewInt(100000000000),
		},
		{
			testName: "QueryContract func Allowance fail arg len validation.",
			funcName: "allowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0xfd6f5A60D2D8b12039F906D112f10Fb66F881087", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedBigIntReturn: nil,
		},
		{
			testName: "QueryContract func Allowance instance failure.",
			funcName: "allowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0xfd6f5A60D2D8b12039F906D112f10Fb66F881087"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedBigIntReturn: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			owner := common.HexToAddress(tt.funcArgs[0])
			spender := common.HexToAddress(tt.funcArgs[1])
			mInstance := new(MockContractInstance)
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Allowance", nil, owner, spender).Return(tt.expectedBigIntReturn,
				tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			dttc.Allowance = map[common.Address]map[common.Address]*big.Int{}
			err := dttc.QueryContract(tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.Allowance[owner][spender], tt.expectedBigIntReturn)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}


func TestWriteContractTransferMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func Transfer successful all data returned.",
			funcName: "transfer",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "info: Transferred 100000000000 tokens at DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000) to address " +
				"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func Transfer fail arg len validation.",
			funcName: "transfer",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func Transfer instance failure.",
			funcName: "transfer",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			recipient := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Transfer", auth, recipient, amount).Return(
				tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestWriteContractApproveMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func Approve successful all data returned.",
			funcName: "approve",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "info: Approved 100000000000 tokens at DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000) to address " +
				"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func Approve fail arg len validation.",
			funcName: "approve",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func Approve instance failure.",
			funcName: "approve",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			recipient := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Approve", auth, recipient, amount).Return(
				tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestWriteContractTransferFromMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func TransferFrom successful all data returned.",
			funcName: "transferfrom",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0x59Ba9FfE3bE7E39479B39eAD755AF9994E974384", "100000000000"},
			strToPrint: "info: Transferred From 0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df " +
				"100000000000 tokens at DetailedTestToken " +
				"(0x0000000000000000000000000000000000000000) to address " +
				"0x59Ba9FfE3bE7E39479B39eAD755AF9994E974384\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func TransferFrom fail arg len validation.",
			funcName: "transferfrom",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0x59Ba9FfE3bE7E39479B39eAD755AF9994E974384", "100000000000",
				"fail"},
			strToPrint: "",
			expectedError: errors.New("error: 4 arguments does not match required 3"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func TransferFrom instance failure.",
			funcName: "transferfrom",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"0x59Ba9FfE3bE7E39479B39eAD755AF9994E974384", "100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			sender := common.HexToAddress(tt.funcArgs[0])
			recipient := common.HexToAddress(tt.funcArgs[1])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[2], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("TransferFrom", auth, sender, recipient,
				amount).Return(tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestWriteContractIncreaseAllowanceFromMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func IncreaseAllowance successful all data returned.",
			funcName: "increaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df", "100000000000"},
			strToPrint: "info: Increased Allowance by 100000000000 tokens at " +
				"DetailedTestToken (0x0000000000000000000000000000000000000000) " +
				"to address 0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func IncreaseAllowance fail arg len validation.",
			funcName: "increaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func IncreaseAllowance instance failure.",
			funcName: "increaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			spender := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("IncreaseAllowance", auth, spender,
				amount).Return(tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}


func TestWriteContractDecreaseAllowanceFromMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func DecreaseAllowance successful all data returned.",
			funcName: "decreaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df", "100000000000"},
			strToPrint: "info: Decreased Allowance by 100000000000 tokens at " +
				"DetailedTestToken (0x0000000000000000000000000000000000000000) " +
				"from address 0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func DecreaseAllowance fail arg len validation.",
			funcName: "decreaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func DecreaseAllowance instance failure.",
			funcName: "decreaseallowance",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			spender := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("DecreaseAllowance", auth, spender,
				amount).Return(tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestWriteContractMintMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func Mint successful all data returned.",
			funcName: "mint",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "info: Minted 100000000000 tokens to " +
				"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df at " +
				"DetailedTestToken (0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func Mint fail arg len validation.",
			funcName: "mint",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func Mint instance failure.",
			funcName: "mint",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			to := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Mint", auth, to, amount).Return(
				tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}

func TestWriteContractBurnMethod(t *testing.T) {
	tests := []struct {
		testName string
		funcName string
		funcArgs []string
		strToPrint string
		expectedError error
		expectedTx *types.Transaction
	}{
		{
			testName: "WriteContract func Burn successful all data returned.",
			funcName: "burn",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "info: Burned 100000000000 tokens from " +
				"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df at " +
				"DetailedTestToken (0x0000000000000000000000000000000000000000)\n",
			expectedError: nil,
			expectedTx: &types.Transaction{},
		},
		{
			testName: "WriteContract func Burn fail arg len validation.",
			funcName: "burn",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000", "fail"},
			strToPrint: "",
			expectedError: errors.New("error: 3 arguments does not match required 2"),
			expectedTx: nil,
		},
		{
			testName: "WriteContract func Burn instance failure.",
			funcName: "burn",
			funcArgs: []string{"0x86Be6FC9B05B55CBD04F3161f9b481f27F90a8Df",
				"100000000000"},
			strToPrint: "",
			expectedError: errors.New("error: something bad happened"),
			expectedTx: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			to := common.HexToAddress(tt.funcArgs[0])
			amount := new(big.Int)
			amount.SetString(tt.funcArgs[1], 10)
			mInstance := new(MockContractInstance)
			auth :=  &bind.TransactOpts{}
			// Since this is run after validation it's fine to use only one error var
			mInstance.On("Burn", auth, to, amount).Return(
				tt.expectedTx, tt.expectedError)
			dttc := DetailedTestTokenContract{}
			dttc.Instance = mInstance
			err := dttc.WriteContract(auth, tt.funcName, tt.funcArgs)
			if err != nil {
				assert.Equal(t, err.Error(), tt.expectedError.Error())
			}
			assert.Equal(t, dttc.LastTx, tt.expectedTx)
			assert.Equal(t, dttc.StrToPrint, tt.strToPrint)
		})
	}
}