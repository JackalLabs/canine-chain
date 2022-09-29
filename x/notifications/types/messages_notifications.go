package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateNotifications = "create_notifications"
	TypeMsgUpdateNotifications = "update_notifications"
	TypeMsgDeleteNotifications = "delete_notifications"
)

var _ sdk.Msg = &MsgCreateNotifications{}

func NewMsgCreateNotifications(
	creator string,
	count uint64,
	notification string,
	address string,

) *MsgCreateNotifications {
	return &MsgCreateNotifications{
		Creator:      creator,
		Count:        count,
		Notification: notification,
		Address:      address,
	}
}

func (msg *MsgCreateNotifications) Route() string {
	return RouterKey
}

func (msg *MsgCreateNotifications) Type() string {
	return TypeMsgCreateNotifications
}

func (msg *MsgCreateNotifications) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNotifications) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNotifications) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNotifications{}

func NewMsgUpdateNotifications(
	creator string,
	count uint64,
	notification string,
	address string,

) *MsgUpdateNotifications {
	return &MsgUpdateNotifications{
		Creator:      creator,
		Count:        count,
		Notification: notification,
		Address:      address,
	}
}

func (msg *MsgUpdateNotifications) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNotifications) Type() string {
	return TypeMsgUpdateNotifications
}

func (msg *MsgUpdateNotifications) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNotifications) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNotifications) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNotifications{}

func NewMsgDeleteNotifications(
	creator string,
	count uint64,

) *MsgDeleteNotifications {
	return &MsgDeleteNotifications{
		Creator: creator,
		Count:   count,
	}
}
func (msg *MsgDeleteNotifications) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNotifications) Type() string {
	return TypeMsgDeleteNotifications
}

func (msg *MsgDeleteNotifications) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNotifications) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNotifications) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
