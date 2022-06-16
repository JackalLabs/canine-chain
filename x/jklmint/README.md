<!--
order: 0
title: Jackal Mint Overview
parent:
  title: "jklmint"
-->

# `jklmint`

## Contents
1. [Concept](#concept)
2. [Begin-Block](#begin-block)
3. [Client](#client)

## Concept
The Jackal Mint module is a drop-in replacement for the cosmos-sdk module: [Mint](https://github.com/cosmos/cosmos-sdk/blob/main/x/mint/spec/README.md). The key differences between this and the pre-existing minting module are that jklmint does not adjust inflation based on rate of bonded tokens. The jklmint module prints 10000000ujkl per block and distributes it to both the jklmining module and the default distribution module.

## Begin-Block
At the start of the block, the module creates 10000000ujkl (10JKL) and it is then sent to the fee collector. This is in contrast to the inflation model that the old mint module used.

## Client
There is no client interaction with the `jklmint` module.

