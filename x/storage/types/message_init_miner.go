package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitMiner = "init_miner"

var _ sdk.Msg = &MsgInitMiner{}

func NewMsgInitMiner(creator string, ip string, totalspace string) *MsgInitMiner {
	return &MsgInitMiner{
		Creator:    creator,
		Ip:         ip,
		Totalspace: totalspace,
	}
}

func (msg *MsgInitMiner) Route() string {
	return RouterKey
}

func (msg *MsgInitMiner) Type() string {
	return TypeMsgInitMiner
}

func (msg *MsgInitMiner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitMiner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitMiner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
