#!/bin/bash

# Infinite loop to run the command every 10 seconds
while true
do
    # Run your specific command
    canined tx storage post-random 1000000 --from deposit_account --gas-prices=0.02ujkl -y

    # Wait for 10 seconds before the next run
    sleep 10
done
