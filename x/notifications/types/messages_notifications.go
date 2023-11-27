package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const (
	TypeMsgCreateNotification = "create_notification"
	TypeMsgDeleteNotification = "delete_notification"
)

var _ sdk.Msg = &MsgCreateNotification{}

func NewMsgCreateNotification(
	from string,
	to string,
	contents string,
) *MsgCreateNotification {
	return &MsgCreateNotification{
		Creator:  from,
		To:       to,
		Contents: contents,
	}
}

func (msg *MsgCreateNotification) Route() string {
	return RouterKey
}

func (msg *MsgCreateNotification) Type() string {
	return TypeMsgCreateNotification
}

func (msg *MsgCreateNotification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNotification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNotification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !json.Valid([]byte(msg.Contents)) {
		return sdkerrors.Wrapf(ErrInvalidContents, "cannot verify contents")
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteNotification{}

func NewMsgDeleteNotification(
	to string,
	from string,
	timeStamp time.Time,
) *MsgDeleteNotification {
	return &MsgDeleteNotification{
		Creator: to,
		From:    from,
		Time:    timeStamp,
	}
}

func (msg *MsgDeleteNotification) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNotification) Type() string {
	return TypeMsgDeleteNotification
}

func (msg *MsgDeleteNotification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNotification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNotification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
