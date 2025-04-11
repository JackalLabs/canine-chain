#!/bin/bash

OLD_VERSION=$1
NEW_VERSION=$2
UPGRADE_HEIGHT=35
HOME=mytestnet
ROOT=$(pwd)
DENOM=ujkl
SOFTWARE_UPGRADE_NAME=$3

# underscore so that go tool will not take gocache into account
mkdir -p ${ROOT}/../_build/gocache
export GOMODCACHE=${ROOT}/../_build/gocache

# install old binary

mkdir -p ../_build/old
git checkout $OLD_VERSION
go build -mod=readonly -o build/canined ./cmd/canined
mv build/canined ./../_build/old/canined
git checkout $NEW_VERSION
go install -mod=readonly ./cmd/canined

# start old node
screen -L -Logfile upgrade-node.log  -dmS node1 bash scripts/run-upgrade-node.sh ./../_build/old/canined $DENOM

sleep 30

./../_build/old/canined version --home $HOME
./../_build/old/canined config broadcast-mode block --home $HOME

sleep 35

./../_build/old/canined q storage params --home $HOME

./../_build/old/canined tx storage buy-storage jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg 720h 1000000000 ujkl --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 7

./../_build/old/canined provider init http://localhost:3333 10000000 ""  --from test1 --keyring-backend test --chain-id test --home $HOME -y
./../_build/old/canined provider init http://localhost:3333 10000000 ""  --from test2 --keyring-backend test --chain-id test --home $HOME -y

sleep 7

./../_build/old/canined tx rns init --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 7

./../_build/old/canined tx gov submit-proposal software-upgrade "$SOFTWARE_UPGRADE_NAME" --upgrade-height $UPGRADE_HEIGHT --upgrade-info "temp" --title "upgrade" --description "upgrade"  --from test1 --keyring-backend test --chain-id test --home $HOME -y --deposit "20000000${DENOM}"

sleep 7

./../_build/old/canined tx gov vote 1 yes --from test --keyring-backend test --chain-id test --home $HOME -y

sleep 7

./../_build/old/canined tx gov vote 1 yes --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 12

./../_build/old/canined tx storage post-contract 1688dc719d1a41ff567fd54e66953f5c518044f6fed6ce814ba777b7dead4ab7d1c193448dc1c04eac05e6708dfd7a8999e9afdf6ba5c525ab7fb9c7f1e2bd4c jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg 10000 jklf1p5cm3z47rrcyaskqge3yc33xm7hdq7lken99ahluvuz67ugctleqmwv43a --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 7

./../_build/old/canined q storage list-contracts --home $HOME

echo "signing contract"

./../_build/old/canined tx storage sign-contract jklc1hft5yqlqermu9l337et6mn6ljs7x9tuqgv2elatlk3w8s8w5uyls3ktp22 --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 7


# determine block_height to halt
while true; do 
    BLOCK_HEIGHT=$(./../_build/old/canined status | jq '.SyncInfo.latest_block_height' -r)
    if [ $BLOCK_HEIGHT = "$UPGRADE_HEIGHT" ]; then
        # assuming running only 1 canined
        echo "BLOCK HEIGHT = $UPGRADE_HEIGHT REACHED, KILLING OLD ONE"
        pkill canined
        break
    else
        ./../_build/old/canined q storage list-active-deals --chain-id test --home $HOME
        ./../_build/old/canined q storage list-strays --chain-id test --home $HOME
        ./../_build/old/canined q storage list-contracts --chain-id test --home $HOME
        ./../_build/old/canined q gov proposal 1 --output=json | jq ".status"
        echo "BLOCK_HEIGHT = $BLOCK_HEIGHT"
        sleep 2
    fi
done

sleep 3
canined start --home $HOME

