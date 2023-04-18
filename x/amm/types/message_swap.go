package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwap = "swap"

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(creator string, poolId uint64, coinInput sdk.Coin, minCoinOutput sdk.Coin) *MsgSwap {
	return &MsgSwap{
		Creator:       creator,
		PoolId:      poolId,
		CoinInput:     coinInput,
		MinCoinOutput: minCoinOutput,
	}
}

func (msg *MsgSwap) Route() string {
	return RouterKey
}

func (msg *MsgSwap) Type() string {
	return TypeMsgSwap
}

func (msg *MsgSwap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !msg.CoinInput.IsPositive() {
		return sdkerrors.Wrapf(ErrNegativeCoin, "coin input is too small (%s)",
			msg.CoinInput.String())
	}

	if !msg.MinCoinOutput.IsPositive() {
		return sdkerrors.Wrapf(ErrNegativeCoin, "minimum coin output is too small (%s)",
			msg.MinCoinOutput.String())
	}
	return nil
}
