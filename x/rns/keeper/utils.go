package keeper

import (
	"fmt"
	"strings"

	"github.com/jackal-dao/canine/x/rns/types"
)

func getTLD(name string) (string, error) {

	for _, tld := range types.SUPPORTED_TLDS {
		tldSize := len(tld)

		if tldSize+1 >= len(name) {
			return "", types.ErrNoTLD
		}

		checkingName := name[len(name)-tldSize:]

		if checkingName == tld {
			return tld, nil
		}
	}

	return "", types.ErrNoTLD
}

func getSubdomain(name string) (string, string, bool) {
	if !strings.Contains(name, ".") {
		return "", name, false
	}

	s := strings.Split(name, ".")

	return s[0], s[1], true

}

func removeTLD(name string, tld string) (string, error) {
	tldSize := len(tld)

	if tldSize+1 >= len(name) {
		return "", types.ErrNoTLD
	}

	checkingName := name[:len(name)-tldSize-1]

	return checkingName, nil
}

func getNameAndTLD(full string) (string, string, error) {
	tld, err := getTLD(full)
	if err != nil {
		return "", "", err
	}

	name, err := removeTLD(full, tld)
	if err != nil {
		return "", "", err
	}

	return name, tld, nil
}

func getCost(tld string) int64 {
	cost := types.TLD_COST[tld]
	fmt.Printf("COST: %d\n", cost)
	return cost
}
