package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProofs = "create_proofs"
	TypeMsgUpdateProofs = "update_proofs"
	TypeMsgDeleteProofs = "delete_proofs"
)

var _ sdk.Msg = &MsgCreateProofs{}

func NewMsgCreateProofs(
	creator string,
	cid string,
	item string,
	hashes string,

) *MsgCreateProofs {
	return &MsgCreateProofs{
		Creator: creator,
		Cid:     cid,
		Item:    item,
		Hashes:  hashes,
	}
}

func (msg *MsgCreateProofs) Route() string {
	return RouterKey
}

func (msg *MsgCreateProofs) Type() string {
	return TypeMsgCreateProofs
}

func (msg *MsgCreateProofs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProofs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProofs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProofs{}

func NewMsgUpdateProofs(
	creator string,
	cid string,
	item string,
	hashes string,

) *MsgUpdateProofs {
	return &MsgUpdateProofs{
		Creator: creator,
		Cid:     cid,
		Item:    item,
		Hashes:  hashes,
	}
}

func (msg *MsgUpdateProofs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProofs) Type() string {
	return TypeMsgUpdateProofs
}

func (msg *MsgUpdateProofs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProofs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProofs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProofs{}

func NewMsgDeleteProofs(
	creator string,
	cid string,

) *MsgDeleteProofs {
	return &MsgDeleteProofs{
		Creator: creator,
		Cid:     cid,
	}
}
func (msg *MsgDeleteProofs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProofs) Type() string {
	return TypeMsgDeleteProofs
}

func (msg *MsgDeleteProofs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProofs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProofs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
