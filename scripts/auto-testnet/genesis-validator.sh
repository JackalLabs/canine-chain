#!/bin/bash
set -e

# collecting keywords
# required keywords: VALIDATOR_NAME,  CHAIN_ID, 
for ARGUMENT in "$@"
do
   KEY=$(echo $ARGUMENT | cut -f1 -d=)

   KEY_LENGTH=${#KEY}
   VALUE="${ARGUMENT:$KEY_LENGTH+1}"

   export "$KEY"="$VALUE"
done

# removing all older jackal inits (if they exist)
rm -rf $HOME/.canine

# making a new jackal directory
mkdir $HOME/.canine

# initializing the validator 
canined init $VALIDATOR_NAME --chain-id=$CHAIN_ID  --home=$HOME/.canine

# creating new keys
canined keys add $VALIDATOR_NAME --keyring-backend=test --home=$HOME/.canine

# create validator node with tokens to transfer to the three other nodes
canined add-genesis-account $(canined keys show $VALIDATOR_NAME -a --keyring-backend=test --home=$HOME/.canine) \
2000000000ujkl --home=$HOME/.canine 

# creating a new gentx
canined gentx $VALIDATOR_NAME 500000000ujkl \
--chain-id=testing \
--pubkey=$(canined tendermint show-validator --chain-id=testing) \
--fees=2500ujkl \
--commission-max-change-rate=0.01 \
--commission-max-rate=0.20 \
--commission-rate=0.05 \
--keyring-backend=test \

# update staking genesis
cat $HOME/.canine/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="ujkl"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json
cat $HOME/.canine/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json

# update crisis variable to ujkl
cat $HOME/.canine/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="ujkl"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json

# update gov genesis
cat $HOME/.canine/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json
cat $HOME/.canine/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ujkl"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json

# update mint genesis
cat $HOME/.canine/config/genesis.json | jq '.app_state["jklmint"]["params"]["mintDenom"]="ujkl"' > $HOME/.canine/config/tmp_genesis.json && mv $HOME/.canine/config/tmp_genesis.json $HOME/.canine/config/genesis.json

# copying the genesis file to the common store
cp /root/.canine/config/genesis.json /home/common_store/genesis.json

# copying gentx to the common store
cp -RT /root/.canine/config/gentx /home/common_store/gentx