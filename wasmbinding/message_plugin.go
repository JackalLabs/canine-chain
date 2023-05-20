package wasmbinding

import (
	"errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	filetreekeeper "github.com/jackalLabs/canine-chain/x/filetree/keeper"
)

// PROBABLY DON'T NEED THIS

// Change to GetFiles
// RetrieveFiles is a function, not method, so the message_plugin can use it
// address is the on-chain address/path of the file and owner is the owner address
func RetrieveFiles(ctx sdk.Context, filetree *filetreekeeper.Keeper, owner string, address string) (string, error) {
	// Address validation
	if _, err := parseAddress(owner); err != nil {
		return "", err
	}
	// Careful to make sure it only errors if not found
	// WrapGetFiles uses built in String() method returned by files.pb.go class file
	// Be warned that the struct will be outputted on one line
	Files, err := wrapGetFiles(ctx, filetree, owner, address)
	if err != nil {
		return "", errorsmod.Wrap(err, "validate sub-denom")
	}

	return Files, nil
}

// TO REPLACE
func wrapGetFiles(ctx sdk.Context, filetree *filetreekeeper.Keeper, owner string, address string) (string, error) {
	Files, found := filetree.GetFiles(ctx, address, owner)
	if !found {
		return "", errors.New("files not found")
	}

	return Files.String(), nil
}

// parseAddress parses address from bech32 string and verifies its format.
func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errorsmod.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, errorsmod.Wrap(err, "verify address format")
	}
	return parsed, nil
}
