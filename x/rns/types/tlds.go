package types

var SupportedTLDs = []string{"ibc", "jkl"}

var IsReserved = map[string]bool{"ibc": false, "jkl": false}

var TLDCost = map[string]int64{"ibc": 50000000, "jkl": 25000000}
