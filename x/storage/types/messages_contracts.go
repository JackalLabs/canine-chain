package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateContracts = "create_contracts"
	TypeMsgUpdateContracts = "update_contracts"
	TypeMsgDeleteContracts = "delete_contracts"
)

var _ sdk.Msg = &MsgCreateContracts{}

func NewMsgCreateContracts(
	creator string,
	cid string,
	chunks string,
	merkle string,
	signee string,
	duration string,
	filesize string,
	fid string,

) *MsgCreateContracts {
	return &MsgCreateContracts{
		Creator:  creator,
		Cid:      cid,
		Chunks:   chunks,
		Merkle:   merkle,
		Signee:   signee,
		Duration: duration,
		Filesize: filesize,
		Fid:      fid,
	}
}

func (msg *MsgCreateContracts) Route() string {
	return RouterKey
}

func (msg *MsgCreateContracts) Type() string {
	return TypeMsgCreateContracts
}

func (msg *MsgCreateContracts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContracts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContracts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateContracts{}

func NewMsgUpdateContracts(
	creator string,
	cid string,
	chunks string,
	merkle string,
	signee string,
	duration string,
	filesize string,
	fid string,

) *MsgUpdateContracts {
	return &MsgUpdateContracts{
		Creator:  creator,
		Cid:      cid,
		Chunks:   chunks,
		Merkle:   merkle,
		Signee:   signee,
		Duration: duration,
		Filesize: filesize,
		Fid:      fid,
	}
}

func (msg *MsgUpdateContracts) Route() string {
	return RouterKey
}

func (msg *MsgUpdateContracts) Type() string {
	return TypeMsgUpdateContracts
}

func (msg *MsgUpdateContracts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateContracts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateContracts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteContracts{}

func NewMsgDeleteContracts(
	creator string,
	cid string,

) *MsgDeleteContracts {
	return &MsgDeleteContracts{
		Creator: creator,
		Cid:     cid,
	}
}
func (msg *MsgDeleteContracts) Route() string {
	return RouterKey
}

func (msg *MsgDeleteContracts) Type() string {
	return TypeMsgDeleteContracts
}

func (msg *MsgDeleteContracts) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteContracts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteContracts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
