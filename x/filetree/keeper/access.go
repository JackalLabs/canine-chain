package keeper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jackal-dao/canine/x/filetree/types"
)

func HasViewingAccess(file types.Files, user string) bool {
	pvacc := file.ViewingAccess
	trackingNumber := file.TrackingNumber

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("v%s%s", trackingNumber, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	if _, ok := jvacc[addressString]; ok {
		return ok
	}

	return true
}

func HasEditAccess(file types.Files, user string) bool {

	peacc := file.EditAccess
	trackingNumber := file.TrackingNumber

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(peacc), &jvacc)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)


	if _, ok := jvacc[addressString]; ok {
		return ok
	}


	return false
}

func IsOwner(file types.Files, user string) bool {
	merklePath := file.Address

	h := sha256.New()
	h.Write([]byte(user))
	hash := h.Sum(nil)
	accountHash := fmt.Sprintf("%x", hash)


	h1 := sha256.New()
	h1.Write([]byte(fmt.Sprintf("o%s%s", merklePath, accountHash)))
	hash1 := h1.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash1)

	return ownerAddress == file.Owner
}

func MakeViewerAddress(trackingNumber string, user string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("v%s%s", trackingNumber, user)))
	hash := h.Sum(nil)
	addressString := fmt.Sprintf("%x", hash)

	return addressString
}

func MakeEditorAddress(trackingNumber string, user string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, user)))
	hash := h.Sum(nil)
	addressString := fmt.Sprintf("%x", hash)

	return addressString
}

// Owner address is whoever owns this file/folder
func MakeOwnerAddress(merklePath string, user string) string {

	//this function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, user)))
	hash := h.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash)

	return ownerAddress
}

// Deleted old code: 'MakeAddress' and 'MakeChainAddress'
