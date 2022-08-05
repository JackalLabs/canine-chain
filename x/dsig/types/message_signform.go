package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSignform = "signform"

var _ sdk.Msg = &MsgSignform{}

func NewMsgSignform(creator string, ffid string, vote int32) *MsgSignform {
	return &MsgSignform{
		Creator: creator,
		Ffid:    ffid,
		Vote:    vote,
	}
}

func (msg *MsgSignform) Route() string {
	return RouterKey
}

func (msg *MsgSignform) Type() string {
	return TypeMsgSignform
}

func (msg *MsgSignform) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignform) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
