package types

import (
	fmt "fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != addressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.ForAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != addressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	if _, err := strconv.Atoi(msg.Bytes); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse bytes (%s)", err)
	}

	if _, err := strconv.Atoi(msg.Duration); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse bytes (%s)", err)
	}

	return nil
}
