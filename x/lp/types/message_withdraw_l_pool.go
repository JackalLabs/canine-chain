package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawLPool = "withdraw_l_pool"

var _ sdk.Msg = &MsgWithdrawLPool{}

func NewMsgWithdrawLPool(creator string, poolName string, shares int64) *MsgWithdrawLPool {
	return &MsgWithdrawLPool{
		Creator:  creator,
		PoolName: poolName,
		Shares:   shares,
	}
}

func (msg *MsgWithdrawLPool) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawLPool) Type() string {
	return TypeMsgWithdrawLPool
}

func (msg *MsgWithdrawLPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawLPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawLPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
