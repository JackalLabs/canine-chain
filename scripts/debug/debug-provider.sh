#!/bin/bash

seed="$1"
address="$2"
index="$3"

command -v gdlv > /dev/null 2>&1 || { echo >&2 "gdlv not installed. More info https://github.com/aarzilli/gdlv"; exit 1; }

localRepo=$(dirname "$(find $HOME -wholename "*/canine-provider/.git")")
if [[ ${#localRepo} -eq 0 ]]; then
    echo "canine-provider repository is not found, can't open gdlv without it"
fi

cd ${localRepo}

rm -rf "$HOME/providers/provider$index"
jprovd client gen-key --home="$HOME/providers/provider$index"
echo "{\"key\":\"$seed\",\"address\":\"$address\"}" > "$HOME/providers/provider$index/config/priv_storkey.json"

sleep 10
jprovd client balance --home="$HOME/providers/provider$index"
jprovd client config chain-id "test-1" --home="$HOME/providers/provider$index"
jprovd init "http://127.0.0.1:333$index" "1000000000" "" --home="$HOME/providers/provider$index" -y

sleep 20


cd -- "$(dirname "$(find . -name "main.go")")"
gdlv debug start --home="$HOME/providers/provider$index" -y --port "333$index" --moniker="provider$index" --threads=1 --interval=5

