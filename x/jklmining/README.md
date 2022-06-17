<!--
order: 0
title: Jackal Mining Overview
parent:
  title: "jklmining"
-->

# `jklmint`

## Contents
1. [Concept](#concept)
2. [Interaction Outline](#interaction-outline)
2. [Client](#client)
    + [Query](#query)
        + [list-mined](#list-mined)
        + [show-mined](#show-mined)
        + [list-miners](#list-miners)
        + [show-miners](#show-miners)
        + [list-save-requests](#list-save-requests)
        + [show-save-requests](#show-save-requests)
        + [params](#params)
    + [Transactions](#transactions)
        + [allow-save](#allow-save)
        + [claim-save](#claim-save)
        + [create-miners](#create-miners)
        + [delete-miners](#delete-miners)
        + [update-miners](#update-miners)


## Concept
The Jackal Mining Module manages the Jackal Peer-to-Peer network. Miners report back to the chain with information about the file transactions to ensure that users are allowed to store data, as well as update the state of the storage network on-chain.

## Interaction Outline
A user first makes a request to the chain with a hashed key & the size of the file they wish to store. A user sends the unhashed key along with their file to the miner network, this allows the miner to check if they have posted the allowance to chain. If the user can store the file according to `x/jklpayments`, and the file is equal to the size allowance, then the miner stores the file and claims the allowance, setting the `Approved` field to `true`. 

## Client
### Query
The `query` commands allow users to query `jklmining` state.
```sh
canined q jklmining --help
```
#### list-mined
The `list-mined` command allows users to see a list of currently mined files, more specifically, if files are approved or not.
```sh
canined q jklmining list-mined
``` 
#### show-mined
The `show-mined` command allows users to view information about a specific file by passing in the hashed key as the identifier.
```sh
canined q jklmining show-mined [id]
```
#### list-miners
The `list-miners` command allows users to view a list of currently registered mining nodes.
```sh
canined q jklmining list-miners
```
#### show-miners
The `show-miners` command allows users to view information about a specific miner by passing in its jackal address.
```sh
canined q jklmining show-miners [address]
```
#### list-save-requests
The `list-save-requests` command allows users to view a list of currently claimed/unclaimed save-requests.
```sh
canined q jklmining list-save-requests
```
#### show-save-requests
The `show-save-requests` command allows users to view information about a specific save-request by passing in the hashed key.
```sh
canined q jklmining show-save-requests [index]
```
#### params
The `params` command allows users to view the params of the module.
```sh
canined q jklmining params
```

### Transactions
The `tx` commands allow users to interact with the `jklmining` module.
```sh
canined tx jklmining --help
```
#### allow-save
The `allow-save` command lets a user provision a saving of a file by specifying the size of the data as well as a hashed key that they can then give out to a miner to claim the saving of the data.
```sh
canined tx jklmining allow-save [passkey] [size]
```
#### claim-save
The `claim-save` command lets a miner claim a currently unclaimed save-request by passing in the unhashed key.
```sh
canined tx jklmining claim-save [key]
```
#### create-miners
The `create-miners` command lets a miner initialize their presence on the network by specifying their public facing IP address.
```sh
canined tx jklmining create-miners [ip]
```
#### delete-miners
The `delete-miners` command lets a miner delete themselves from the network.
```sh
canined tx jklmining delete-miners [address]
```
#### update-miners
The `update-miners` command lets a miner update their IP address in the event that it changes.
```sh
canined tx jklmining update-miners [address] [ip]
```