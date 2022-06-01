package types

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
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
