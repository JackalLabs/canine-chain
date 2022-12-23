package simulation

import (
	"math/rand"
	"strconv"
	"strings"

	sdksim "github.com/cosmos/cosmos-sdk/types/simulation"
)

// Generate random IPv4 address
// It may not be unique
func RandIPv4(r *rand.Rand) string {
	b := make([]string, 4)

	for i := 0; i < len(b); i++ {
		b[i] = strconv.Itoa(sdksim.RandIntBetween(r, 0, 255))
	}

	return strings.Join(b, ".")
}
