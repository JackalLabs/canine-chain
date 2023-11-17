export JKL_HOME="$HOME/canine-test"
export PROV_HOME="$HOME/canine-prov-test"

rm -rf $JKL_HOME
rm -rf $PROV_HOME

export CHAIN="canine-1"
export ALIAS="marston"
export MONIKER="jackal"

canined init $MONIKER --home=$JKL_HOME --chain-id=$CHAIN
canined config chain-id $CHAIN --home=$JKL_HOME
canined config keyring-backend test --home=$JKL_HOME

sed -i.bak -e 's/chain-id = ""/chain-id = "canine-1"/' $JKL_HOME/config/client.toml

echo "video pluck level diagram maximum grant make there clog tray enrich book hawk confirm spot you book vendor ensure theory sure jewel sort basket" | canined keys add $ALIAS --keyring-backend=test --recover --home=$JKL_HOME
echo "flock stereo dignity lawsuit mouse page faith exact mountain clinic hazard parent arrest face couch asset jump feed benefit upper hair scrap loud spirit" | canined keys add charlie --keyring-backend=test --recover --home=$JKL_HOME
echo "brief enhance flee chest rabbit matter chaos clever lady enable luggage arrange hint quarter change float embark canoe chalk husband legal dignity music web" | canined keys add danny --keyring-backend=test --recover --home=$JKL_HOME

canined add-genesis-account $(canined keys show -a $ALIAS --keyring-backend=test --home=$JKL_HOME) 500000000ujkl --home=$JKL_HOME
canined add-genesis-account $(canined keys show -a charlie --keyring-backend=test --home=$JKL_HOME) 500000000ujkl --home=$JKL_HOME
canined add-genesis-account $(canined keys show -a danny --keyring-backend=test --home=$JKL_HOME) 500000000ujkl --home=$JKL_HOME
canined add-genesis-account jkl13327d3ntsy849s622y4ft05ynwkfyqss3v4pzg 500000000000000ujkl --home=$JKL_HOME

canined gentx $ALIAS 200000000ujkl \
--chain-id=$CHAIN \
--moniker="$MONIKER" \
--commission-max-change-rate=0.01 \
--commission-max-rate=0.20 \
--commission-rate=0.05 \
--fees=2500ujkl \
--from=$ALIAS \
--keyring-backend=test \
--home=$JKL_HOME

canined collect-gentxs --home=$JKL_HOME

sed -i.bak -e "s/stake/ujkl/" $JKL_HOME/config/genesis.json
sed -i.bak -e "s/cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8/jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgc3sexc/" $JKL_HOME/config/genesis.json
sed -i.bak -e "s/\"proof_window\": \"50\"/\"proof_window\": \"200\"/" $JKL_HOME/config/genesis.json
sed -i.bak -e "s/\"check_window\": \"100\"/\"check_window\": \"20\"/" $JKL_HOME/config/genesis.json
sed -i.bak -e "s/\"chunk_size\": \"1024\"/\"chunk_size\": \"1024\"/" $JKL_HOME/config/genesis.json
sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"0.0025ujkl\"/" $JKL_HOME/config/app.toml
sed -i.bak -e 's/enable = false/enable=true/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/' $JKL_HOME/config/config.toml
sed -i.bak -e 's/chain-id = ""/chain-id = "canine-1"/' $JKL_HOME/config/client.toml
screen -d -m canined start --home=$JKL_HOME --log_level info
sleep 20
canined tx storage buy-storage $(canined keys show charlie -a --home=$JKL_HOME) 30 1000000000000 ujkl --from charlie --gas-prices=0.02ujkl --home=$JKL_HOME -y

sequoia init --home=$PROV_HOME

yq -i '.proof_interval=120' $PROV_HOME/config.yaml
yq -i '.queue_interval=7' $PROV_HOME/config.yaml
yq -i '.chain_config.rpc_addr="http://localhost:26657"' $PROV_HOME/config.yaml
yq -i '.chain_config.grpc_addr="localhost:9090"' $PROV_HOME/config.yaml
yq -i '.domain="http://localhost:3334"' $PROV_HOME/config.yaml
yq -i '.data_directory="'$PROV_HOME'/data"' $PROV_HOME/config.yaml
yq -i '.api_config.port=3334' $PROV_HOME/config.yaml

rm $PROV_HOME/provider_wallet.json
echo "{\"seed_phrase\":\"forward service profit benefit punch catch fan chief jealous steel harvest column spell rude warm home melody hat broccoli pulse say garlic you firm\",\"derivation_path\":\"m/44'/118'/0'/0/0\"}" > $PROV_HOME/provider_wallet.json
sequoia start --home=$PROV_HOME
