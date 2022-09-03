package types

var SUPPORTED_TLDS = []string{"ibc", "jkl"}

var IS_RESERVED = map[string]bool{"ibc": false, "jkl": false}

var TLD_COST = map[string]int64{"ibc": 11250000, "jkl": 7500000}
