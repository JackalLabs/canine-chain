package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDelist = "delist"

var _ sdk.Msg = &MsgDelist{}

func NewMsgDeList(creator string, name string) *MsgDelist {
	return &MsgDelist{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDelist) Route() string {
	return RouterKey
}

func (msg *MsgDelist) Type() string {
	return TypeMsgDelist
}

func (msg *MsgDelist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDelist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, _, err = GetNameAndTLD(msg.Name)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid name/tld (%s)", err)
	}
	return nil
}
