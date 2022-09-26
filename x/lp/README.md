<!--
order: 0
title: Jackal Storage Overview
parent:
  title: "storage"
-->
[◀ modules](/x/README.md)

# `lp`

## Contents

1. [Jackal Automated Market Maker (AMM)](#jackal-automated-market-maker-(amm))
2. [Client](#client)
    + [Query](#query)
        + [list-l-pool](#list-l-pool)
        + [show-l-pool](#show-l-pool)
        + [estimate-deposit](#estimate-deposit)
        + [estimate-return](#estimate-return)
        + [params](#params)
    + [Transactions](#transactions)
        + [create-l-pool](#create-l-pool)
        + [deposit-l-pool](#deposit-l-pool)
        + [withdraw-l-pool](#withdraw-l-pool)
        + [swap](#swap)
3. [Development](#development)
    + [Todo](#todo)
    + [Coins](#coins)
    + [States](#states)
         + [LPool](#lpool-(liquidity-pool))
         + [LProviderRecord](#lproviderrecord)

## Jackal Automated Market Maker (AMM)

Jackal AMM provides allows users to create a liquidity pool and swap tokens.
Aimed to provide wide variety of tokens as a payment for Jackal storage.

Swap fees are taken from deposit and deposited directly back into the liquidity pool.
Liquidity providers can burn their liquidity pool tokens to collect their reward
in the liquidity pool tokens.

## Client

### Query

The `query` (`q` for shortcut) commands allow users to query state of liquidity pools.

```sh
canined q storage --help
```

#### list-l-pool

The `list-l-pool` command allows users to see a list of currently active
liquidity pools.  

```sh
canined q storage list-contracts
```

#### show-l-pool

The `show-l-pool` command allows users to view information about a specific
liquidity pool.

```sh
canined q lp show-l-pool [index]
```

#### estimate-deposit

The `estimate-deposit` command allows users to estimate deposit amount for
desired amount of tokens in a swap.  
Expected format for `[desired-coins]` (omit curly brackets): "{denom0}{amount0}
...{denomN}{amountN}"

```sh
canined q estimate-deposit "pool-name" [desired-coins] 
```

#### estimate-return

The `estimate-return` command allows users to estimate return amount when
depositing x amount in a swap.  
Expected format for `[depositing-coins]` (omit curly brackets): "{denom0}{amount0}
...{denomN}{amountN}"

```sh
canined q estimate-return "pool-name" [depositing-coins] 
```

#### params

The `params` command allows users to view the params of the module.

```sh
canined q lp params
```

### Transactions

The `tx` commands allow users to interact with the `storage` module.

```sh
canined tx storage --help
```

#### create-l-pool

The `create-l-pool` command initializes a liquidity pool.  
Expected format for `[pool-coins]` (omit curly brackets): "{denom0}{amount0}
...{denomN}{amountN}"

```sh
canined tx lp create-l-pool [pool-coins] [invariant model id] \
   "swap-fee-percentage" [pool-lock-time] "withdraw-penalty-percentage"
```

#### deposit-l-pool

The `deposit-l-pool` command allows users to contribute to a pool and receive
liquidity pool token.
Expected format for `[deposit-coins]` (omit curly brackets): "{denom0}{amount0}
...{denomN}{amountN}"

```sh
canined tx lp deposit-l-pool "pool-name" [deposit-coins]
```

#### withdraw-l-pool

The `withdraw-l-pool` command allows users to burn their liquidity pool token
to receive pool coins.  
The `[burn-amount]` is an integer.

```sh
canined tx withdraw-l-pool "pool-name" [burn-amount]
```

#### swap

The `swap` command allows users to swap token in a liquidity pool.  
Expected format for `[deposit-coins]` (omit curly brackets): "{denom0}{amount0}
...{denomN}{amountN}"

```sh
canined tx withdraw-l-pool "pool-name" [deposit-coins]
```

# Development

## Todo

- [ ] Implement withdraw penalty mechanism (in progress).
- [ ] Change LPool.poolLockTime `uint64` type to `time`.
- [ ] Implement service fee mechanism.
- [ ] Add more query to aid liquidity pool creation.
- [x] Add estimate query regarding liquidity pool tokens.

## States

[LPool](#lpool-(liquidity-pool)) and [LProviderRecord](#lproviderrecord) states
are managed by `lp` module.

### LPool (liquidity pool)

It keeps the record of coins that are in the pool and amount of liquidity pool
token that exists.

#### `LPool`

```proto
message LPool {
  string index = 1; 
  // Liquidity pool coins that users can swap.
  string coins = 2;
  // AMM model used to balance coins. Valid model id must be used to initialize
  // the pool.
  uint32 invariantModelId = 3;
  // Swap fee that is taken away from deposit. Swap return is calculated after
  // swap fee is subtracted.
  float swapFeePerc = 4;
  // Denoms of pool coins in format: {denom0}-{denom1}...-{denomN}
  string poolDenoms = 5;
  // Duration where burning LP token is locked. Panelty is applied when
  // a user decides to burn the token during lock time.
  // Note: This type will change to `google.protobuf.duration` in later updates.
  uint64 poolLockTime = 6;
  // Portion of coins taken away from pool coins return.
  float withdrawPaneltyPerc = 7;
  // Liquidity pool name.
  string name = 8;
  // Denom of this LP token.
  string lPTokenDenom = 9;
  // Total amount of LP token that exists.
  string totalLPToken = 10;
}
```

### LProviderRecord

This keeps the record of liquidity provider of a liquidity pool. When a user
contributes to multiple pools the record is created for every pool.
This allows users to keep track of their contribution, collected reward and
LP token unlock time.

#### `LProviderRecord`

This is created only once when a user contributes to a pool and only updated on
succeeding burn or deposit transactions. This is deleted when the user burns
all of their LP token.

```proto
message LProviderRecord {
	string provider = 1;
	string poolName = 2;
	string coinsProvided = 3;
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

