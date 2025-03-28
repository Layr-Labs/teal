#!/bin/bash

function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo $contract
    cd $contract_dir
    forge b
    # pop back to the original directory
    cd -
    mkdir -p $binding_dir/${contract}
    contract_json="$contract_dir/out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')
    solc_bin=$(cat ${contract_json} | jq -r '.bytecode.object')

    mkdir -p data
    echo ${solc_abi} > data/tmp.abi
    echo ${solc_bin} > data/tmp.bin

    rm -f $binding_dir/${contract}/binding.go
    abigen --bin=data/tmp.bin --abi=data/tmp.abi --pkg=contract${contract} --out=$binding_dir/${contract}/binding.go

    rm -f data/tmp.abi
    rm -f data/tmp.bin
}

contracts="MinimalCertificateVerifier"
for contract in $contracts; do
    create_binding . $contract ./bindings
done