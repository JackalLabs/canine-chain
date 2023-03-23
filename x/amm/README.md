<!--
order: 0
title: Jackal Storage Overview
parent:
  title: "storage"
-->
[◀ modules](/x/README.md)

# `amm`

## Contents

1. [Jackal Automated Market Maker (AMM)](#jackal-automated-market-maker-(amm))
2. [Client](#client)
    + [Query](#query)
        + [list-pool](#list-pool)
        + [show-pool](#show-pool)
        + [estimate-swap-in](#estimate-swap-in)
        + [estimate-swap-out](#estimate-swap-out)
        + [list-records-from-pool](#list-records-from-pool)
        + [show-liquidity-provider-record](#show-liquidity-provider-record)
        + [params](#params)
    + [Transactions](#transactions)
        + [create-pool](#create-l-pool)
        + [join-pool](#join-pool)
        + [exit-pool](#exit-pool)
        + [swap](#swap)
3. [Development](#development)
    + [Todo](#todo)
    + [Coins](#coins)
    + [States](#states)
         + [Pool](#lpool-(liquidity-pool))
         + [LiquidityProviderRecord](#lproviderrecord)

## Jackal Automated Market Maker (AMM)

Jackal AMM provides allows users to create a liquidity pool and swap tokens.
Aimed to provide wide variety of tokens as a payment for Jackal storage.

Swap fees are taken from deposit and deposited directly back into the liquidity pool.
Liquidity providers can burn their liquidity pool tokens to collect their reward
in the liquidity pool tokens.

## Client

Expected format for [`coin`](https://buf.build/cosmos/cosmos-sdk/docs/main:cosmos.base.v1beta1#cosmos.base.v1beta1.Coin):
{amount}{denom} e.g. 100000ujkl  
List of coins: {amount0}{denom0},...,{amountN}{denomN}  
e.g. 100000ujkl, 200000stake  
*amount has to be an integer value

### Query

The `query` (`q` for shortcut) commands allow users to query state of liquidity pools.

```sh
canined q storage --help
```

#### list-pool

The `list-pool` command allows users to see a list of currently active
liquidity pools.  

```sh
canined q storage list-contracts
```

#### show-pool

The `show-pool` command allows users to view information about a specific
liquidity pool.

```sh
canined q lp show-pool [index]
```

#### estimate-swap-in

The `estimate-swap-in` command allows users to estimate swap in coin for
desired swap out coin.  

```sh
canined q estimate-swap-in "pool-name" [desired-coin-swap-out] 
```

#### estimate-swap-out

The `estimate-swap-out` command allows users to estimate swap out coin for swap in coin.  

```sh
canined q estimate-swap-out "pool-name" [swap-in-coin] 
```

#### list-pool-providers

The `list-pool-providers` command allows users to see all liquidity providers
of a liquidity pool.  

```sh
canined q list-pool-providers "pool-name" 
```

#### show-liquidity-provider-record

The `show-liquidity-provider-record` command allows users to see a liquidity provider
record.  

```sh
canined q show-l-provider-record "pool-name" "provider-address"
```

#### params

The `params` command allows users to view the params of the module.

```sh
canined q amm params
```

### Transactions

The `tx` commands allow users to interact with the `amm` module.

```sh
canined tx amm --help
```

#### create-pool

The `create-pool` command initializes a liquidity pool.  

```sh
canined tx amm create-pool [pool-coins] [AMM_Id] \
   "swap-fee-multiplier" [pool-lock-time (int64)] "withdraw-penalty-multiplier"
```

#### join-pool

The `join-pool` command allows users to provide liquidity to a pool and receive
liquidity pool token.

```sh
canined tx amm join-pool "pool-name" [coins]
```

#### exit-pool

The `exit-pool` command allows users to burn their liquidity pool token
to receive pool coins.  
*The `[burn-amount]` is an integer.

```sh
canined tx amm exit-pool "pool-name" [burn-amount]
```

#### swap

The `swap` command allows users to swap token in a liquidity pool.  
Use `[minimum-swap-out]` to prevent swap if swap output is below that amount.  

```sh
canined tx amm swap "pool-name" [swap-in] [minimum-swap-out]
```

# Development

## States

[Pool](#pool-(liquidity-pool)) and
[LiquidityProviderRecord](#liquidityproviderrecord) states
are managed by `amm` module.

### Pool (liquidity pool)

Pool keeps the record of coins that are in the pool and amount of liquidity pool
token that exists.

#### `Pool`

```proto
message Pool {
  string index = 1; 
  // Pool name
  string name = 2;
  // Pool coins
  repeated cosmos.base.v1beta1.Coin coins = 3 [(gogoproto.nullable) = false];
  // AMM model used to balance the pool
  uint32 aMM_Id = 4;
  // sdk.Dec in string format
  // This is deducted from swap input before swap output is calculated
  // swap_input = swap_input - (swap_input * swap_fee_multi)
  string swap_fee_multi = 5;
  // Duration of LPToken being locked from burning
  // Penalty is applied when LPToken is burned before lock expires
  // Duration is in seconds
  int64 min_lock_duration = 6;
  // sdk.Dec in string format
  // Penalty applied to LPToken burn
  // Full requested amount is burned but penalty is applied before
  // LP tokens returned is calculated
  string penalty_multi = 7;
  string lpToken_denom = 8;
  // Amount of total LPToken that exists
  string LPTokenBalance = 9;
}
```

### LiquidityProviderRecord

This keeps the record of liquidity provider of a liquidity pool. When a user
contributes to multiple pools the record is created for every pool.
This allows users to keep track of their contribution, collected reward and
LP token unlock time.

#### `LiquidityProviderRecord`

This is created only once when a user contributes to a pool and only updated on
succeeding burn or deposit transactions. This is deleted when the user burns
all of their LP token.

```proto
message LiquidityProviderRecord {
	string provider = 1;
	string poolName = 2;
	repeated cosmos.base.v1beta1.Coin coins coinsProvided = 3;
	google.protobuf.Timestamp unlockTime = 4;
	google.protobuf.Duration lockDuration = 5;
}
```

This is stored at KVStore with
`{LProviderRecordKeyPrefix}{poolName}{provider}` key.
It is advised to use [`Prefix` store](https://docs.cosmos.network/master/core/store.html#prefix-store)
to skip prepending `{LProviderRecordKeyPrefix}` key.

#### LProviderRecord References

Additional reference to `LProviderRecord` key is stored at KVStore for
efficient query.  
The reference key is in format of
`{LProviderRecordKeyPrefix}{provider}{poolName}`.
This can be used effectively with [`KVStorePrefixIterator()`](https://github.com/cosmos/cosmos-sdk/blob/v0.46.1/types/store.go#L30).  
This stored in KVStore will look like `{ReferenceKey} = {LProviderRecordKey}`.

For example, to query all `LProviderRecord` of account `123`, put the
address to reference `{provider}` and use `KVStorePrefixIterator()` to iterate
through all `LProviderRecord` keys that has `123` as `{provider}`.


## Notes

### :warning: Coins

All interaction regarding coins must be normalized before further interaction.
Normalizing converts coins to their smallest unit.
This eliminates confusion and headaches dealing with coin units and decimals.  
Useful functions are [`ParseCoinsNormalized()`](https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/types/coin.go#L919)
and [`NormalizeCoins()`](https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/types/denom.go#L135)

