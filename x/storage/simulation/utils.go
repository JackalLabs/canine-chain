package simulation

import (
	"math/rand"
	"strconv"
	"strings"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	sdksim "github.com/cosmos/cosmos-sdk/types/simulation"
)

// Generate random IPv4 url e.g http://1.1.1.1
// It may not be unique
func RandIPv4Url(r *rand.Rand) string {
	b := make([]string, 4)

	for i := 0; i < len(b); i++ {
		b[i] = strconv.Itoa(sdksim.RandIntBetween(r, 0, 255))
	}

	return "http://" + strings.Join(b, ".")
}
