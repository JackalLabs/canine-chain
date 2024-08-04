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

# Check and install dependencies
check_install_dependency() {
    if ! command -v $1 &> /dev/null
    then 
        echo "$1 could not be found, attempting to install."
        apt-get install -y $1
    fi 
}

check_install_dependency "jq"
check_install_dependency "make"

from_scratch() {
    make install 

    # should be no need to remove existing daemon because a clean docker container is booted up each time 
    
    # sl suite already makes accounts
    # sl suite already inits a moniker, keyring-backend and chain-id. Not sure it inits the correct broadcast mode. 

    canined config broadcast-mode $BROADCASTMODE
    canined config output "json"

    # Function updates the config based on a jq argument as a string 
    # CONCERN: 
    # When the image is pulled and run, does our script run before or after sl suite starts executing canined genesis commands?
    # What if this happens:
        # devnet-ica.sh inits config viles in /var/cosmos-chain/canined-2
        # sl suite starts executing canined genesis commands and over writes everything inside of /var/cosmos-chain/canined-2
        # which means the allow_messages param will be overwritten to be empty []
    # SOLUTION: if needed, look at the sl suite to see how canined's commands are executed 
    # I wonder if it's possible to set the allow_messages param at genesis time through sl's suite--by forking the repo. 

    # Function updates the config based on a jq argument as a string
    update_test_genesis () {
        # EX: update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
        GENESIS_PATH="/var/cosmos-chain/canined-2/config/genesis.json"
        TMP_GENESIS_PATH="/var/cosmos-chain/canined-2/config/tmp_genesis.json"

        # Not sure at this moment if sl suite creates the genesis file before or after this script is run
        # so we loop until the genesis.json file is found

        # NOTE: not sure if this loop will block sl suite from executing canined so we can try running devnet-ica.sh in the background 
        # of our docker container--even though this violates the 'one process per container philosophy'
        while [ ! -f "$GENESIS_PATH" ]; do
            echo "Waiting for genesis.json to be created..."
            sleep 5 # Waits for 5 seconds before checking again
        done

        echo "genesis.json found. Updating file..." 

        # Now proceed with the update 
        cat "$GENESIS_PATH" | jq "$1" > "$TMP_GENESIS_PATH" && mv "$TMP_GENESIS_PATH" "$GENESIS_PATH"
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

    # grant the ica host execute permissions to all modules' messages
    update_test_genesis '.app_state["interchainaccounts"]["host_genesis_state"]["params"]["allow_messages"]=["*"]'

    # sl suite already creates genesis accounts, a gentx, collects the gentx, and validates the gensis file.

}

cleanup() {
	echo "SIGINT captured, starting cleanup"
    # need to clean up resources or re-arrange directories?
	exit
}

from_scratch

# Opens the RPC endpoint to outside connections
if [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' '/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:26657"' /var/cosmos-chain/canined-2/config/config.toml
  sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /var/cosmos-chain/canined-2/config/config.toml
else
  sed -i '/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:26657"' /var/cosmos-chain/canined-2/config/config.toml
  sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /var/cosmos-chain/canined-2/config/config.toml
fi

trap "cleanup" SIGINT

# I believe sl suite already starts canined 

# clean after program termination without SIGINT
cleanup