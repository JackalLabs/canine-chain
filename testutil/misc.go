package testutil

import (
	"math/rand"
	"time"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
)

// createTestAddresses retunrs a []string of random addresses.
func createTestAddresses(prefix string, n int) ([]string, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randomAccounts := simulation.RandomAccounts(r, n)
	var s []string

	for i := 0; i < n; i++ {
		b := randomAccounts[i].PubKey.Bytes()
		address, err := sdkTypes.Bech32ifyAddressBytes(prefix, b)
		if err != nil {
			return nil, err
		}

		s = append(s, address)
	}

	return s, nil
}
