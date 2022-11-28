package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteFile = "delete_file"

var _ sdk.Msg = &MsgDeleteFile{}

func NewMsgDeleteFile(creator string, hashPath string, account string) *MsgDeleteFile {
	return &MsgDeleteFile{
		Creator:  creator,
		HashPath: hashPath,
		Account:  account,
	}
}

func (msg *MsgDeleteFile) Route() string {
	return RouterKey
}

func (msg *MsgDeleteFile) Type() string {
	return TypeMsgDeleteFile
}

func (msg *MsgDeleteFile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteFile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteFile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Check empty values
	if msg.HashPath == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, 
			"invalid hash path: %s", msg.HashPath)
	}
	if msg.Account == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, 
			"invalid account: %s", msg.Account)
	}

	return nil
}
