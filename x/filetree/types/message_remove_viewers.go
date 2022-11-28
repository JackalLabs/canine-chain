package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveViewers = "remove_viewers"

var _ sdk.Msg = &MsgRemoveViewers{}

func NewMsgRemoveViewers(creator string, viewerIds string, address string, fileowner string) *MsgRemoveViewers {
	return &MsgRemoveViewers{
		Creator:   creator,
		ViewerIds: viewerIds,
		Address:   address,
		Fileowner: fileowner,
	}
}

func (msg *MsgRemoveViewers) Route() string {
	return RouterKey
}

func (msg *MsgRemoveViewers) Type() string {
	return TypeMsgRemoveViewers
}

func (msg *MsgRemoveViewers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveViewers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveViewers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ViewerIds == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid viewer ids: %s", msg.ViewerIds)
	}
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid address: %s", msg.Address)
	}
	if msg.Fileowner == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid file owner: %s", msg.Fileowner)
	}

	return nil
}
