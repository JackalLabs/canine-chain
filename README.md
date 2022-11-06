![canine banner](banner.png)
# Canine Chain by Jackal
**Canine** is a blockchain built using Cosmos SDK and Tendermint.

## Pages

1. [Modules](x/README.md)
2. [Tokens](TOKENS.md)
3. [Storage Providers](/cmd/canined/README.md)


## Installing the Canine CLI
To install `canined` on your Linux machine:

`go install ./...`

### Pre-built Binary
[Releases](https://github.com/JACKAL-DAO/canine-chain/releases) and download the latest release. Move the executable to a folder in your `$PATH` and download [this](https://github.com/CosmWasm/wasmvm/raw/v1.1.1/internal/api/libwasmvm.x86_64.so) to `/lib/libwasmvm.x86_64.so` 

```sh
sudo wget https://github.com/CosmWasm/wasmvm/raw/v1.1.1/internal/api/libwasmvm.x86_64.so -O /lib/libwasmvm.x86_64.so
```

You may also need to run `sudo chmod +x canined` inside the executables directory to allow it to run.

## Testing this chain

`go test ./...`

## License

Canine by Jackal uses the [MIT License](/LICENSE.md).

## Bug Bounty

Refer to the bug bounty program proposed by Jackal Labs [Here](https://jackaldao.medium.com/announcement-jackal-bug-bounty-program-31d4e03ab7e2)

### [Developer Contact](/ABOUT.md)

