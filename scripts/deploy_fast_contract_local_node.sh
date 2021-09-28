privateKey=$(ethermintd keys unsafe-export-eth-key mykey --keyring-backend test)
echo $privateKey
go run cmd/contract_deployer/main.go -p $privateKey -r "http://127.0.0.1:8545" -c "fast_test_token"
