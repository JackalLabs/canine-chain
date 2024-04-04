package keeper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func HasViewingAccess(file types.Files, user string) (bool, error) {
	pvacc := file.ViewingAccess
	trackingNumber := file.TrackingNumber

	jvacc := make(map[string]string)
	err := json.Unmarshal([]byte(pvacc), &jvacc)
	if err != nil {
		return false, sdkerrors.Wrapf(err, "cannot unmarshal viewers: %s", pvacc)
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("v%s%s", trackingNumber, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	_, ok := jvacc[addressString]

	return ok, nil
}

func HasEditAccess(file types.Files, user string) (bool, error) {
	peacc := file.EditAccess
	trackingNumber := file.TrackingNumber

	jeacc := make(map[string]string)

	err := json.Unmarshal([]byte(peacc), &jeacc)
	if err != nil {
		return false, sdkerrors.Wrapf(err, "cannot unmarshal editors: %s", peacc)
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, user)))
	hash := h.Sum(nil)

	addressString := fmt.Sprintf("%x", hash)

	_, ok := jeacc[addressString]

	return ok, nil
}

func IsOwner(file types.Files, user string) bool {
	merklePath := file.Address

	h := sha256.New()
	h.Write([]byte(user))
	hash := h.Sum(nil)
	accountHash := fmt.Sprintf("%x", hash)

	// h1 is so named as to differentiate it from h above--else compiler complains
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

// MakeOwnerAddress Owner address is whoever owns this file/folder
func MakeOwnerAddress(merklePath string, user string) string {
	// make sure that user was already hex(hashed) before it was passed into
	// this function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, user)))
	hash := h.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash)

	return ownerAddress
}
