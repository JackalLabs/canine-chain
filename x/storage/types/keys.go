package types

import (
	"crypto/sha256"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ModuleName defines the module name
	ModuleName = "storage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_storage"

	AddressPrefix = "jkl"
	CidPrefix     = "jklc"
	FidPrefix     = "jklf"

	CollateralCollectorName    = "storage_collateral_name"
	TokenHolderName            = "token_holder_name"
	ProtocolOwnedLiquidityName = "protocol_owned_liq"
)

func GetTokenHolderAccount() (sdk.AccAddress, error) {
	return GetAccount(TokenHolderName)
}

func GetPOLAccount() (sdk.AccAddress, error) {
	return GetAccount(ProtocolOwnedLiquidityName)
}

func GetAccount(name string) (sdk.AccAddress, error) {
	s := sha256.New()
	s.Write([]byte(name))
	m := s.Sum(nil)
	mh := hex.EncodeToString(m)
	adr, err := sdk.AccAddressFromHex(mh)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get account account")
	}
	return adr, nil
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
