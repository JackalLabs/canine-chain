#!/bin/bash
rm -rf $HOME/.canine/


# make four jackal directories
mkdir $HOME/.canine
mkdir $HOME/.canine/validator1
mkdir $HOME/.canine/validator2
mkdir $HOME/.canine/validator3

# init all three validators
canined init --chain-id=testing validator1 --home=$HOME/.canine/validator1
canined init --chain-id=testing validator2 --home=$HOME/.canine/validator2
canined init --chain-id=testing validator3 --home=$HOME/.canine/validator3
# create keys for all three validators
canined keys add validator1 --keyring-backend=test --home=$HOME/.canine/validator1
canined keys add validator2 --keyring-backend=test --home=$HOME/.canine/validator2
canined keys add validator3 --keyring-backend=test --home=$HOME/.canine/validator3

# change staking denom to ujkl
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="ujkl"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
canined add-genesis-account $(canined keys show validator1 -a --keyring-backend=test --home=$HOME/.canine/validator1) 100000000000ujkl,100000000000stake --home=$HOME/.canine/validator1
canined gentx validator1 500000000ujkl --keyring-backend=test --home=$HOME/.canine/validator1 --chain-id=testing
canined collect-gentxs --home=$HOME/.canine/validator1


# update staking genesis
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json

# update crisis variable to ujkl
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="ujkl"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json

# udpate gov genesis
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ujkl"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json

# update mint genesis
cat $HOME/.canine/validator1/config/genesis.json | jq '.app_state["jklmint"]["params"]["mintDenom"]="ujkl"' > $HOME/.canine/validator1/config/tmp_genesis.json && mv $HOME/.canine/validator1/config/tmp_genesis.json $HOME/.canine/validator1/config/genesis.json


# port key (validator1 uses default ports)
# validator1 1317, 9090, 9091, 26658, 26657, 26656, 6060
# validator2 1316, 9088, 9089, 26655, 26654, 26653, 6061
# validator3 1315, 9086, 9087, 26652, 26651, 26650, 6062
# validator4 1314, 9084, 9085, 26649, 26648, 26647, 6063


# change app.toml values

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $HOME/.canine/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.canine/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.canine/validator2/config/app.toml

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $HOME/.canine/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.canine/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.canine/validator3/config/app.toml


# change config.toml values

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.canine/validator1/config/config.toml
# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $HOME/.canine/validator2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $HOME/.canine/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $HOME/.canine/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.canine/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.canine/validator2/config/config.toml
# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $HOME/.canine/validator3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $HOME/.canine/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.canine/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.canine/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.canine/validator3/config/config.toml


# copy validator1 genesis file to validator2-3
cp $HOME/.canine/validator1/config/genesis.json $HOME/.canine/validator2/config/genesis.json
cp $HOME/.canine/validator1/config/genesis.json $HOME/.canine/validator3/config/genesis.json


# copy tendermint node id of validator1 to persistent peers of validator2-3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(canined tendermint show-node-id --home=$HOME/.canine/validator1)@$(curl -4 icanhazip.com):26656\"|g" $HOME/.canine/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(canined tendermint show-node-id --home=$HOME/.canine/validator1)@$(curl -4 icanhazip.com):26656\"|g" $HOME/.canine/validator3/config/config.toml


# start all three validators
screen -S validator1 canined start --home=$HOME/.canine/validator1
screen -S validator2 canined start --home=$HOME/.canine/validator2
screen -S validator3 canined start --home=$HOME/.canine/validator3


# send ujkl from first validator to second validator
sleep 7
canined tx bank send validator1 $(canined keys show validator2 -a --keyring-backend=test --home=$HOME/.canine/validator2) 500000000ujkl --keyring-backend=test --home=$HOME/.canine/validator1 --chain-id=testing --yes
sleep 7
canined tx bank send validator1 $(canined keys show validator3 -a --keyring-backend=test --home=$HOME/.canine/validator3) 400000000ujkl --keyring-backend=test --home=$HOME/.canine/validator1 --chain-id=testing --yes

# create second validator
sleep 7
canined tx staking create-validator --amount=500000000ujkl --from=validator2 --pubkey=$(canined tendermint show-validator --home=$HOME/.canine/validator2) --moniker="validator2" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="500000000" --keyring-backend=test --home=$HOME/.canine/validator2 --yes
sleep 7
canined tx staking create-validator --amount=400000000ujkl --from=validator3 --pubkey=$(canined tendermint show-validator --home=$HOME/.canine/validator3) --moniker="validator3" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="400000000" --keyring-backend=test --home=$HOME/.canine/validator3 --yes
