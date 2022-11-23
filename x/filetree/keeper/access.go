package keeper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func HasEditAccess(file types.Files, user string) (bool, error) {
	peacc := file.EditAccess
	trackingNumber := file.TrackingNumber

	jeacc := make(map[string]string)

	if err := json.Unmarshal([]byte(peacc), &jeacc); err != nil {
		return false, err	
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	if _, ok := jeacc[addressString]; ok {
		return ok, nil
	}

	//During sandbox testing, if editor doesn't exist, the body of the if statement never executes, so we need to return false
	return false, nil
}

func IsOwner(file types.Files, user string) bool {

	merklePath := file.Address

	h := sha256.New()
	h.Write([]byte(user))
	hash := h.Sum(nil)
	accountHash := fmt.Sprintf("%x", hash)

	//h1 is so named as to differentiate it from h above--else compiler complains
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
	//make sure that user was already hex(hashed) before it was passed into
	//this function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, user)))
	hash := h.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash)

	return ownerAddress
}

