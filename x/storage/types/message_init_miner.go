package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitProvider = "init_provider"

var _ sdk.Msg = &MsgInitProvider{}

func NewMsgInitProvider(creator string, ip string, totalspace string) *MsgInitProvider {
	return &MsgInitProvider{
		Creator:    creator,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgInitProvider) Route() string {
	return RouterKey
}

func (msg *MsgInitProvider) Type() string {
	return TypeMsgInitProvider
}

func (msg *MsgInitProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
