package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLPool = "create_l_pool"

var _ sdk.Msg = &MsgCreateLPool{}

func NewMsgCreateLPool(
	creator string,
	coins sdk.DecCoins,
	aMMId uint32,
	swapFeeMulti sdk.Dec,
	minLockDuration int64,
	penaltyMulti sdk.Dec,
) *MsgCreateLPool {
	return &MsgCreateLPool{
		Creator:         creator,
		Coins:           coins,
		Amm_Id:          aMMId,
		SwapFeeMulti:    swapFeeMulti.String(),
		MinLockDuration: minLockDuration,
		PenaltyMulti:    penaltyMulti.String(),
	}
}

func (msg *MsgCreateLPool) Route() string {
	return RouterKey
}

func (msg *MsgCreateLPool) Type() string {
	return TypeMsgCreateLPool
}

func (msg *MsgCreateLPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLPool) ValidateBasic() error {
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
