package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateform = "createform"

var _ sdk.Msg = &MsgCreateform{}

func NewMsgCreateform(creator string, fid string, signees string) *MsgCreateform {
	return &MsgCreateform{
		Creator: creator,
		Fid:     fid,
		Signees: signees,
	}
}

func (msg *MsgCreateform) Route() string {
	return RouterKey
}

func (msg *MsgCreateform) Type() string {
	return TypeMsgCreateform
}

func (msg *MsgCreateform) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateform) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
