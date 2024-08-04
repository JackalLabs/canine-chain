package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddEditors = "add_editors"

var _ sdk.Msg = &MsgAddEditors{}

func NewMsgAddEditors(creator string, editorIDs string, editorKeys string, address string, fileowner string) *MsgAddEditors {
	return &MsgAddEditors{
		Creator:    creator,
		EditorIds:  editorIDs,
		EditorKeys: editorKeys,
		Address:    address,
		FileOwner:  fileowner,
	}
}

func (msg *MsgAddEditors) Route() string {
	return RouterKey
}

func (msg *MsgAddEditors) Type() string {
	return TypeMsgAddEditors
}

func (msg *MsgAddEditors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddEditors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddEditors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Check empty values
	if msg.EditorIds == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid editor ids: %s", msg.EditorIds)
	}
	if msg.EditorKeys == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid editor keys: %s", msg.EditorKeys)
	}
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid address: %s", msg.Address)
	}
	if msg.FileOwner == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid file owner: %s", msg.FileOwner)
	}

	return nil
}
