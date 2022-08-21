<!--
order: 0
title: Canine Storage Provider
parent:
  title: "cmd"
-->
[◀ readme](/README.md)

# Canine Storage Provider

## Overview
The storage provider is a web-server that accepts incoming files from users and creates contracts for the users to approve. These contracts last until the user either cancels them or the provider itself goes offline.

## Quickstart
This assumes you have either already set up a node or are using another RPC provider in your `~/.canined/config/client.toml` file.

To quickly set up a storage provider, one must initialize their provider & announce themselves to the network. Then they start the provider from their account of choice which stores files in the `~/.canined/config/networkfiles` folder (this can be changed with the --home flag).

To use a different home directory to store the files, you must run `init` with the home flag set to where you wish to initialize it.

You must also be using the keyring-backend: `test`. (`canined config keyring-backend test`)

```sh
$ canined tx storage init-miner {IP_ADDRESS} 100000000000 --from {KEY_NAME} --gas-prices=0.002ujkl --gas-adjustment=1.5

$ canined start-miner --from {KEY_NAME} --gas-prices=0.002ujkl --gas-adjustment=1.5 -y
```

## Posting files
Files can be uploaded through a POST request to `localhost:3333/upload` with form data.
### Form Data
|Key   |Data      |
|------|----------|
|file  |{filedata}|
|sender|{address} |

### Response
The response will be a JSON response formatted as:
```JSON
{
    "CID": "cid...",
    "FID": "fid..."
}
```

## Getting files
Gettings files is as easy as running a GET request at `localhost:3333/download/{FID}`. This will return the file as a blob to the browser.

