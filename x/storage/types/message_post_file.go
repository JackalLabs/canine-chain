package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostFile = "post_file"

var _ sdk.Msg = &MsgPostFile{}

func NewMsgPostFile(creator string, merkle []byte, fileSize int64, proofInterval int64, proofType int64, maxProofs int64, note string) *MsgPostFile {
	return &MsgPostFile{
		Creator:       creator,
		Merkle:        merkle,
		FileSize:      fileSize,
		ProofInterval: proofInterval,
		ProofType:     proofType,
		MaxProofs:     maxProofs,
		Note:          note,
	}
}

func (msg *MsgPostFile) Route() string {
	return RouterKey
}

func (msg *MsgPostFile) Type() string {
	return TypeMsgPostFile
}

func (msg *MsgPostFile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostFile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostFile) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
