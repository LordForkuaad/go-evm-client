privateKey=$(ethermintd keys unsafe-export-eth-key mykey --keyring-backend test)
echo $privateKey
go run cmd/contract_deployer/main.go -p $privateKey -r "http://127.0.0.1:8545" -c "detailed_test_token" -a "MintSwapToken" -a "MST" -a "100000000000000000000000000"
