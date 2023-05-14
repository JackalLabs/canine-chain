package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostZKContract = "post_zk_contract"

var _ sdk.Msg = &MsgPostZKContract{}

func NewMsgPostZKContract(creator string, signer string, filesize int64, fid string, merkle string) *MsgPostZKContract {
	return &MsgPostZKContract{
		Creator:  creator,
		Signer:   signer,
		FileSize: filesize,
		Fid:      fid,
		Merkle:   merkle,
	}
}

func (msg *MsgPostZKContract) Route() string {
	return RouterKey
}

func (msg *MsgPostZKContract) Type() string {
	return TypeMsgPostZKContract
}

func (msg *MsgPostZKContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostZKContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostZKContract) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signee address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signee prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Fid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fid (%s)", err)
	}
	if prefix != FidPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklf`", prefix))
	}

	return nil
}
