package keeper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jackal-dao/canine/x/filetree/types"
)

func HasViewingAccess(file types.Files, user string) bool {
	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("v%s%s", file.Address, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	if _, ok := jvacc[addressString]; ok {
		return ok
	}

	return true
}

func HasEditAccess(file types.Files, user string) bool {
	pvacc := file.EditAccess

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("e%s%s", file.Address, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	if _, ok := jvacc[addressString]; ok {
		return ok
	}

	return true
}

func IsOwner(file types.Files, user string) bool {

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", file.Address, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	return addressString == file.Owner
}

func MakeAddress(path string, user string) string {

	h := sha256.New()
	h.Write([]byte(path))
	hash := h.Sum(nil)

	pathString := fmt.Sprintf("%x", hash)
	return pathString
}

func MakeChainAddress(path string, user string) string {

	h := sha256.New()
	h.Write([]byte(path))
	hash := h.Sum(nil)

	pathString := fmt.Sprintf("%x", hash)

	h = sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", user, pathString)))
	hash = h.Sum(nil)

	pathString = fmt.Sprintf("%x", hash)

	return pathString
}

func MakeViewerAddress(path string, user string) string {

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("v%s%s", path, user)))
	hash := h.Sum(nil)
	addressString := fmt.Sprintf("%x", hash)

	return addressString
}
