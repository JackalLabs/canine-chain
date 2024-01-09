#!/bin/bash

# devnet for ica project ensuring compatibility with strangelove's (sl) e2e test suite - specifically how 
# the suite builds cosmos chains and sets up the chains config files

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

# sl test suite sets this path for a cosmos-chain's config files. All commands
# are executed with this path as the home flag 
# NOTE: need to run script in clean linux sandbox to ensure this path is the only
# path which contains canined config files 
CANINE_HOME="/var/cosmos-chain/canined-2"