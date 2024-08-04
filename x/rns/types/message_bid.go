package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBid = "bid"

var _ sdk.Msg = &MsgBid{}

func NewMsgBid(creator string, name string, bid sdk.Coin) *MsgBid {
	return &MsgBid{
		Creator: creator,
		Name:    name,
		Bid:     bid,
	}
}

func (msg *MsgBid) Route() string {
	return RouterKey
}

func (msg *MsgBid) Type() string {
	return TypeMsgBid
}

func (msg *MsgBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, _, err = GetNameAndTLD(msg.Name)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid name/tld (%s)", err)
	}
	return nil
}
