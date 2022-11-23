package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddEditors = "add_editors"

var _ sdk.Msg = &MsgAddEditors{}

func NewMsgAddEditors(creator string, editorIds string, editorKeys string, address string, fileowner string, notifyEditors string, notiForEditors string) *MsgAddEditors {
	return &MsgAddEditors{
		Creator:        creator,
		EditorIds:      editorIds,
		EditorKeys:     editorKeys,
		Address:        address,
		Fileowner:      fileowner,
		NotifyEditors:  notifyEditors,
		NotiForEditors: notiForEditors,
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
	return nil
}
