package types

import (
	fmt "fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostContract = "post_contract"

var _ sdk.Msg = &MsgPostContract{}

func NewMsgPostContract(creator string, signee string, filesize string, fid string, merkle string) *MsgPostContract {
	return &MsgPostContract{
		Creator:  creator,
		Signee:   signee,
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
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != addressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Signee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signee address (%s)", err)
	}
	if prefix != addressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signee prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Fid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fid (%s)", err)
	}
	if prefix != fidPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklf`", prefix))
	}

	if _, err := strconv.Atoi(msg.Filesize); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse file size (%s)", err)
	}

	return nil
}
