package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMinerIp = "set_miner_ip"

var _ sdk.Msg = &MsgSetMinerIp{}

func NewMsgSetMinerIp(creator string, ip string) *MsgSetMinerIp {
	return &MsgSetMinerIp{
		Creator: creator,
		Ip:      ip,
	}
}

func (msg *MsgSetMinerIp) Route() string {
	return RouterKey
}

func (msg *MsgSetMinerIp) Type() string {
	return TypeMsgSetMinerIp
}

func (msg *MsgSetMinerIp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMinerIp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMinerIp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
