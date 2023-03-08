package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBlockSenders = "add_senders"

var _ sdk.Msg = &MsgBlockSenders{}

func NewMsgBlockSenders(creator string, senderIds string) *MsgBlockSenders {
	return &MsgBlockSenders{
		Creator:   creator,
		SenderIds: senderIds,
	}
}

func (msg *MsgBlockSenders) Route() string {
	return RouterKey
}

func (msg *MsgBlockSenders) Type() string {
	return TypeMsgBlockSenders
}

func (msg *MsgBlockSenders) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBlockSenders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBlockSenders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
