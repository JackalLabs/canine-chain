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

echo "$@"

# removing all older jackal inits (if they exist)
rm -rf $HOME/.canine

# making a new jackal directory
mkdir $HOME/.canine

# initializing the validator 
canined init $VALIDATOR_NAME --chain-id=$CHAIN_ID 

# copying the genesis
cp /home/common_store/genesis.json /root/.canine/config/genesis.json

# creating new keys
canined keys add $VALIDATOR_NAME --keyring-backend=test

# adding a genesis account 
canined add-genesis-account \
$(canined keys show $VALIDATOR_NAME -a --keyring-backend=test) \
2000000000ujkl

# creating the gentx
canined gentx $VALIDATOR_NAME 500000000ujkl \
--chain-id=testing \
--pubkey=$(canined tendermint show-validator --chain-id=testing) \
--fees=2500ujkl \
--commission-max-change-rate=0.01 \
--commission-max-rate=0.20 \
--commission-rate=0.05 \
--keyring-backend test \

# copying the modified genesis back
cp /root/.canine/config/genesis.json /home/common_store/genesis.json

# copying the gentxs
cp -RT /root/.canine/config/gentx /home/common_store/gentx