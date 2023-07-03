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

    # adding providers to genesis
    canined add-genesis-account jkl1xclg3utp4yuvaxa54r39xzrudc988s82ykve3f 1100000000000ujkl
    canined add-genesis-account jkl1tcveayn80pe3d5wallj9kev3rfefctsmrqf6ks 1100000000000ujkl
    canined add-genesis-account jkl1eg3gm3e3k4dypvvme26ejmajnyvtgwwlaaeu2y 1100000000000ujkl
    canined add-genesis-account jkl1ga0348r8zhn8k4xy3fagwvkwzvyh5lynxr5kak 1100000000000ujkl
    canined add-genesis-account jkl18encuf0esmxv3pxqjqvn0u4tgd6yzuc8urzlp0 1100000000000ujkl
    canined add-genesis-account jkl1sqt9v0zwwx362szrek7pr3lpq29aygw06hgyza 1100000000000ujkl
    canined add-genesis-account jkl1yu099xns2qpslvyrymxq3hwrqhevs7qxksvu8p 1100000000000ujkl

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

sed -i.bak -e 's/stake/ujkl/' $HOME/.canine/config/genesis.json
sed -i.bak -e 's/^minimum-gas-prices =""/minimum-gas-prices = \"0.0025ujkl\"/' $HOME/.canine/config/app.toml
sed -i.bak -e 's/enable = false/enable=true/' $HOME/.canine/config/app.toml
sed -i.bak -e 's/enable=false/enable=true/' $HOME/.canine/config/app.toml
sed -i.bak -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $HOME/.canine/config/app.toml
sed -i.bak -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/' $HOME/.canine/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $HOME/.canine/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26656"/laddr = "tcp:\/\/0.0.0.0:26656"/' $HOME/.canine/config/config.toml
sed -i.bak -e 's/chain-id = ""/chain-id = "canine-1"/' $HOME/.canine/config/client.toml

trap "cleanup" SIGINT

rm "*.log"

# Start the node
screen -d -m -S "canined" bash -c "canined start --pruning=nothing --minimum-gas-prices=0ujkl"

sleep 30

echo "starting providers..."

screen -d -m -L -Logfile "provider0.log" -S "provider0" bash -c "./scripts/start-provider.sh 54f86a701648e8324e920f9592c21cc591b244ae46eac935d45fe962bba1102c jkl1xclg3utp4yuvaxa54r39xzrudc988s82ykve3f 0"
screen -d -m -L -Logfile "provider1.log" -S "provider1" bash -c "./scripts/start-provider.sh a29c5f0033606d1ac47db6a3327bc13a6b0c426dbfe5c15b2fcd7334b4165033 jkl1tcveayn80pe3d5wallj9kev3rfefctsmrqf6ks 1"
screen -d -m -L -Logfile "provider2.log" -S "provider2" bash -c "./scripts/start-provider.sh a490cb438024cddca16470771fb9a21938c4cf61176a46005c6a7b25ee25a649 jkl1eg3gm3e3k4dypvvme26ejmajnyvtgwwlaaeu2y 2"
screen -d -m -L -Logfile "provider3.log" -S "provider3" bash -c "./scripts/start-provider.sh 6c8a948c347079706e404ab48afc5f03203556e34ea921f3b132f2b2e9bcc87d jkl1ga0348r8zhn8k4xy3fagwvkwzvyh5lynxr5kak 3"
screen -d -m -L -Logfile "provider4.log" -S "provider4" bash -c "./scripts/start-provider.sh 8144389a23c6535e276068ff9043b2b6ff95aa3c103c35486c8f2d2363606fd5 jkl18encuf0esmxv3pxqjqvn0u4tgd6yzuc8urzlp0 4"
screen -d -m -L -Logfile "provider5.log" -S "provider5" bash -c "./scripts/start-provider.sh 0e019088a0fafa8f77cb5c0d0f6cb6b63a0015f20d2450480cbcdee44d170aab jkl1sqt9v0zwwx362szrek7pr3lpq29aygw06hgyza 5"
screen -d -m -L -Logfile "provider6.log" -S "provider6" bash -c "./scripts/start-provider.sh adf5a86ac54146b172c20b865c548e900c51439c3723af14aeab668ccd2b8ecf jkl1yu099xns2qpslvyrymxq3hwrqhevs7qxksvu8p 6"

echo "done!"

sleep 60

echo "starting file uploads"


curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/1.png http://localhost:3330/upload

sleep 10

canined tx storage sign-contract jklc107a4hj35fg4jlcapl9h3y7rhw4p24pjtunv8hwwg2hp9dcwatgwsw229ql --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/3.png http://localhost:3331/upload

sleep 10

canined tx storage sign-contract jklc1mnjchx4g27929y5wfm6luvx86lhca688vqh3vtts6hjllfupgvdqgmr04k --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/4.png http://localhost:3332/upload

sleep 10

canined tx storage sign-contract jklc14y6hk074svd8dyjg5g6c2xzkcfv4ge0w2ey96plr98k4pyepk5xq46wjkn --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/test.txt http://localhost:3333/upload

sleep 10

canined tx storage sign-contract jklc102jpqmfj5w9pz555zfjd6e9v5nfcnjsy4vh08qhlmrnedrg5jvwqluwrl8 --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/2.png http://localhost:3334/upload

sleep 10

canined tx storage sign-contract jklc16xdn85h9fel3ruawc3dxjdu7xtkdfp23j66fkkl7dc4277t03rds2rf7g7 --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/5.svg http://localhost:3335/upload

sleep 10

canined tx storage sign-contract jklc1g80djchxzjxztkff98wrelh86ha0alhwu0j9ce84sqvljy8vpddsg637q5 --from charlie -y --pay-upfront

sleep 30

curl -v -F sender=jkl10k05lmc88q5ft3lm00q30qkd9x6654h3lejnct -F file=@./scripts/dummy_data/6.wav http://localhost:3336/upload

sleep 10

canined tx storage sign-contract jklc1sfwyeu8sz7vuhwhej5qclffjrcw99tc6t527y4jnu4t9cd8yc98sm7ufrc --from charlie -y --pay-upfront

sleep 20

read -rsp $'Press any key to continue...\n' -n1 key

killall screen
killall canined
# clean after program termination without SIGINT
cleanup

