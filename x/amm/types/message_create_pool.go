package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePool = "create_pool"

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(
	creator string,
	coins sdk.Coins,
	aMMId uint32,
	swapFeeMulti sdk.Dec,
	minLockDuration int64,
	penaltyMulti sdk.Dec,
) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:         creator,
		Coins:           coins,
		AmmId:          aMMId,
		SwapFeeMulti:    swapFeeMulti.String(),
		MinLockDuration: minLockDuration,
		PenaltyMulti:    penaltyMulti.String(),
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	if msg == nil {
		return sdkerrors.ErrInvalidRequest
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.MinLockDuration < 0 {
		return sdkerrors.Wrap(errors.New(fmt.Sprintf(
			"MinLockDuration cannot be negative: %d > 0", msg.MinLockDuration)),
			sdkerrors.ErrInvalidRequest.Error())
	}

	sfm, err := sdk.NewDecFromStr(msg.SwapFeeMulti)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if sfm.LT(sdk.NewDec(0)) || sfm.GTE(sdk.NewDec(1)) {
		return sdkerrors.Wrap(errors.New(fmt.Sprintf(
			"swap fee multiplier should be non-negative number less than 1:"+
				" 0 < %s < 1", msg.SwapFeeMulti)),
			sdkerrors.ErrInvalidRequest.Error())
	}

	pm, err := sdk.NewDecFromStr(msg.PenaltyMulti)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if pm.LT(sdk.NewDec(0)) || pm.GTE(sdk.NewDec(1)) {
		return sdkerrors.Wrap(errors.New(fmt.Sprintf(
			"penalty multiplier fee multiplier should be non-negative number "+
				"and less than 1: 0 < %s < 1", msg.PenaltyMulti)),
			sdkerrors.ErrInvalidRequest.Error())
	}

	return nil
}
