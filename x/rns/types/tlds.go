package types

var SupportedTLDs = []string{"ibc", "jkl"}

var IsReserved = map[string]bool{"ibc": false, "jkl": false}

var TLDCost = map[string]int64{"ibc": 50_000_000, "jkl": 10_000_000}
