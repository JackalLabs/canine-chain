#!/bin/bash

OLD_VERSION=23.01-beta
UPGRADE_HEIGHT=20
HOME=mytestnet
ROOT=$(pwd)
DENOM=ujkl
SOFTWARE_UPGRADE_NAME=fixstrays

# underscore so that go tool will not take gocache into account
mkdir -p ${ROOT}/../_build/gocache
export GOMODCACHE=${ROOT}/../_build/gocache

# install old binary
if ! command -v ./../_build/old/canined &> /dev/null
then
    mkdir -p ../_build/old
    wget -c "https://github.com/JackalLabs/canine-chain/archive/refs/tags/v${OLD_VERSION}.zip" -O ../_build/v${OLD_VERSION}.zip
    unzip ../_build/v${OLD_VERSION}.zip -d ../_build
    cd ../_build/canine-chain-${OLD_VERSION}
      make build
      mv build/canined ../old/canined
    cd ../..
fi


# start old node
screen -dmS node1 bash scripts/run-upgrade-node.sh ./../_build/old/canined $DENOM

./../_build/old/canined version --home $HOME

sleep 20

./../_build/old/canined tx gov submit-proposal software-upgrade "$SOFTWARE_UPGRADE_NAME" --upgrade-height $UPGRADE_HEIGHT --upgrade-info "temp" --title "upgrade" --description "upgrade"  --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./../_build/old/canined tx gov deposit 1 "20000000${DENOM}" --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./../_build/old/canined tx gov vote 1 yes --from test --keyring-backend test --chain-id test --home $HOME -y

sleep 3

./../_build/old/canined tx gov vote 1 yes --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 3

# determine block_height to halt
while true; do 
    BLOCK_HEIGHT=$(./../_build/old/canined status | jq '.SyncInfo.latest_block_height' -r)
    if [ $BLOCK_HEIGHT = "$UPGRADE_HEIGHT" ]; then
        # assuming running only 1 canined
        echo "BLOCK HEIGHT = $UPGRADE_HEIGHT REACHED, KILLING OLD ONE"
        pkill canined
        break
    else
        ./../_build/old/canined q gov proposal 1 --output=json | jq ".status"
        echo "BLOCK_HEIGHT = $BLOCK_HEIGHT"
        sleep 10
    fi
done

#sleep 3

#canined start --log_level debug --home $HOME