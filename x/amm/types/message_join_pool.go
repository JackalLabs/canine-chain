package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJoinPool = "join_pool"

var _ sdk.Msg = &MsgJoinPool{}

func NewMsgJoinPool(creator string, poolName string, coins sdk.Coins) *MsgJoinPool {
	return &MsgJoinPool{
		Creator:      creator,
		PoolName:     poolName,
		Coins:        coins,
	}
}

func (msg *MsgJoinPool) Route() string {
	return RouterKey
}

func (msg *MsgJoinPool) Type() string {
	return TypeMsgJoinPool
}

func (msg *MsgJoinPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJoinPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoinPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if len(msg.PoolName) == 0 {
		return ErrInvalidPoolName
	}

	coins := sdk.NewCoins(msg.Coins...)
	if !coins.IsAllPositive() {
		return sdkerrors.Wrapf(ErrInvalidValue,
			"amount of coins are too small: (%s)", coins.String())
	}

	return nil
}
