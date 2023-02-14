#!/bin/bash

KEY="j1"
CHAINID="test-1"
MONIKER="localjack"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"

canined config keyring-backend $KEYRING
canined config chain-id $CHAINID
canined config output "json"

command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

from_scratch () {
    make install

    # remove existing daemon
    rm -rf ~/.canine/*
    
    # jkl1hj5fveer5cjtn4wd6wstzugjfdxzl0xpljur4u '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApZa31BR3NWLylRT6Qi5+f+zXtj2OpqtC76vgkUGLyww"}'
    echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | canined keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
    # j2 jkl1s00nvkagel9xe6luqmmd09jt6jgjl7qu57prct  '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Ah3VzRghgXLn8IA2AH6qaoiuBwZv3ADg3gNPFTo92FwM"}'
    echo "guess census arena parent ribbon among advice green electric almost wink muffin size unfold hedgehog gather warfare embrace float entry cargo ice fade best" | canined keys add j2 --keyring-backend $KEYRING --algo $KEYALGO --recover

    canined init $MONIKER --chain-id $CHAINID 

    # Function updates the config based on a jq argument as a string
    update_test_genesis () {
        # EX: update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
        cat $HOME/.canine/config/genesis.json | jq "$1" > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json
    }

    # Set gas limit in genesis
    update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    update_test_genesis '.app_state["gov"]["voting_params"]["voting_period"]="15s"'
    
    update_test_genesis '.app_state["staking"]["params"]["bond_denom"]="ujkl"'  
    
    update_test_genesis '.app_state["mint"]["params"]["mint_denom"]="ujkl"'  
    update_test_genesis '.app_state["gov"]["deposit_params"]["min_deposit"]=[{"denom": "ujkl","amount": "1000000"}]'
    update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom": "ujkl","amount": "1000"}'

    # Allocate genesis accounts
    canined add-genesis-account $KEY 10000000ujkl --keyring-backend $KEYRING
    canined add-genesis-account j2 10000000ujkl --keyring-backend $KEYRING
    
    canined gentx $KEY 1000000ujkl --keyring-backend $KEYRING --chain-id $CHAINID
    
    canined collect-gentxs
    
    canined validate-genesis
}

from_scratch

# Opens the RPC endpoint to outside connections
sed -i '/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:26657"' ~/.canine/config/config.toml
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' ~/.canine/config/config.toml

# Update genesis file
sed -i 's/"deposit_account": "[a-z0-9]*",/"deposit_account": "jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8",/' ~/.canine/config/genesis.json

# Start the node 
canined start --pruning=nothing  --minimum-gas-prices=0ujkl
