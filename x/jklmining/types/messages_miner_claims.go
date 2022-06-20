package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMinerClaims = "create_miner_claims"
	TypeMsgUpdateMinerClaims = "update_miner_claims"
	TypeMsgDeleteMinerClaims = "delete_miner_claims"
)

var _ sdk.Msg = &MsgCreateMinerClaims{}

func NewMsgCreateMinerClaims(
	creator string,
	hash string,

) *MsgCreateMinerClaims {
	return &MsgCreateMinerClaims{
		Creator: creator,
		Hash:    hash,
	}
}

func (msg *MsgCreateMinerClaims) Route() string {
	return RouterKey
}

func (msg *MsgCreateMinerClaims) Type() string {
	return TypeMsgCreateMinerClaims
}

func (msg *MsgCreateMinerClaims) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMinerClaims) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMinerClaims) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMinerClaims{}

func NewMsgUpdateMinerClaims(
	creator string,
	hash string,

) *MsgUpdateMinerClaims {
	return &MsgUpdateMinerClaims{
		Creator: creator,
		Hash:    hash,
	}
}

func (msg *MsgUpdateMinerClaims) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMinerClaims) Type() string {
	return TypeMsgUpdateMinerClaims
}

func (msg *MsgUpdateMinerClaims) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMinerClaims) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMinerClaims) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMinerClaims{}

func NewMsgDeleteMinerClaims(
	creator string,
	hash string,

) *MsgDeleteMinerClaims {
	return &MsgDeleteMinerClaims{
		Creator: creator,
		Hash:    hash,
	}
}
func (msg *MsgDeleteMinerClaims) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMinerClaims) Type() string {
	return TypeMsgDeleteMinerClaims
}

func (msg *MsgDeleteMinerClaims) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMinerClaims) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMinerClaims) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
