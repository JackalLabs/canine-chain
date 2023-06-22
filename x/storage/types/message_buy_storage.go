package types

import (
	fmt "fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgBuyStorage      = "buy_storage"
	TypeMsgBuyStorageToken = "buy_storage_token"
)

var (
	_ sdk.Msg = &MsgBuyStorage{}
	_ sdk.Msg = &MsgBuyStorageToken{}
)

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
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.ForAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	if _, err := strconv.ParseInt(msg.Bytes, 10, 64); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse bytes (%s)", err)
	}

	duration, err := time.ParseDuration(msg.Duration)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse bytes (%s)", err)
	}

	if duration < 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duration cannot be negative (%s)", msg.Duration)
	}

	return nil
}

func NewMsgBuyStorageToken(creator string, amount int64, paymentDenom string) *MsgBuyStorageToken {
	return &MsgBuyStorageToken{
		Creator:      creator,
		Amount:       amount,
		PaymentDenom: paymentDenom,
	}
}

func (msg *MsgBuyStorageToken) Route() string {
	return RouterKey
}

func (msg *MsgBuyStorageToken) Type() string {
	return TypeMsgBuyStorageToken
}

func (msg *MsgBuyStorageToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyStorageToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyStorageToken) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "amount cannot be less than 1: %s", err)
	}

	return nil
}
