package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgList = "list"

var _ sdk.Msg = &MsgList{}

func NewMsgList(creator string, name string, price string) *MsgList {
	return &MsgList{
		Creator: creator,
		Name:    name,
		Price:   price,
	}
}

func (msg *MsgList) Route() string {
	return RouterKey
}

func (msg *MsgList) Type() string {
	return TypeMsgList
}

func (msg *MsgList) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
