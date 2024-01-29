#!/bin/bash

seed="$1"
address="$2"
index="$3"

jprovd start --home="$HOME/providers/provider$index" -y --port "333$index" --moniker="provider$index" --threads=1 --interval=5 
