<!--
order: 0
title: Jackal Accounts Overview
parent:
  title: "jklaccounts"
-->
[◀ modules](/x/README.md)

# `jklaccounts`

## Contents
1. [Concept](#concept)
2. [Client](#client)

## Concept
The Jackal Account module handles the user accounts on the Jackal system. Not to be confused with [`jklpayments`](/x/jklpayments/README.md) which handles the payment structure of the Jackal eco-system. Each user account is attributed an `available` tag and a `used` tag. These tags determine how much storage space a user is authorized to use, as well as how much space the user is currently using. This lets the miners know whether or not to store the incoming files on behalf of users and prevents clogging/cheating of the system.

## Client
### Query
The `query` command allows users to query the `jklaccounts` state.
```sh
canined q jklaccounts --help
```
#### check-can-store
The `check-can-store` command takes a users address and a file size and checks to see if the user is allowed to store a file of that size.
```sh
canined q jklaccoutns check-can-store [user_address] [size_in_bytes]
```
#### system-storage
The `system-storage` command returns the total amount of space that the Jackal system is using. This is purely used for analytics and no information about which users make up this data is returned.
```sh
canined q jklaccounts system-storage
```

