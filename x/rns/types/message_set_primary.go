package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakePrimary = "make_primary"

var _ sdk.Msg = &MsgMakePrimary{}

func NewMsgMakePrimary(name string) *MsgMakePrimary {
	return &MsgMakePrimary{
		Name: name,
	}
}

func (msg *MsgMakePrimary) Route() string {
	return RouterKey
}

func (msg *MsgMakePrimary) Type() string {
	return TypeMsgMakePrimary
}

func (msg *MsgMakePrimary) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakePrimary) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakePrimary) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	name, _, err := GetNameAndTLD(msg.Name)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid name/tld (%s)", err)
	}
	if !IsValidName(name) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid name")
	}

	return nil
}
