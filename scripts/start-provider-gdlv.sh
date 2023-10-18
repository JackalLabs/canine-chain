#!/bin/bash
# for debug use
# required: gdlv [https://github.com/aarzilli/gdlv]

command -v gdlv > /dev/null 2>&1 || { echo >&2 "gdlv not installed. Get it here: https://github.com/aarzilli/gdlv"; exit 1; }

seed="$1"
address="$2"
index="$3"

rm -rf "$HOME/providers/provider$index"
jprovd client gen-key --home="$HOME/providers/provider$index"
echo "{\"key\":\"$seed\",\"address\":\"$address\"}" > "$HOME/providers/provider$index/config/priv_storkey.json"

sleep 10
jprovd client balance --home="$HOME/providers/provider$index"
jprovd client config chain-id "test-1" --home="$HOME/providers/provider$index"
jprovd init "http://127.0.0.1:333$index" "1000000000" "" --home="$HOME/providers/provider$index" -y

sleep 20

# change this to where your providers main.go location:
cd $HOME/Workplace/canine-provider/jprov/jprovd
gdlv debug start --home="$HOME/providers/provider$index" -y --port "333$index" --moniker="provider$index" --threads=1 --interval=5 &

# wait for gdlv to start running the code
sleep 20
