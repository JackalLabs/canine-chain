package types

import "fmt"

var MinterKey = []byte{0x00}

const (
	// ModuleName defines the module name
	ModuleName = "jklmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_jklmint"

	// Query endpoints supported by the minting querier
	QueryParameters   = "parameters"
	QueryInflationKey = "inflation"

	LastBlockMinted = "last_block_minted"

	DevGrantsPool = "dev_grant_pool"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func MintedBlockKey(height int64) []byte {
	return []byte(fmt.Sprintf("minted_at_%d", height))
}
