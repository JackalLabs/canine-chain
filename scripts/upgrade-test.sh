#!/bin/bash

OLD_VERSION=$1
NEW_VERSION=$2
UPGRADE_HEIGHT=30
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
make build
mv build/canined ./../_build/old/canined
git checkout $NEW_VERSION
make install

# start old node
screen -dmS node1 bash scripts/run-upgrade-node.sh ./../_build/old/canined $DENOM

sleep 30

./../_build/old/canined version --home $HOME
./../_build/old/canined config broadcast-mode block --home $HOME

sleep 30

./../_build/old/canined q storage params --home $HOME

./../_build/old/canined tx storage buy-storage jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg 30 1000000000 ujkl --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined provider init http://localhost:3333 10000000 ""  --from test1 --keyring-backend test --chain-id test --home $HOME -y
./../_build/old/canined provider init http://localhost:3333 10000000 ""  --from test2 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx rns init --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx storage post-contract jklc10zpctnu6qu4dkvzx2u9jpvk6hr8z3qnfpzlmf63ejqdgjgr5h5qszdqylf jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg 10000 jklf16yfjl9t8u9ztt0e5wzudpfr3e2u5cxwsru6l6jh930zrvav5cz2q2wrfkc --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

echo "signing contract"
./../_build/old/canined tx storage sign-contract jklc1p8u0a9yhqc9vr2rq2ksz8jekkq002y4chcv4k7thw5ntaqgdrvss5mqx3s --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx gov submit-proposal software-upgrade "$SOFTWARE_UPGRADE_NAME" --upgrade-height $UPGRADE_HEIGHT --upgrade-info "temp" --title "upgrade" --description "upgrade"  --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx gov deposit 1 "20000000${DENOM}" --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx gov vote 1 yes --from test --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx gov vote 1 yes --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx storage post-contract jklc1hufzy29uulvfcpw9xsat0fs87vglyslj8sc0cz270necl9u23hzsczp0s3 jkl12g4qwenvpzqeakavx5adqkw203s629tf6k8vdg 10000 jklf1rqry0q7a55tanxkv34rnza82pewfm292pr77m78vc8avjk5p3e9sc6qgnq --from test1 --keyring-backend test --chain-id test --home $HOME -y

sleep 6

./../_build/old/canined tx storage sign-contract jklc1hufzy29uulvfcpw9xsat0fs87vglyslj8sc0cz270necl9u23hzsczp0s3 --from test1 --keyring-backend test --chain-id test --home $HOME -y


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

