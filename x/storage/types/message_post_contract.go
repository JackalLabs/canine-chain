package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostContract = "post_contract"

var _ sdk.Msg = &MsgPostContract{}

func NewMsgPostContract(creator string, signee string, duration string, filesize string, fid string, merkle string) *MsgPostContract {
	return &MsgPostContract{
		Creator:  creator,
		Signee:   signee,
		Duration: duration,
		Filesize: filesize,
		Fid:      fid,
		Merkle:   merkle,
	}
}

func (msg *MsgPostContract) Route() string {
	return RouterKey
}

func (msg *MsgPostContract) Type() string {
	return TypeMsgPostContract
}

func (msg *MsgPostContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
