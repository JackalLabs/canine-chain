package types

var SupportedTLDs = []string{"ibc", "jkl"}

var IsReserved = map[string]bool{"ibc": false, "jkl": false}

var TLDCost = map[string]int64{"ibc": 11250000, "jkl": 7500000}
