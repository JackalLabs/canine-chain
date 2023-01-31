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
        + [list-active-deals](#list-active-deals)
        + [show-active-deals](#show-active-deals)
        + [list-miners](#list-miners)
        + [show-miners](#show-miners)
        + [list-contracts](#list-contracts)
        + [show-contracts](#show-contracts)
        + [find-file](#find-file)
        + [freespace](#freespace)
        + [get-client-free-space](#get-client-free-space)
        + [list-pay-blocks](#list-pay-blocks)
        + [show-pay-blocks](#show-pay-blocks)
        + [list-proofs](#list-proofs)
        + [show-proofs](#show-proofs)
        + [list-strays](#list-strays)
        + [show-strays](#show-strays)
        + [params](#params)
    + [Transactions](#transactions)
        + [init-miner](#init-miner)
        + [set-miner-ip](#set-miner-ip)
        + [set-miner-totalspace](#set-miner-totalspace)
        + [sign-contract](#sign-contract)
        + [post-contract](#post-contract)
        + [postproof](#postproof)
        + [buy-storage](#buy-storage)
        + [cancel-contract](#cancel-contract)


## Jackal Proof-of-Persistence (JPOP)
Jackal Storage functions by a Proof-of-Storage algorithm we call Proof-of-Persistence. The Jackal Proof-of-Persistence (JPOP) works through a series of contracts formed between the storage provider and the user. These contracts contain the Merkle Tree root hash of the file and the information required to prove ownership of the file. Miners, or as well call them, Storage Providers, are responsible for posting Merkle Proofs within a challenge window determined by the blockchain. These challenge windows require the miner to post the raw data chunk of data corresponding to the index of the challenge window alongside the required Merkle Hashes to prove the data belongs to the Merkle Root stored on the contract. These challenge indexes are chosen at random by the blockchain using a block-hash-based random number generator paired with a random data oracle. 

### Internal Detection Of Loss (IDOL) Protocol
If a Storage Provider successfully posts a Merkle Proof within the challenge window for the contract and the data is verified by the Validators to be valid Merkle Proofs for the challenge index, the Storage Provider is paid out. Storage Provider rewards are proportional to the file size the contract is associated with relative to every other active contract on the network. If a Storage Provider fails to provide a valid proof within the allotted timeframe, the contract is struck with a missed proof. After X missed proofs the contract is burned and the User is alerted the next time they query the contract. For every contract burned through missing proofs, the Storage Provider is struck with a penalty that remains on their record for a period of time adjustable through governance. These contracts are then moved to a new list where they are able to be claimed by other providers. This system follows our secondary protocol IDOL (Internal Detection Of Loss), where a contract is claimed by a new provider and downloaded from one of the two online providers storing the same file, thus resuming that contract's proof action, returning redundancy to 3x.

## Interaction Outline
A user first sends a file to an available Storage Provider. A list of Storage Providers can be found on the blockchain and miners can deny any incoming request if they wish not to store new files. The Storage Provider, after receiving the entire file, keeps that file in memory and posts a contract to the blockchain. If the contract is not signed by the sender in X blocks (configurable by the Storage Provider), then the file is removed from memory and the contract is burned. However, if the contract is signed by the user within the given blocks, the file is committed to the Storage Provider's hard storage and the challenge windows start being created for the now active contract.

## Client
### Query
The `query` commands allow users to query `storage` state.
```sh
canined q storage --help
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
#### find-file
The `find-file` command lets a user enter a file id (fid) and return every storage provider currently storing that file.
```sh
canined q storage find-file [fid]
```
#### freespace
The `freespace` command returns the free space in bytes of the miner.
```sh
canined q storage freespace [miner-address]
```
#### get-client-free-space
The `get-client-free-space` command takes an address and returns the total data they have available to use in bytes.
```sh
canined q storage get-client-free-space [address]
```
#### list-pay-blocks
The `list-pay-blocks` command lists all of the payblocks created by users.
```sh
canined q storage list-pay-blocks
```
### show-pay-blocks
The `show-pay-blocks` command shows information about a single pay block.
```sh
canined q storage show-pay-blocks [blockid]
```
#### list-strays
The `list-strays` command lists all of stray contracts.
```sh
canined q storage list-strays
```
### show-strays
The `show-strays` command shows information about a single stray contract from a given contract id (cid).
```sh
canined q storage show-strays [cid]
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
### postproof
The `postproof` command allows storage providers to post a proof claiming they have the file available.
```sh
canined tx storage postproof [chunk_data] [proof_data]
```
---
### buy-storage
The `buy-storage` command allows users to pay for a specific amount of storage for a specified period of time.
```sh
canined tx storage buy-storage [for-address] [duration] [bytes] [payment-denom]
```
Example:
```sh
canined tx storage buy-storage jkl1t3stAcc0unt 720h 6000000000 ujkl
```

#### Failed Cases:
 - buy storage while having an active plan
 - buy less than the current usage (SpaceUsed)
 - buy less than a GB
 - buy less than a month
 - pay with anything other than ujkl
 --- 

### cancel-contract
The `cancel-contract` command allows users to cancel currently active contracts removing the data usage from their account.
```sh
canined tx storage cancel-contract [cid]
```
