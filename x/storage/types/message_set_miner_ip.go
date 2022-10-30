package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderIP = "set_provider_ip"

var _ sdk.Msg = &MsgSetProviderIp{}

func NewMsgSetProviderIP(creator string, ip string) *MsgSetProviderIp {
	return &MsgSetProviderIp{
		Creator: creator,
		Ip:      ip,
	}
}

func (msg *MsgSetProviderIp) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderIp) Type() string {
	return TypeMsgSetProviderIP
}

func (msg *MsgSetProviderIp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderIp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderIp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
