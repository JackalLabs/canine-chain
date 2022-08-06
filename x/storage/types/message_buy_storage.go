package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyStorage = "buy_storage"

var _ sdk.Msg = &MsgBuyStorage{}

func NewMsgBuyStorage(creator string, forAddress string, duration string, bytes string, paymentDenom string) *MsgBuyStorage {
	return &MsgBuyStorage{
		Creator:      creator,
		ForAddress:   forAddress,
		Duration:     duration,
		Bytes:        bytes,
		PaymentDenom: paymentDenom,
	}
}

func (msg *MsgBuyStorage) Route() string {
	return RouterKey
}

func (msg *MsgBuyStorage) Type() string {
	return TypeMsgBuyStorage
}

func (msg *MsgBuyStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
