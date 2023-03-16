package types

import (
	"regexp"
	"strings"
)

var rnsRegexp = *regexp.MustCompile(`^[\w-]+$`)

func IsValidName(name string) bool {
	return rnsRegexp.MatchString(name)
}

func GetTLD(name string) (string, error) {
	for _, tld := range SupportedTLDs {
		tldSize := len(tld)

		if tldSize+1 >= len(name) {
			return "", ErrNoTLD
		}

		checkingName := name[len(name)-tldSize:]

		if checkingName == tld {
			return tld, nil
		}
	}

	return "", ErrNoTLD
}

func GetSubdomain(name string) (string, string, bool) {
	if !strings.Contains(name, ".") {
		return "", name, false
	}

	s := strings.Split(name, ".")

	return s[0], s[1], true
}

func removeTLD(name string, tld string) (string, error) {
	tldSize := len(tld)

	if tldSize+1 >= len(name) {
		return "", ErrNoTLD
	}

	checkingName := name[:len(name)-tldSize-1]

	return checkingName, nil
}

func GetNameAndTLD(full string) (string, string, error) {
	tld, err := GetTLD(full)
	if err != nil {
		return "", "", err
	}

	name, err := removeTLD(full, tld)
	if err != nil {
		return "", "", err
	}

	return name, tld, nil
}

func GetCost(tld string) int64 {
	cost := TLDCost[tld]
	return cost
}
