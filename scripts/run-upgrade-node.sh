#!/bin/bash

rm -rf mytestnet

BINARY=$1
DENOM=$2
HOME=mytestnet
CHAIN_ID="test"
KEYRING="test"
KEY="test"
KEY1="test1"
KEY2="test2"

# Function updates the config based on a jq argument as a string
update_test_genesis () {
    # EX: update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    cat $HOME/config/genesis.json | jq --arg DENOM "$2" "$1" > $HOME/config/tmp_genesis.json && mv $HOME/config/tmp_genesis.json $HOME/config/genesis.json
}

$BINARY init --chain-id $CHAIN_ID moniker --home $HOME

$BINARY keys add $KEY --keyring-backend $KEYRING --home $HOME

echo "brief enhance flee chest rabbit matter chaos clever lady enable luggage arrange hint quarter change float embark canoe chalk husband legal dignity music web" | $BINARY keys add $KEY1 --keyring-backend $KEYRING --recover --home $HOME
echo "bulk whisper now clump write donor jump menu option muffin crack absent angle dumb deposit empower calm across dawn slice simple crisp soon oak" | $BINARY keys add $KEY2 --keyring-backend $KEYRING --recover --home $HOME

# Allocate genesis accounts (cosmos formatted addresses)
$BINARY add-genesis-account $KEY "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME

$BINARY add-genesis-account $KEY1 "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME
$BINARY add-genesis-account $KEY2 "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME

update_test_genesis '.app_state["gov"]["voting_params"]["voting_period"] = "50s"'
update_test_genesis '.app_state["mint"]["params"]["mint_denom"]=$DENOM' $DENOM
update_test_genesis '.app_state["gov"]["deposit_params"]["min_deposit"]=[{"denom": $DENOM,"amount": "1000000"}]' $DENOM
update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom": $DENOM,"amount": "1000"}' $DENOM
update_test_genesis '.app_state["staking"]["params"]["bond_denom"]=$DENOM' $DENOM
update_test_genesis '.app_state["storage"]["params"]["misses_to_burn"]=2' $DENOM
update_test_genesis '.app_state["storage"]["params"]["proof_window"]=3' $DENOM
update_test_genesis '.app_state["storage"]["params"]["deposit_account"]="jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg"' $DENOM


sed -i '' 's/enable = false/enable = true/' $HOME/config/app.toml
sed -i '' 's/pruning = "default"/pruning = "nothing"/' $HOME/config/app.toml

# Sign genesis transaction
$BINARY gentx $KEY "1000000${DENOM}" --keyring-backend $KEYRING --chain-id $CHAIN_ID --home $HOME

# Collect genesis tx
$BINARY collect-gentxs --home $HOME

# Run this to ensure everything worked and that the genesis file is setup correctly
$BINARY validate-genesis --home $HOME

$BINARY start --home $HOME