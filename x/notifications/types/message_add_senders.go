package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddSenders = "add_senders"

var _ sdk.Msg = &MsgAddSenders{}

func NewMsgAddSenders(creator string, senderIds string) *MsgAddSenders {
  return &MsgAddSenders{
		Creator: creator,
    SenderIds: senderIds,
	}
}

func (msg *MsgAddSenders) Route() string {
  return RouterKey
}

func (msg *MsgAddSenders) Type() string {
  return TypeMsgAddSenders
}

func (msg *MsgAddSenders) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgAddSenders) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgAddSenders) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

