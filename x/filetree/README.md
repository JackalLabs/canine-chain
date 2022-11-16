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
2. [Client](#client)
    + [Query](#query)
    + [Transactions](#transactions)


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
