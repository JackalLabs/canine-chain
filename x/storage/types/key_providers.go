package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProvidersKeyPrefix is the prefix to retrieve all Providers
	ProvidersKeyPrefix       = "Providers/value/"
	ActiveProvidersKeyPrefix = "ActiveProviders/value/"

	AttestationKeyPrefix = "Attestation/value/"
)

// ActiveProvidersKey returns the store key to retrieve an Active Provider from the index fields
func ActiveProvidersKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// ProvidersKey returns the store key to retrieve a Providers from the index fields
func ProvidersKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// AttestationKey returns the store key to retrieve a Providers from the index fields
func AttestationKey(
	cid string,
) []byte {
	var key []byte

	cidBytes := []byte(cid)
	key = append(key, cidBytes...)
	key = append(key, []byte("/")...)

	return key
}
