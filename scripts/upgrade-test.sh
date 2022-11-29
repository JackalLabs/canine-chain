#!/bin/bash

OLD_VERSION=1.1.2-hotfix
UPGRADE_HEIGHT=20
HOME=mytestnet
ROOT=$(pwd)
DENOM=ujkl
SOFTWARE_UPGRADE_NAME=v3

GOMODCACHE=~/go/pkg/mod

# install old binary
if ! command -v build/old/canined &> /dev/null
then
    mkdir -p build/old
    wget -c "https://github.com/JackalLabs/canine-chain/archive/refs/tags/v${OLD_VERSION}.zip" -O build/v${OLD_VERSION}.zip
    unzip build/v${OLD_VERSION}.zip -d build
    cd ./build/canine-chain-${OLD_VERSION}
    GOBIN="$ROOT/build/old" go install -mod=readonly ./...
    cd ../..
fi

# install new binary
if ! command -v build/new/canined &> /dev/null
then
    GOBIN="$ROOT/build/new" go install -mod=readonly ./... 2> /dev/null
fi

# start old node
screen -dmS node1 bash scripts/run-upgrade-node.sh build/old/canined $DENOM

sleep 20

./build/old/canined tx gov submit-proposal software-upgrade "$SOFTWARE_UPGRADE_NAME" --upgrade-height $UPGRADE_HEIGHT --upgrade-info "temp" --title "upgrade" --description "upgrade"  --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./build/old/canined tx gov deposit 1 "20000000${DENOM}" --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./build/old/canined tx gov vote 1 yes --from test --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./build/old/canined tx gov vote 1 yes --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

# determine block_height to halt
while true; do 
    BLOCK_HEIGHT=$(./build/new/canined status | jq '.SyncInfo.latest_block_height' -r)
    if [ $BLOCK_HEIGHT = "$UPGRADE_HEIGHT" ]; then
        # assuming running only 1 canined
        echo "BLOCK HEIGHT = $UPGRADE_HEIGHT REACHED, KILLING OLD ONE"
        break
    else
        ./build/old/canined q gov proposal 1 --output=json | jq ".status"
        echo "BLOCK_HEIGHT = $BLOCK_HEIGHT"
        sleep 10
    fi
done

sleep 3

./build/new/canined start --log_level debug --home $HOME