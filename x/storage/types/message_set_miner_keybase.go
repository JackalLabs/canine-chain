package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderKeybase = "set_provider_keybase"

var _ sdk.Msg = &MsgSetProviderKeybase{}

func NewMsgSetProviderKeybase(creator string, key string) *MsgSetProviderKeybase {
	return &MsgSetProviderKeybase{
		Creator: creator,
		Keybase: key,
	}
}

func (msg *MsgSetProviderKeybase) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderKeybase) Type() string {
	return TypeMsgSetProviderKeybase
}

func (msg *MsgSetProviderKeybase) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderKeybase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderKeybase) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
