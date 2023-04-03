#!/bin/bash
set -x

# building the validator docker image
docker buildx build . -t validator \
-f ValidatorDockerfile \
--build-arg RELEASE=v1.2.1 \
--build-arg VALIDATOR_SCRIPT=genesis-validator.sh \
--build-arg CHAIN_ID=testing \
--build-arg VALIDATOR_NAME=genesis-val \

# running the docker image
sudo docker run -d -v "/root/common_store:/home/common_store" \
--name=genesis-validator \
validator

# executing internal scripts
sudo docker exec genesis-validator bash /home/canine-validator/scripts/genesis-validator.sh \
VALIDATOR_NAME=genesis-val \
CHAIN_ID=testing

# spawning additional docker images, w/o the genesis validator
for i in {1..3}
do
    # running the docker image
    docker run -d -v "/root/common_store:/home/common_store" \
    --name=validator_${i} \
    validator

    # executing internal scripts
    docker exec validator_${i} bash /home/canine-validator/scripts/validator.sh \
    VALIDATOR_NAME=val_${i} \
    CHAIN_ID=testing \

done

# collecting the genesis_validator gentx
# docker exec genesis-validator rm -rf /root/.canine/config/gentx
docker exec genesis-validator cp /home/common_store/genesis.json /root/.canine/config/genesis.json
docker exec genesis-validator cp -RT /home/common_store/gentx /root/.canine/config/gentx
docker exec genesis-validator canined collect-gentxs --chain-id=testing
docker exec genesis-validator canined tendermint unsafe-reset-all --chain-id=testing 

# changing the genesis-validator listening address in config.toml
export IPADDR=$(docker exec genesis-validator hostname -I | tr -d ' ')
docker exec genesis-validator sed -i "s|tcp://0\.0\.0\.0:26656|tcp://$IPADDR:26656|" $HOME/.canine/config/config.toml
docker exec genesis-validator sed -i "s|tcp://127\.0\.0\.1:26657|tcp://$IPADDR:26657|" $HOME/.canine/config/config.toml

for i in {1..3}
    # collecting the gentxs
    # docker exec validator_${i} rm -rf /root/.canine/config/gentx
    docker exec validator_${i} cp /home/common_store/genesis.json /root/.canine/config/genesis.json
    docker exec validator_${i} cp -RT /home/common_store/gentx /root/.canine/config/gentx
    docker exec validator_${i} canined collect-gentxs --chain-id=testing
    docker exec validator_${i} canined tendermint unsafe-reset-all --chain-id=testing 

    # changing the genesis-validator listening address in config.toml
    export IPADDR=$(docker exec validator_${i} hostname -I | tr -d ' ')
    docker exec validator_${i} sed -i "s|tcp://0\.0\.0\.0:26656|tcp://$IPADDR:26656|" $HOME/.canine/config/config.toml
    docker exec validator_${i} sed -i "s|tcp://127\.0\.0\.1:26657|tcp://$IPADDR:26657|" $HOME/.canine/config/config.toml
done

# starting the chain 
docker exec -d genesis-validator canined start

for i in {1..3}
do
    docker exec -d validator_${i} canined start 
done

# # creating a new validator
# docker exec genesis-validator canined tx staking create-validator \
# --amount=400000000ujkl \
# --pubkey=$(docker exec genesis-validator canined tendermint show-validator --chain-id=testing) \
# --moniker=genesis-val \
# --chain-id=testing \
# --commission-rate="0.10" \
# --commission-max-rate="0.20" \
# --commission-max-change-rate="0.05" \
# --min-self-delegation="1" \
# --from=genesis-val
# --keyring-backend=test \
# --yes 

# for i in {1..3}
# do
#     # creating a new validator
#     docker exec validator_${i} canined tx staking create-validator --amount=400000000ujkl \
#     --from=val_${i} \
#     --pubkey=$(docker exec validator_${i} canined tendermint show-validator --chain-id=testing) \
#     --moniker=val_${i} \
#     --chain-id=$CHAIN_ID \
#     --commission-rate="0.1" \
#     --commission-max-rate="0.2" \
#     --commission-max-change-rate="0.05" \
#     --min-self-delegation="1" \
#     --keyring-backend=test \
#     --yes 
# done
