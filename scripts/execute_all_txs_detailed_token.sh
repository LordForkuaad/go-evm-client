#!/bin/bash
# USAGE -p Private Key -a Contract address
declare -A flags
declare -A booleans
args=()

while [ "$1" ];
do
    arg=$1
    if [ "${1:0:1}" == "-" ]
    then
      shift
      rev=$(echo "$arg" | rev)
      if [ -z "$1" ] || [ "${1:0:1}" == "-" ] || [ "${rev:0:1}" == ":" ]
      then
        bool=$(echo ${arg:1} | sed s/://g)
        booleans[$bool]=true
        echo \"$bool\" is boolean
      else
        value=$1
        flags[${arg:1}]=$value
        shift
        echo \"$arg\" is flag with value \"$value\"
      fi
    else
      args+=("$arg")
      shift
      echo \"$arg\" is an arg
    fi
done

go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "name"
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "symbol"
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "decimals"
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "totalsupply"
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "balanceof" -fa ${flags["public"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "allowance" -fa ${flags["public"]} -fa ${flags["recipient"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "transfer" -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "approve" -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "transferfrom" -fa ${flags["public"]} -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "increaseallowance" -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "decreaseallowance" -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "mint" -fa ${flags["recipient"]} -fa ${flags["amount"]}
go run cmd/contract_interactor/main.go -p ${flags["private"]} -r "http://127.0.0.1:8545" -c "detailed_test_token" -a ${flags["address"]} -f "burn" -fa ${flags["recipient"]} -fa ${flags["amount"]}
