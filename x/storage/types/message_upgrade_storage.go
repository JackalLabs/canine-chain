package types

import (
	fmt "fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpgradeStorage = "upgrade_storage"

var _ sdk.Msg = &MsgUpgradeStorage{}

func NewMsgUpgradeStorage(creator string, forAddress string, duration string, bytes string, paymentDenom string) *MsgUpgradeStorage {
	return &MsgUpgradeStorage{
		Creator:      creator,
		ForAddress:   forAddress,
		Duration:     duration,
		Bytes:        bytes,
		PaymentDenom: paymentDenom,
	}
}

func (msg *MsgUpgradeStorage) Route() string {
	return RouterKey
}

func (msg *MsgUpgradeStorage) Type() string {
	return TypeMsgUpgradeStorage
}

func (msg *MsgUpgradeStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpgradeStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpgradeStorage) ValidateBasic() error {
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
