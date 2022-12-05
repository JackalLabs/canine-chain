<!--
order: 0
title: filetree Overview
parent:
  title: "filetree"
-->
[◀ modules](/x/README.md)

# `filetree`

## Contents
1. [Concept](#concept)
2. [Merkle Paths](#merkle-paths)
3. [Client](#client)
    + [Query](#query)
    + [Transactions](#transactions)
4. [Transaction Messages](#transaction-messages)
    + [Postkey](#postkey)
    + [Makeroot](#makeroot)
    + [Postfile](#postfile)
    + [Delete](#delete)
5. [Query Requests](#query-requests)
    + [GetPubkey](#getpubkey)
    + [GetFiles](#getfiles)



    


## Concept
The `filetree` module is responsible for keeping records of a user's files and organizing them in a way that is accessible. When a user uploads a file using the Storage module, the file is only accessible from the File ID (`fid`) which makes the process clunky and obtuse to remember every file uploaded to Jackal. Furthermore, every single upload would be required to be public, or the user would need to keep track of every symmetric key used to encrypt the files and manually map them to the `fids`. The solution to this is a tree structure storing each file as an entry in the tree. Organizing this structure is also trivial as we can assign children to pseudo files that we call folders. Finally to keep track of encryption keys, the protocol maps every file to its respective key.

## Merkle Paths
To keep files hidden from the public eye, we fully hash the file paths for every file. Usually just hashing the entire path would result in two file paths having no relationship even if in the plaintext path, one of the files would be a child of the other.

Ex: `/home/path/` & `/home/path/child` are very clearly related.

Through our system we call Merkle Paths, we can keep these relations intact while still pushing for privacy through one-way hash functions.

```go
let path := 'home/path/'
let child := 'child'

let H := func() return hash(item)

// we split `path` by `/`
let h1 := 'home'
let h2 := 'path'

// we hash each value of split(path)
let hh1 := H(h1)
let hh2 := H(h2)

// we also hash the child
let hh3 := H(child)

// this leaves us with each chunk hashed and protected from prying eyes already, 
// but it's clear this has 3 items in its path (NOT GOOD)

// we can then hash every item in this list with the proceeding item

let s := H(hh1)
ps := H(s + hh2)
cs := H(ps + hh3)

// what we are left with is the merkle path of the child (cs) and the path of the parent (ps)

// at any time we can check if cs is a child of ps if we have H(child)

func checkChild(parent ps, child cs, hc H(child)) {
    return H(parent + hc) == child
}

// this gives us the ability to keep the usefulness of plaintext paths but add a 
// level of privacy that would be impossible to achieve with plaintext.
```


## Client

Below are CLI query and transaction commands to interact with canined.
### Query
The `query` commands allow users to query `filetree` state.
```sh
canined q filetree --help
```

### Transactions
The `tx` commands allow users to interact with the `filetree` module.
```sh
canined tx filetree --help
```

## Transaction Messages

Below is a full description of valid transaction messages that can be broadcasted to affect state change. These descriptions aim to be "implementation agnostic", i.e., they make sense to both the CLI/Golang and TS implementations. 

### Postkey

Post a ecies.PublicKey on chain for the encryption scheme

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address 
|key  | String  | ecies.PublicKey 

#### Response

Coming soon

### Makeroot 

Create an absolute root folder for a storage account. 

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in Bech32 address<br /> 
|account  | String  | Hex[ hash( Bech32 address of user that will own this account)]. <br /> Please note that the broadcaster of this message will always be making a storage account for themselves, but there are other filetree transaction messages that can be called by userA to affect a change in userB's account. It is for this purpose that the Account field exists.<br /> 
|rootHashPath  | String  | MerklePath("s")<br />
|contents  | String  | FID<br />
|editors  | String  | string(json encoded map) with: <br />let c = concatenate( "e", trackingNumber, Bech32 address )<br />map_key: hex[ hash("c") ]<br />map_value: ECIES.encrypt( aesIV + aesKey )<br />
|viewers  | String  | Pass in "NONE." Do not pass in an emptry string else message validation will fail. Root folder has no viewers. Unknown at this time if this field will be needed in the future so we leave it in for now. <br />
|trackingNumber  | String  | UUID. This trackingNumber is one and the same as what is used in editors map

#### Response

Coming soon

### PostFile

Create and save a new 'Files' struct on chain. The distinction between a folder and a file is very clear in jackalJS, but the filetree module does not care whether a Files struct is being used to save data for a folder or a file. 

Let it be that alice wants to create a home folder

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in alice's Bech32 address<br /> 
|account  | String  | Hex[ hash( alice's Bech32 address )]<br />
|hashParent  | String  | MerklePath("s")<br />
|hashChild  | String  |  Hex[ hash("home") ]<br />
|contents  | String  | FID<br />
|viewers  | String  | string(json encoded map) with: <br />let c = concatenate( "v", trackingNumber, Bech32 address )<br />map_key: hex[ hash("c") ]<br />map_value: ECIES.encrypt( aesIV + aesKey )<br />
|editors  | String  | same as above but with c = concatenate( "e", trackingNumber, Bech32 address )<br />
|trackingNumber  | String  | UUID. This trackingNumber is one and the same as what is used in editors AND viewers map<br />

alice can add other users to her viewers and editors map aswell. 
#### Response

let fullMerklePath = MerklePath("s/home")

```json
{
  
    "path": "fullMerklePath"

}
```
### Delete

Let it be that alice wants to delete her "s/home" folder

|Name|Type|Description|                                                                                       
|--|--|--|
|creator  | String  | The creator and broadcaster of this message. Pass in alice's Bech32 address<br />
|hashPath  | String  | MerklePath("s/home")<br />
|account  | String  | Hex[ hash( alice's Bech32 address )]<br />

#### Response

Coming soon

## Query Requests

Below is a full description of valid query requests that can be made to retrieve state information. These descriptions aim to be "implementation agnostic", i.e., they make sense to both the CLI/Golang and TS implementations.

### GetPubkey

Retrieve a user's ecies.PublicKey

|Name|Type|Description|                                                                                       
|--|--|--|
|address  | String  | user's Bech32 address<br />

#### Response

types.PubKey

### GetFiles

Retrieve a Files struct. Let it be that alice want's to retrieve "s/home/bunny.jpg"

|Name|Type|Description|                                                                                       
|--|--|--|
|address  | String  | MerklePath("s/home/bunny.jpg")    <br />
|ownerAddress  | String  | accountHash = Hex[hash(alice's bech32 address)] <br /> let c = concatenate("o", MerklePath("s/home/bunny.jpg"), accountHash) <br /> OwnerAddress = hex[hash(c)]<br />


#### Response

types.Files


