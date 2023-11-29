#!/bin/bash

KEY="j1"
KEY1="j2"
KEY2="charlie"
DEPOACCKEY="deposit_account"

CHAINID="test-1"
MONIKER="localjack"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
BROADCASTMODE="block"

canined config keyring-backend $KEYRING
canined config chain-id $CHAINID
canined config broadcast-mode $BROADCASTMODE
canined config output "json"

command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

from_scratch () {
    make install

    # remove existing daemon
    rm -rf ~/.canine/*
    
    # jkl1hj5fveer5cjtn4wd6wstzugjfdxzl0xpljur4u '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApZa31BR3NWLylRT6Qi5+f+zXtj2OpqtC76vgkUGLyww"}'
    echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | canined keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
    # j2 jkl1s00nvkagel9xe6luqmmd09jt6jgjl7qu57prct  '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Ah3VzRghgXLn8IA2AH6qaoiuBwZv3ADg3gNPFTo92FwM"}'
    echo "guess census arena parent ribbon among advice green electric almost wink muffin size unfold hedgehog gather warfare embrace float entry cargo ice fade best" | canined keys add $DEPOACCKEY --keyring-backend $KEYRING --algo $KEYALGO --recover

	echo "video pluck level diagram maximum grant make there clog tray enrich book hawk confirm spot you book vendor ensure theory sure jewel sort basket" | canined keys add $KEY1 --algo $KEYALGO --keyring-backend $KEYRING --recover

	echo "flock stereo dignity lawsuit mouse page faith exact mountain clinic hazard parent arrest face couch asset jump feed benefit upper hair scrap loud spirit" | canined keys add $KEY2 --algo $KEYALGO --keyring-backend $KEYRING --recover

    canined init $MONIKER --chain-id $CHAINID 

	canined config keyring-backend $KEYRING
	canined config chain-id $CHAINID
	canined config broadcast-mode $BROADCASTMODE

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

	# Use jkl bech32 prefix account for storage and oracle modules
	update_test_genesis '.app_state["storage"]["params"]["deposit_account"]="'"$(canined keys show -a $DEPOACCKEY)"'"'
    update_test_genesis '.app_state["storage"]["params"]["chunk_size"]="'10240'"'
    update_test_genesis '.app_state["storage"]["params"]["proof_window"]="'25'"'
    update_test_genesis '.app_state["oracle"]["params"]["deposit"]="'"$(canined keys show -a $DEPOACCKEY)"'"'
    update_test_genesis '.app_state["rns"]["params"]["deposit_account"]="'"$(canined keys show -a $DEPOACCKEY)"'"'

    update_test_genesis '.app_state["interchainaccounts"]["host_genesis_state"]["params"]["allow_messages"]=["*"] | map(.)'

    # Allocate genesis accounts
    canined add-genesis-account $KEY 1000000000000ujkl --keyring-backend $KEYRING
    canined add-genesis-account $DEPOACCKEY 10000000000ujkl  --keyring-backend $KEYRING
    canined add-genesis-account $KEY1 10000000000ujkl  --keyring-backend $KEYRING
    canined add-genesis-account $KEY2 10000000000ujkl  --keyring-backend $KEYRING
    
    canined gentx $KEY 1000000ujkl --keyring-backend $KEYRING --chain-id $CHAINID
    
    canined collect-gentxs
    
    canined validate-genesis
}

startup() {
	mv $HOME/.canine $HOME/.canine.old
}

cleanup() {
	echo "SIGINT captured, starting cleanup"
	mv $HOME/.canine.old $HOME/.canine
	exit
}

startup

from_scratch

# Opens the RPC endpoint to outside connections
sed -i '/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:26657"' ~/.canine/config/config.toml
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' ~/.canine/config/config.toml

trap "cleanup" SIGINT

# Start the node 
canined start --pruning=nothing  --minimum-gas-prices=0ujkl

# clean after program termination without SIGINT
cleanup
