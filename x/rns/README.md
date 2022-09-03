<!--
order: 0
title: RNS Overview
parent:
  title: "rns"
-->
[◀ modules](/x/README.md)

# `rns`

## Contents
1. [Concept](#concept)
2. [Client](#client)
    + [Query](#query)
        + [list-names](#list-names)
        + [show-names](#show-names)
        + [list-bids](#list-bids)
        + [show-bids](#show-bids)
        + [list-forsale](#list-forsale)
        + [show-forsale](#show-forsale)
        + [params](#params)
    + [Transactions](#transactions)
        + [register](#register)
        + [add-record](#add-record)
        + [del-record](#del-record)
        + [list](#list)
        + [delist](#delist)
        + [buy](#buy)
        + [bid](#bid)
        + [cancel-bid](#cancel-bid)
        + [accpet-bid](#accept-bid)


## Concept
The `rns` module is a nameservice that allows users to use a human readable name when interacting with the Jackal Blockchain. Users can register names, list names for sale, buy names on sale, and place/accept bids from other users on their names.

## Client
### Query
The `query` commands allow users to query `rns` state.
```sh
canined q rns --help
```
#### list-names
The `list-names` command allows users to see a list of registered names.
```sh
canined q rns list-names
``` 
#### show-name
The `show-names` command allows users to see information about a single name.
```sh
canined q rns show-names
``` 
#### list-bids
The `list-bids` command allows users to see a list of currently active bids.
```sh
canined q rns list-bids
```
#### show-bids
The `show-bids` command allows users to see information about an active bid.
```sh
canined q rns show-bids
```
#### list-forsale
The `list-forsale` command allows users to see a list of names on sale.
```sh
canined q rns list-forsale
```
#### show-forsale
The `show-forsale` command allows users to see information about a name on sale.
```sh
canined q rns show-forsale
```
#### params
The `params` command allows users to view the params of the module.
```sh
canined q rns params
```

### Transactions
The `tx` commands allow users to interact with the `rns` module.
```sh
canined tx rns --help
```
#### register
The `register` command registers a name for the years specified under the user account. Must also pass in a data field, used for storing a JSON structure defining extra characteristics of the name. Will take tokens from user account proportional to years registered & length of name.
```sh
canined tx rns register [name] [years] [data]
```
#### add-record
The `add-record` command appends the new record to the name acting as a subdomain.
```sh
canined tx rns add-record [name] [record] [data]
```
#### del-record
The `del-record` command removes the record from the name.
```sh
canined tx rns del-record [name] [record] [data]
```
#### list
The `list` command lists a name for sale at a specified price. When bought, the amount specified will be sent to the user.
```sh
canined tx rns list [name] [price]
```
#### delist
The `delist` command removes a name from the sale listings.
```sh
canined tx rns delist [name]
```
#### buy
The `buy` command buys a name that is listed for sale.
```sh
canined tx rns buy [name]
```
#### bid
The `bid` command places a bid for a name. The price of the bid is locked up in escrow until either the bid is cancelled or is accepted. The owner can accept this bid and the name will be transfered.
```sh
canined tx rns bid [name] [price]
```
#### cancel-bid
The `cancel-bid` command removes a bid from the bid list and returns the funds locked in escrow.
```sh
canined tx rns cancel-bid [bid-id]
```
#### accept-bid
The `accept-bid` command accepts a bid, transferring ownership of the name specified and taking the tokens locked in escrow.
```sh
canined tx rns accept-bid [bid-id]
```