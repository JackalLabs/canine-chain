package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMiners = "create_miners"
	TypeMsgUpdateMiners = "update_miners"
	TypeMsgDeleteMiners = "delete_miners"
)

var _ sdk.Msg = &MsgCreateMiners{}

func NewMsgCreateMiners(
	creator string,
	address string,
	ip string,
	totalspace string,

) *MsgCreateMiners {
	return &MsgCreateMiners{
		Creator:    creator,
		Address:    address,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgCreateMiners) Route() string {
	return RouterKey
}

func (msg *MsgCreateMiners) Type() string {
	return TypeMsgCreateMiners
}

func (msg *MsgCreateMiners) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMiners) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMiners) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMiners{}

func NewMsgUpdateMiners(
	creator string,
	address string,
	ip string,
	totalspace string,

) *MsgUpdateMiners {
	return &MsgUpdateMiners{
		Creator:    creator,
		Address:    address,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgUpdateMiners) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMiners) Type() string {
	return TypeMsgUpdateMiners
}

func (msg *MsgUpdateMiners) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMiners) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMiners) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMiners{}

func NewMsgDeleteMiners(
	creator string,
	address string,

) *MsgDeleteMiners {
	return &MsgDeleteMiners{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteMiners) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMiners) Type() string {
	return TypeMsgDeleteMiners
}

func (msg *MsgDeleteMiners) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMiners) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMiners) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
