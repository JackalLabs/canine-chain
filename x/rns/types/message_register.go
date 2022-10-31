package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegister = "register"

var _ sdk.Msg = &MsgRegister{}

func NewMsgRegister(creator string, name string, years string, data string) *MsgRegister {
	return &MsgRegister{
		Creator: creator,
		Name:    name,
		Years:   years,
		Data:    data,
	}
}

func (msg *MsgRegister) Route() string {
	return RouterKey
}

func (msg *MsgRegister) Type() string {
	return TypeMsgRegister
}

func (msg *MsgRegister) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegister) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, _, err = GetNameAndTLD(msg.Name)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid name/tld (%s)", err)
	}
	_, ok := sdk.NewIntFromString(msg.Years)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid years")
	}
	return nil
}
