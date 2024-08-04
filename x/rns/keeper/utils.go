package keeper

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func GetTLD(name string) (string, error) {
	for _, tld := range types.SupportedTLDs {
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

func GetSubdomain(name string) (string, string, bool) {
	if !strings.Contains(name, ".") {
		return "", name, false
	}

	s := strings.Split(name, ".")

	return s[0], s[1], true
}

func RemoveTLD(name string, tld string) (string, error) {
	tldSize := len(tld)

	if tldSize+1 >= len(name) {
		return "", types.ErrNoTLD
	}

	checkingName := name[:len(name)-tldSize-1]

	return checkingName, nil
}

func GetCostOfName(name string, tld string) (int64, error) {
	baseCost := GetCost(tld)

	var cost int64
	chars := len(name)
	switch chars {
	case 0:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Must be 1 or more characters.")
	case 1:
		cost = baseCost * 32
	case 2:
		cost = baseCost * 16
	case 3:
		cost = baseCost * 8
	case 4:
		cost = baseCost * 4
	case 5:
		cost = baseCost * 2
	default:
		cost = baseCost
	}

	return cost, nil
}

func GetNameAndTLD(full string) (string, string, error) {
	tld, err := GetTLD(full)
	if err != nil {
		return "", "", err
	}

	name, err := RemoveTLD(full, tld)
	if err != nil {
		return "", "", err
	}

	return name, tld, nil
}

func GetCost(tld string) int64 {
	cost := types.TLDCost[tld]
	return cost
}
