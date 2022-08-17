<!--
order: 0
title: dsig Overview
parent:
  title: "dsig"
-->
[◀ modules](/x/README.md)

# `dsig`

## Contents
1. [Concept](#concept)
2. [Client](#client)
    + [Query](#query)
        + [list-user-uploads](#list-user-uploads)
        + [show-user-uploads](#show-user-uploads)
        + [list-form](#list-form)
        + [show-form](#show-form)
    + [Transactions](#transactions)
        + [uploadfile](#uploadfile)
        + [createform](#createform)
        + [signform](#signform)


## Concept
The `dsig` module is a digital signature service that allows users to collect signatures from multiple users who are registered on the Jackal Blockchain. Users can create 'forms' associated with a unique file stored on Jackal and can add signees (users) to collect their signatures. The signees have the following options to respond: Approve, Deny, Abstain, and No Response (Default). The form can execute a custom function after all users have voted to Approve the form. 

Use cases for the `dsig` module include:

- Signing Contracts
- Approving a proposal
- Sending invitations
- General Voting

## Client
### Query
The `query` commands allow users to query `dsig` state.
```sh
canined q dsig --help
```
### list-user-uploads
The `list-user-uploads` command allow users to access a database of files compatible for generating Forms.
```sh
canined q dsig list-user-uploads [flags]
``` 

### show-user-uploads
The `show-user-uploads` command finds a specific file uploaded from its unique file id (FID)
```sh
canined q dsig show-user-uploads [fid] [flags]
```

### list-form
The `list-form` command allows users to access all Forms stored on the blockchain. 
```sh
canined q dsig list-form [flags]
``` 

### show-form
The `show-form` command finds a specific form from its unique Form ID (FFID)
```sh
canined qdsig show-form [ffid] [flags]
```

### Transactions
The `tx` commands allow users to interact with the `dsig` module.
```sh
canined tx dsig --help
```
### uploadfile
The `uploadfile` command stores the metadata relevant to the file on the blockchain.
```sh
canined tx dsig uploadfile [fid] [flags]
```
### createform
The `createform` command creates a form for digital signature collection. It accepts a unique File ID (fid) and a space-delimited string of Canine-Compatible user addresses (signees).

```sh
canined tx dsig createform [fid] [signees] [flags]
```
Example of a valid signees entry:
```golang
"canineaddr1 canineaddr2 canineaddr3"
```

### signform
The `signform` command applies the chosen vote for a valid signee. 
The vote is passed as a integer from [0-3].

Valid vote entries are:
```
Deny = 0 
Approve = 1
Abstain = 2 
No Response (Defualt) = 3
```
Command structure:
```sh
canined tx dsig signform [ffid] [vote]
```

