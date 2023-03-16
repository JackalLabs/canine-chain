package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExitPool = "exit_pool"

var _ sdk.Msg = &MsgExitPool{}

func NewMsgExitPool(creator string, poolName string, shares int64) *MsgExitPool {
	return &MsgExitPool{
		Creator:  creator,
		PoolName: poolName,
		Shares:   shares,
	}
}

func (msg *MsgExitPool) Route() string {
	return RouterKey
}

func (msg *MsgExitPool) Type() string {
	return TypeMsgExitPool
}

func (msg *MsgExitPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExitPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExitPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
