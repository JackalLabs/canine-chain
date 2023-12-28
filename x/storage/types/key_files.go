package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	FileSecondaryKeyPrefix = "FilesByOwner/value/"
	FilePrimaryKeyPrefix   = "FilesByMerkle/value/"
	ProofKeyPrefix         = "FileProof/value/"

	LegacyActiveDealsKeyPrefix = "ActiveDeals/value/" // OLD! DO NOT USE!

)

// FilesPrimaryKey returns the store key to retrieve a File from the index fields ordered by merkle
func FilesPrimaryKey(
	merkle []byte,
	owner string,
	start int64,
) []byte {
	return []byte(fmt.Sprintf("%x/%s/%d/", merkle, owner, start))
}

// FilesMerklePrefix returns the prefix for a merkle
func FilesMerklePrefix(
	merkle []byte,
) []byte {
	return []byte(fmt.Sprintf("%s%x", FilePrimaryKeyPrefix, merkle))
}

// FilesOwnerPrefix returns the prefix for a owner
func FilesOwnerPrefix(
	owner string,
) []byte {
	return []byte(fmt.Sprintf("%s%s", FileSecondaryKeyPrefix, owner))
}

// FilesSecondaryKey returns the store key to retrieve a File from the index fields ordered by owner
func FilesSecondaryKey(
	merkle []byte,
	owner string,
	start int64,
) []byte {
	return []byte(fmt.Sprintf("%s/%x/%d/", owner, merkle, start))
}

// ProofKey returns the store key to retrieve a proof from the index fields ordered by owner
func ProofKey(
	prover string,
	merkle []byte,
	owner string,
	start int64,
) []byte {
	return []byte(fmt.Sprintf("%s/%s/%x/%d/", prover, owner, merkle, start))
}

// ProofKey returns the store key to retrieve a proof from the index fields ordered by owner
func ProofPrefix(
	prover string,
) []byte {
	return []byte(fmt.Sprintf("%s/%s", ProofKeyPrefix, prover))
}

// LegacyActiveDealsKey returns the store key to retrieve a ActiveDeals from the index fields
//
// Deprecated: UnifiedFile replaced Active Deals as the correct data structure
func LegacyActiveDealsKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
