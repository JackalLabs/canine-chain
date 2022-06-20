package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSaveRequests{}

func NewMsgCreateSaveRequests(
	creator string,
	index string,
	size string,
	approved string,

) *MsgCreateSaveRequests {
	return &MsgCreateSaveRequests{
		Creator:  creator,
		Index:    index,
		Size_:    size,
		Approved: approved,
	}
}

func (msg *MsgCreateSaveRequests) Route() string {
	return RouterKey
}

func (msg *MsgCreateSaveRequests) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSaveRequests) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSaveRequests) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSaveRequests{}

func NewMsgUpdateSaveRequests(
	creator string,
	index string,
	size string,
	approved string,

) *MsgUpdateSaveRequests {
	return &MsgUpdateSaveRequests{
		Creator:  creator,
		Index:    index,
		Size_:    size,
		Approved: approved,
	}
}

func (msg *MsgUpdateSaveRequests) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSaveRequests) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSaveRequests) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSaveRequests) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSaveRequests{}

func NewMsgDeleteSaveRequests(
	creator string,
	index string,

) *MsgDeleteSaveRequests {
	return &MsgDeleteSaveRequests{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteSaveRequests) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSaveRequests) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSaveRequests) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSaveRequests) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
