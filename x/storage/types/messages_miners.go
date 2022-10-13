package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProviders = "create_providers"
	TypeMsgUpdateProviders = "update_providers"
	TypeMsgDeleteProviders = "delete_providers"
)

var _ sdk.Msg = &MsgCreateProviders{}

func NewMsgCreateProviders(
	creator string,
	address string,
	ip string,
	totalspace string,

) *MsgCreateProviders {
	return &MsgCreateProviders{
		Creator:    creator,
		Address:    address,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgCreateProviders) Route() string {
	return RouterKey
}

func (msg *MsgCreateProviders) Type() string {
	return TypeMsgCreateProviders
}

func (msg *MsgCreateProviders) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProviders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProviders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProviders{}

func NewMsgUpdateProviders(
	creator string,
	address string,
	ip string,
	totalspace string,

) *MsgUpdateProviders {
	return &MsgUpdateProviders{
		Creator:    creator,
		Address:    address,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgUpdateProviders) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProviders) Type() string {
	return TypeMsgUpdateProviders
}

func (msg *MsgUpdateProviders) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProviders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProviders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProviders{}

func NewMsgDeleteProviders(
	creator string,
	address string,

) *MsgDeleteProviders {
	return &MsgDeleteProviders{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteProviders) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProviders) Type() string {
	return TypeMsgDeleteProviders
}

func (msg *MsgDeleteProviders) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProviders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProviders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
