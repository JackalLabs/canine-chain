export JKL_HOME="$HOME/canine-test"

rm -rf $JKL_HOME
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
canined add-genesis-account $1 500000000ujkl --home=$JKL_HOME

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

sed -i.bak -e 's/stake/ujkl/' $JKL_HOME/config/genesis.json
sed -i.bak -e 's/^minimum-gas-prices =""/minimum-gas-prices = \"0.0025ujkl\"/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/enable = false/enable=true/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/enable=false/enable=true/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $JKL_HOME/config/app.toml
sed -i.bak -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/' $JKL_HOME/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $JKL_HOME/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26656"/laddr = "tcp:\/\/0.0.0.0:26656"/' $JKL_HOME/config/config.toml
sed -i.bak -e 's/chain-id = ""/chain-id = "canine-1"/' $JKL_HOME/.canine/config/client.toml

canined start --home=$JKL_HOME --log_level info