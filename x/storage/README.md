<!--
order: 0
title: Jackal Storage Overview
parent:
  title: "storage"
-->
[◀ modules](/x/README.md)

# `storage`

## Contents
1. [Jackal Proof-of-Persistence (JPOP)](#jackal-proof-of-persistence-jpop)
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
        + [init-miner](#init-miner)
        + [set-miner-ip](#set-miner-ip)
        + [set-miner-totalspace](#set-miner-totalspace)
        + [sign-contract](#sign-contract)
        + [post-contract](#post-contract)
        + [postproof](#postproof)


## Jackal Proof-of-Persistence (JPOP)
Jackal Storage functions by a Proof-of-Storage algorithm we call Proof-of-Persistence. The Jackal Proof-of-Persistence (JPOP) works through a series of contracts formed between the storage provider and the user. These contracts contain the Merkle Tree root hash of the file and the information required to prove ownership of the file. Storage Providers, or as well call them, Miners, are responsible for posting Merkle Proofs within a challenge window determined by the blockchain. These challenge windows require the miner to post the raw data chunk of data corresponding to the index of the challenge window alongside the required Merkle Hashes to prove the data belongs to the Merkle Root stored on the contract. These challenge indexes are chosen at random by the blockchain using a block-hash-based random number generator paired with a random data oracle. 

If a miner successfully posts a Merkle Proof within the challenge window for the contract and the data is verified by the Validators to be valid Merkle Proofs for the challenge index, the miner is paid out. Miner rewards are proportional to the file size the contract is associated with relative to every other active contract on the network. If a Miner fails to provide a valid proof within the allotted timeframe, the contract is struck with a missed proof. After X missed proofs the contract is burned and the User is alerted the next time they query the contract. For every contract burned through missing proofs, the Miner is struck with a penalty that remains on the Miner's record for a long period of time.

## Interaction Outline
A user first sends a file to an available Miner. A list of miners can be found on the blockchain and miners can deny any incoming request if they wish not to store new files. The Miner, after receiving the entire file, keeps that file in memory and posts a contract to the blockchain. If the contract is not signed by the sender in X blocks (configurable by the miner), then the file is removed from memory and the contract is burned. However if the contract is signed by the user within the given blocks, the file is committed to the Miner's hard storage and the challenge windows start being created for the now active contract.

## Client
### Query
The `query` commands allow users to query `storage` state.
```sh
canined q storage --help
```
#### list-contracts
The `list-contracts` command allows users to see a list of currently alive contracts that are awaiting signatures.
```sh
canined q storage list-contracts
``` 
#### show-contracts
The `show-contracts` command allows users to view information about a specific contract as the identifier.
```sh
canined q storage show-contracts [cid]
```
#### list-miners
The `list-miners` command allows users to view a list of currently registered mining nodes.
```sh
canined q storage list-miners
```
#### show-miners
The `show-miners` command allows users to view information about a specific miner by passing in its jackal address.
```sh
canined q storage show-miners [address]
```
#### list-active-deals
The `list-active-deals` command allows users to view a list of currently active contracts that have already been signed.
```sh
canined q storage list-active-deals
```
#### show-active-deals
The `show-active-deals` command allows users to view information about a specific active deal by passing in the contract id.
```sh
canined q storage show-active-deals [cid]
```
#### freespace
The `freespace` command returns the free space in bytes of the miner.
```sh
canined q storage freespace [miner-address]
```
#### params
The `params` command allows users to view the params of the module.
```sh
canined q storage params
```

### Transactions
The `tx` commands allow users to interact with the `storage` module.
```sh
canined tx storage --help
```
#### init-miner
The `init-miner` command posts a miner's interface info alongside the total storage offered by the miner. 
```sh
canined tx storage init-miner [ip-address] [total-size]
```
#### set-miner-ip
The `set-miner-ip` command updates a miner's IP address.
```sh
canined tx storage set-miner-ip [ip]
```
#### set-miner-totalspace
The `set-miner-totalspace` command updates a miner's total space available.
```sh
canined tx storage set-miner-totalspace [total-space]
```
#### sign-contract
The `sign-contract` command signs a contract by passing in the contract id.
```sh
canined tx storage sign-contract [cid]
```
#### post-contract
The `post-contract` command posts a contract to the blockchain from a miner.
```sh
canined tx storage post-contract [hashes] [signee] [duration] [filesize] [file-id]
```