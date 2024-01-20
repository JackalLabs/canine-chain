#!/bin/bash

localRepo=$(dirname "$(find $HOME -wholename "*/canine-chain/.git")")
if [[ ${#localRepo} -eq 0 ]];then
    echo "canine-chain repository is not found, can't open gdlve without it"
fi

cd ${localRepo}

echo "starting debugger for chain..."
cd -- "$(dirname "$(find . -name "main.go")")"
gdlv debug start --pruning=nothing --minimum-gas-prices=0ujkl &

read -rsp $'Press anything to continue test... \n' -n1 key
